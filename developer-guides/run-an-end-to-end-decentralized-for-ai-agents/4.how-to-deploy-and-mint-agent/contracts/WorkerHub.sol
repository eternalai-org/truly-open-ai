// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {Random} from "./library/Random.sol";
import {TransferHelper} from "./library/TransferHelper.sol";
import {WorkerHubStorage, Set} from "./storages/WorkerHubStorage.sol";
import {IDAOToken} from "./tokens/IDAOToken.sol";
import {IStakingHub} from "./interfaces/IStakingHub.sol";

contract WorkerHub is
    WorkerHubStorage,
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    using Random for Random.Randomizer;
    using Set for Set.Uint256Set;
    using Set for Set.Bytes32Set;

    string private constant VERSION = "v0.0.2";
    uint256 internal constant PERCENTAGE_DENOMINATOR = 100_00;
    uint256 private constant BLOCK_PER_YEAR = 365 days / 2; // 2s per block

    receive() external payable {}

    function initialize(
        address _wEAI,
        address _l2Owner,
        address _treasury,
        address _daoToken,
        address _stakingHub,
        uint16 _feeL2Percentage,
        uint16 _feeTreasuryPercentage,
        uint8 _minerRequirement,
        uint40 _submitDuration,
        uint40 _commitDuration,
        uint40 _revealDuration,
        uint16 _feeRatioMinerValidor,
        uint256 _daoTokenReward,
        DAOTokenPercentage memory _daoTokenPercentage
    ) external initializer {
        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        require(
            _l2Owner != address(0) &&
                _treasury != address(0) &&
                _daoToken != address(0) &&
                _stakingHub != address(0) &&
                _wEAI != address(0),
            "Zero address"
        );

        l2Owner = _l2Owner;
        treasury = _treasury;
        daoToken = _daoToken;
        stakingHub = _stakingHub;
        feeL2Percentage = _feeL2Percentage;
        feeTreasuryPercentage = _feeTreasuryPercentage;
        feeRatioMinerValidator = _feeRatioMinerValidor;
        minerRequirement = _minerRequirement;

        daoTokenReward = _daoTokenReward;
        submitDuration = _submitDuration;
        commitDuration = _commitDuration;
        revealDuration = _revealDuration;
        daoTokenPercentage = _daoTokenPercentage;
        wEAI = _wEAI;
    }

    function version() external pure returns (string memory) {
        return VERSION;
    }

    function pause() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpause() external onlyOwner whenPaused {
        _unpause();
    }

    function _registerReferrer(address _referrer, address _referee) internal {
        if (_referrer == address(0) || _referee == address(0))
            revert InvalidData();
        if (referrerOf[_referee] != address(0)) revert AlreadySubmitted();

        referrerOf[_referee] = _referrer;
    }

    function registerReferrer(
        address[] memory _referrers,
        address[] memory _referees
    ) external onlyOwner {
        if (_referrers.length != _referees.length) revert InvalidData();

        for (uint256 i = 0; i < _referrers.length; i++) {
            _registerReferrer(_referrers[i], _referees[i]);
        }
    }

    function infer(
        bytes calldata _input,
        address _creator,
        bool _flag
    ) external payable whenNotPaused returns (uint256) {
        return _infer(_input, _creator, 0, _flag);
    }

    function infer(
        bytes calldata _input,
        address _creator
    ) external payable whenNotPaused returns (uint256) {
        return _infer(_input, _creator, 0, false);
    }

    function _infer(
        bytes calldata _input,
        address _creator,
        uint256 _scoringFee,
        bool _flag
    ) internal virtual returns (uint256) {
        IStakingHub.Model memory model = IStakingHub(stakingHub).getModelInfo(
            msg.sender
        );
        if (model.tier == 0) revert Unauthorized();

        uint256 inferenceId = ++inferenceNumber;
        Inference storage inference = inferences[inferenceId];

        uint256 value = msg.value - _scoringFee;
        uint256 feeL2 = (value * feeL2Percentage) / PERCENTAGE_DENOMINATOR;
        uint256 feeTreasury = (value * feeTreasuryPercentage) /
            PERCENTAGE_DENOMINATOR;

        inference.input = _input;
        inference.feeL2 = feeL2;
        inference.feeTreasury = feeTreasury;
        inference.value = value - feeL2 - feeTreasury;
        inference.creator = _creator;
        inference.referrer = referrerOf[_creator];
        inference.modelAddress = msg.sender;

        _assignMiners(inferenceId);

        emit NewInference(inferenceId, msg.sender, _creator, value, 0);
        emit RawSubmitted(
            inferenceId,
            msg.sender,
            _creator,
            value,
            0,
            _input,
            _flag
        );

        return inferenceId;
    }

    function _assignMiners(uint256 _inferenceId) internal {
        uint40 expiredAt = uint40(block.number + submitDuration);
        uint40 commitTimeout = expiredAt + commitDuration;
        inferences[_inferenceId].submitTimeout = expiredAt;
        inferences[_inferenceId].commitTimeout = commitTimeout;
        inferences[_inferenceId].revealTimeout = commitTimeout + revealDuration;
        inferences[_inferenceId].status = InferenceStatus.Solving;

        address model = inferences[_inferenceId].modelAddress;

        address[] memory miners = IStakingHub(stakingHub)
            .getMinerAddressesOfModel(model);
        uint256 minerLen = miners.length;

        if (minerLen < minerRequirement) revert NotEnoughMiners();

        uint256 n = minerRequirement;

        for (uint256 i = 0; i < n; ++i) {
            uint8 index = uint8(randomizer.randomUint256() % (minerLen - i));
            address miner = miners[index];

            miners[index] = miners[minerLen - i - 1];

            uint256 assignmentId = ++assignmentNumber;
            assignments[assignmentId].inferenceId = _inferenceId;
            assignments[assignmentId].worker = miner;
            assignments[assignmentId].role = AssignmentRole.Validating;

            assignmentsByMiner[miner].insert(assignmentId);
            assignmentsByInference[_inferenceId].insert(assignmentId);
            emit NewAssignment(assignmentId, _inferenceId, miner, expiredAt);
        }
    }

    function seizeMinerRole(uint256 _assignmentId) external {
        IStakingHub(stakingHub).updateEpoch();

        if (assignments[_assignmentId].worker != msg.sender)
            revert OnlyAssignedWorker();
        uint256 inferId = assignments[_assignmentId].inferenceId;
        if (inferences[inferId].processedMiner != address(0))
            revert AlreadySeized();

        assignments[_assignmentId].role = AssignmentRole.Mining;
        inferences[inferId].processedMiner = msg.sender;

        emit MinerRoleSeized(_assignmentId, inferId, msg.sender);
    }

    function _validatateSolution(bytes calldata _data) internal pure virtual {
        if (_data.length == 0) revert InvalidData();
    }

    function submitSolution(
        uint256 _assigmentId,
        bytes calldata _data
    ) external virtual whenNotPaused {
        IStakingHub(stakingHub).updateEpoch();
        _validatateSolution(_data);

        // Check whether miner is available (the miner had previously joined). The inactive miner is not allowed to submit solution.
        if (!IStakingHub(stakingHub).isMinerAddress(msg.sender))
            revert InvalidMiner();

        IStakingHub(stakingHub).validateModelOfMiner(msg.sender);

        Assignment memory clonedAssignments = assignments[_assigmentId];
        uint256 inferId = clonedAssignments.inferenceId;

        // Check the msg sender is the assigned miner
        if (msg.sender != clonedAssignments.worker) revert Unauthorized();
        if (clonedAssignments.role != AssignmentRole.Mining)
            revert InvalidRole();

        if (clonedAssignments.output.length != 0) revert AlreadySubmitted();

        Inference memory clonedInference = inferences[inferId];

        if (clonedInference.status != InferenceStatus.Solving) {
            revert InvalidInferenceStatus();
        }

        if (uint40(block.number) > clonedInference.submitTimeout)
            revert SubmitTimeout();

        Inference storage inference = inferences[inferId];

        assignments[_assigmentId].output = _data; //Record the solution
        bytes32 digest = keccak256(abi.encodePacked(inferId, _data)); //Record the solution
        assignments[_assigmentId].digest = digest;
        assignments[_assigmentId].commitment = digest;
        inference.status = InferenceStatus.Commit;
        inference.assignments.push(_assigmentId);

        if (!digests[inferId].hasValue(digest)) {
            digests[inferId].insert(digest);
        }
        countDigest[digest]++;

        emit InferenceStatusUpdate(inferId, InferenceStatus.Commit);
        emit SolutionSubmission(msg.sender, _assigmentId);
    }

    function commit(
        uint256 _assignId,
        bytes32 _commitment
    ) external virtual whenNotPaused {
        IStakingHub(stakingHub).updateEpoch();

        if (_commitment == 0) revert InvalidCommitment();

        Assignment storage assignment = assignments[_assignId];
        uint256 inferId = assignment.inferenceId;
        Inference storage inference = inferences[inferId];

        if (uint40(block.number) > inference.commitTimeout)
            revert CommitTimeout();
        if (inference.status != InferenceStatus.Commit) {
            revert InvalidInferenceStatus();
        }

        // Check the msg sender is the assigned miner
        if (msg.sender != assignment.worker) revert Unauthorized();
        if (assignment.role != AssignmentRole.Validating) revert InvalidRole();
        if (assignment.commitment != 0) revert AlreadyCommitted();

        assignment.commitment = _commitment;
        inference.assignments.push(_assignId);
        votingInfo[inferId].totalCommit++;

        emit CommitmentSubmission(msg.sender, _assignId, _commitment);

        if (
            votingInfo[inferId].totalCommit ==
            assignmentsByInference[inferId].size() - 1
        ) {
            inference.status = InferenceStatus.Reveal;
            emit InferenceStatusUpdate(inferId, InferenceStatus.Reveal);
        }
    }

    function reveal(
        uint256 _assignId,
        uint40 _nonce,
        bytes calldata _data
    ) external virtual whenNotPaused {
        IStakingHub(stakingHub).updateEpoch();

        _validatateSolution(_data);
        if (_nonce == 0) revert InvalidNonce();

        Assignment storage assignment = assignments[_assignId];
        if (assignment.revealNonce != 0) revert AlreadyRevealed();

        uint256 inferId = assignment.inferenceId;
        Inference storage inference = inferences[inferId];

        if (uint40(block.number) > inference.revealTimeout)
            revert RevealTimeout();
        if (
            inference.status != InferenceStatus.Commit &&
            inference.status != InferenceStatus.Reveal
        ) revert InvalidInferenceStatus();

        if (
            uint40(block.number) < inference.commitTimeout &&
            votingInfo[inferId].totalCommit !=
            assignmentsByInference[inferId].size() - 1
        ) revert CannotFastForward();

        if (inference.status == InferenceStatus.Commit) {
            inference.status = InferenceStatus.Reveal;
        }

        // Check the msg sender is the assigned miner
        if (msg.sender != assignment.worker) revert Unauthorized();
        if (assignment.role != AssignmentRole.Validating) revert InvalidRole();
        if (assignment.commitment == 0) revert NotCommitted();

        bytes32 commitment = assignment.commitment;
        bytes32 revealHash = keccak256(
            abi.encodePacked(_nonce, msg.sender, _data)
        );

        if (commitment != revealHash) revert InvalidReveal();
        bytes32 digest = keccak256(abi.encodePacked(inferId, _data));

        assignment.revealNonce = _nonce;
        assignment.output = _data;
        assignment.digest = digest;
        votingInfo[inferId].totalReveal++;

        if (!digests[inferId].hasValue(digest)) {
            digests[inferId].insert(digest);
        }
        countDigest[digest]++;

        emit RevealSubmission(msg.sender, _assignId, _nonce, _data);

        if (
            votingInfo[inferId].totalReveal == votingInfo[inferId].totalCommit
        ) {
            resolveInference(inferId);
        }
    }

    function _findMostVotedDigest(
        uint256 _inferenceId
    ) internal view returns (bytes32, uint8) {
        uint8 maxCount = 0;
        bytes32 mostVotedDigest = 0;
        bytes32[] memory digestArr = digests[_inferenceId].values;
        uint256 len = digests[_inferenceId].size();

        for (uint256 i = 0; i < len; i++) {
            bytes32 currDigest = digestArr[i];
            uint8 count = countDigest[currDigest];
            if (count > maxCount) {
                maxCount = count;
                mostVotedDigest = currDigest;
            }
        }
        return (mostVotedDigest, maxCount);
    }

    function _validateDAOSupplyIncrease(
        bool _isReferred
    ) internal view returns (bool notReachedLimit) {
        if (_isReferred) {
            notReachedLimit = IDAOToken(daoToken).validateSupplyIncrease(
                daoTokenReward
            );
        } else {
            notReachedLimit = IDAOToken(daoToken).validateSupplyIncrease(
                (daoTokenReward *
                    (PERCENTAGE_DENOMINATOR -
                        daoTokenPercentage.referrerPercentage -
                        daoTokenPercentage.refereePercentage)) /
                    PERCENTAGE_DENOMINATOR
            );
        }
    }

    function validateDAOSupplyIncrease(
        bool _isReferred
    ) external view returns (bool notReachedLimit) {
        return _validateDAOSupplyIncrease(_isReferred);
    }

    function _filterCommitment(
        uint256 _inferenceId
    ) internal virtual returns (bool) {
        (bytes32 mostVotedDigest, uint8 maxCount) = _findMostVotedDigest(
            _inferenceId
        );

        // Check the maxCount is greater than the voting requirement
        if (
            maxCount <
            _getThresholdValue(assignmentsByInference[_inferenceId].size())
        ) {
            return false;
        }

        uint256[] memory assignmentIds = inferences[_inferenceId].assignments;
        uint256 len = assignmentIds.length;
        bool isMatchMinerResult = assignments[assignmentIds[0]].digest ==
            mostVotedDigest;

        //EAI
        uint256 feeForMiner = 0;
        uint256 shareFeePerValidator = 0;
        uint256 remainValue = inferences[_inferenceId].value;
        // DAO token
        uint256 tokenForMiner = 0;
        uint256 shareTokenPerValidator = 0;
        uint256 remainToken = (daoTokenPercentage.minerPercentage *
            daoTokenReward) / PERCENTAGE_DENOMINATOR;

        // Calculate fee for miner and share fee for validators
        if (isMatchMinerResult) {
            //if miner result is correct, then fee for miner = feeRatioMinerValidator * remainValue / 10000
            feeForMiner =
                (remainValue * feeRatioMinerValidator) /
                PERCENTAGE_DENOMINATOR;
            shareFeePerValidator = (remainValue - feeForMiner) / (maxCount - 1);
            tokenForMiner =
                (remainToken * feeRatioMinerValidator) /
                PERCENTAGE_DENOMINATOR;
            shareTokenPerValidator =
                (remainToken - tokenForMiner) /
                (maxCount - 1);
        } else {
            //if miner result is incorrect, then fee for miner = 0 and all honest validators will share the remainValue
            shareFeePerValidator = remainValue / maxCount;
            shareTokenPerValidator = remainToken / maxCount;
        }

        for (uint256 i = 0; i < len; i++) {
            Assignment storage assignment = assignments[assignmentIds[i]];
            // Logically, when a worker calls the commit function, it proves that the worker is active.
            // Calling the reveal function is a natural consequence if the worker is honest.
            // Therefore, if a worker calls commit but doesn't call reveal, it is highly likely that they are dishonest,
            // leading to the decision to slash this worker.
            if (assignment.digest != mostVotedDigest) {
                assignment.vote = Vote.Disapproval;
                IStakingHub(stakingHub).slashMiner(assignment.worker, true); // Slash dishonest workers (miner and validators will be slashed in the same way)
            } else {
                // process for honest workers
                assignment.vote = Vote.Approval;
                if (assignment.role == AssignmentRole.Validating) {
                    // if it iss validator, then transfer share fee
                    if (shareFeePerValidator > 0) {
                        TransferHelper.safeTransferNative(
                            assignment.worker,
                            shareFeePerValidator
                        );
                    }
                } else {
                    if (feeForMiner > 0) {
                        // it is miner, if miner is honest, the feeForMiner is greater than 0
                        TransferHelper.safeTransferNative(
                            assignment.worker,
                            feeForMiner
                        );
                    }
                }
            }
        }

        // Transfer the mining fee to treasury
        if (inferences[_inferenceId].feeL2 > 0) {
            TransferHelper.safeTransferNative(
                l2Owner,
                inferences[_inferenceId].feeL2
            );
        }
        if (inferences[_inferenceId].feeTreasury > 0) {
            TransferHelper.safeTransferNative(
                treasury,
                inferences[_inferenceId].feeTreasury
            );
        }

        inferences[_inferenceId].status = InferenceStatus.Transferred;

        return true;
    }

    function setDAOTokenReward(uint256 _newDAOTokenReward) external onlyOwner {
        daoTokenReward = _newDAOTokenReward;
    }

    function resolveInference(
        uint256 _inferenceId
    ) public virtual whenNotPaused nonReentrant {
        IStakingHub(stakingHub).updateEpoch();

        Inference storage inference = inferences[_inferenceId];

        // If the inference is not processed (not seize or not submit solution),
        // we will refund all the value that user spent to get solution
        if (
            inference.status == InferenceStatus.Solving &&
            inference.submitTimeout < block.number &&
            inference.processedMiner != address(0)
        ) {
            inference.status = InferenceStatus.Killed;
            TransferHelper.safeTransferNative(
                inference.creator,
                inference.value + inference.feeL2 + inference.feeTreasury
            );

            // slash miner
            IStakingHub(stakingHub).slashMiner(inference.processedMiner, true);
        }

        if (
            inference.status == InferenceStatus.Commit &&
            inference.commitTimeout < block.number
        ) {
            // if 2/3 miners approve, then move to reveal phase
            if (
                votingInfo[_inferenceId].totalCommit + 1 >=
                _getThresholdValue(assignmentsByInference[_inferenceId].size())
            ) {
                inference.status == InferenceStatus.Reveal;
            } else {
                // else slash miner has not submitted solution and refund to user (because we do not know the correctly result)
                // Processed
                inference.status = InferenceStatus.Processed;
                TransferHelper.safeTransferNative(
                    inference.creator,
                    inference.value + inference.feeL2 + inference.feeTreasury
                );

                // slash validator not submitted commit hash
                uint256[] memory assignmentIds = assignmentsByInference[
                    _inferenceId
                ].values;
                for (uint i; i < assignmentIds.length; i++) {
                    //
                    if (
                        assignments[assignmentIds[i]].commitment == bytes32(0)
                    ) {
                        IStakingHub(stakingHub).slashMiner(
                            assignments[assignmentIds[i]].worker,
                            false
                        );
                    }
                }
            }
        }

        if (
            inference.status == InferenceStatus.Reveal &&
            (inference.revealTimeout < block.number ||
                votingInfo[_inferenceId].totalReveal ==
                votingInfo[_inferenceId].totalCommit)
        ) {
            // call kelvin function to get result
            // if 2/3 miners approve, then mark this infer as processed and trigger resolve infer again
            // else slash miner has not submitted solution and use miner's answer as result
            if (!_filterCommitment(_inferenceId)) {
                // edisable workers not call reveal and refund to user
                // Processed
                _handleNotEnoughVote(_inferenceId);
                inference.status = InferenceStatus.Processed;
            }
        }

        emit InferenceStatusUpdate(_inferenceId, inference.status);
    }

    function _handleNotEnoughVote(uint256 _inferenceId) internal virtual {
        Inference memory inference = inferences[_inferenceId];

        TransferHelper.safeTransferNative(
            inference.creator,
            inference.value + inference.feeL2 + inference.feeTreasury
        );

        // disable workers not call reveal
        uint256[] memory assignmentIds = assignmentsByInference[_inferenceId]
            .values;
        for (uint i; i < assignmentIds.length; i++) {
            //
            if (assignments[assignmentIds[i]].digest == bytes32(0)) {
                IStakingHub(stakingHub).slashMiner(
                    assignments[assignmentIds[i]].worker,
                    false
                );
            }
        }
    }

    function _getThresholdValue(uint x) internal pure returns (uint) {
        return (x * 2) / 3 + (x % 3 == 0 ? 0 : 1);
    }

    function getMinFeeToUse(
        address _modelAddress
    ) external view returns (uint256) {
        return IStakingHub(stakingHub).getMinFeeToUse(_modelAddress);
    }

    function getTreasuryAddress() external view returns (address) {
        return treasury;
    }

    function getInferenceInfo(
        uint256 _inferenceId
    ) external view returns (Inference memory) {
        return inferences[_inferenceId];
    }

    function getAssignmentsByInference(
        uint256 _inferenceId
    ) external view returns (uint256[] memory) {
        return assignmentsByInference[_inferenceId].values;
    }

    function getAssignmentInfo(
        uint256 _assignmentId
    ) external view returns (Assignment memory) {
        return assignments[_assignmentId];
    }
}
