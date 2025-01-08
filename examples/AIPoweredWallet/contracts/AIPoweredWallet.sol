// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

interface AIKernel {
    event NewInference(
        uint256 indexed inferenceId,
        address indexed model,
        address indexed creator,
        uint256 value,
        uint256 originInferenceId
    );

    function infer(
        bytes calldata _data,
        bool _flag
    ) external payable returns (uint256 referenceId);

    function infer(
        bytes calldata _data
    ) external payable returns (uint256 referenceId);
}

interface PromptScheduler {
    enum InferenceStatus {
        Nil,
        Solving,
        Commit,
        Reveal,
        Processed,
        Killed,
        Transferred
    }

    enum AssignmentRole {
        Nil,
        Validating,
        Mining
    }

    enum Vote {
        Nil,
        Disapproval,
        Approval
    }

    struct Assignment {
        uint256 inferenceId;
        bytes32 commitment;
        bytes32 digest;
        uint40 revealNonce;
        address worker;
        AssignmentRole role;
        Vote vote;
        bytes output;
    }

    struct Inference {
        uint256[] assignments;
        bytes input;
        uint256 value;
        uint256 feeL2;
        uint256 feeTreasury;
        address modelAddress;
        uint40 submitTimeout;
        uint40 commitTimeout;
        uint40 revealTimeout;
        InferenceStatus status;
        address creator;
        address processedMiner;
        address referrer;
    }

    function getInferenceInfo(
        uint256 _inferenceId
    ) external view returns (Inference memory);

    function getAssignmentInfo(
        uint256 _assignmentId
    ) external view returns (Assignment memory);
}

contract AIPoweredWallet {
    struct TxInfo {
        address sender;
        address receiver;
        uint256 amount;
    }

    address public kernel;
    address public promptScheduler;
    mapping(address user => string) public context;
    mapping(uint256 inferId => TxInfo) public txInfo;

    event SuspiciousTransaction(uint256 indexed inferenceId, bytes prompt);
    event Sent(uint256 indexed inferenceId, TxInfo txInfo);

    constructor(address _kernelAddress, address _promptSchedulerAddress) {
        require(
            _kernelAddress != address(0) &&
                _promptSchedulerAddress != address(0),
            "AIPoweredWallet: Invalid address"
        );
        kernel = _kernelAddress;
        promptScheduler = _promptSchedulerAddress;
    }

    function buildRequest(
        string memory _prompt
    ) public pure returns (string memory) {
        string memory request = string.concat('{  "messages"');
        request = string.concat(request, " :[");
        request = string.concat(
            request,
            '{"role":"system","content":"You are a helpful assistant"},'
        );
        request = string.concat(request, '{"role":"user","content":"');
        request = string.concat(request, _prompt);
        request = string.concat(request, ' "}');
        request = string.concat(request, "],");
        request = string.concat(request, '"max_tokens":1024,');
        request = string.concat(
            request,
            '"model":"PrimeIntellect/INTELLECT-1-Instruct"'
        );
        request = string.concat(request, "}");
        return request;
    }

    function suspiciousTransaction(
        address _receiver,
        uint256 _amount
    ) external {
        string memory prompt = string.concat(
            "Assess the following Ethereum transaction history for any suspicious patterns. Respond with 'yes' if there is ANY indication of unusual or potentially malicious activity, however slight. Respond with 'no' ONLY if the transaction history is completely clean and normal. Your answer must be one word only. ",
            Strings.toHexString(msg.sender),
            " transfer ",
            Strings.toString(_amount),
            " wei to ",
            Strings.toHexString(_receiver),
            ". ",
            context[msg.sender]
        );

        string memory request = buildRequest(prompt);

        uint256 inferenceId = AIKernel(kernel).infer(bytes(request));
        txInfo[inferenceId] = TxInfo(msg.sender, _receiver, _amount);

        emit SuspiciousTransaction(inferenceId, bytes(request));
    }

    function send(uint256 _inferenceId) external payable {
        TxInfo memory info = txInfo[_inferenceId];

        require(info.sender == msg.sender, "AIPoweredWallet: Unauthorized");

        address receivedWallet = info.receiver;
        require(
            receivedWallet != address(0),
            "AIPoweredWallet: Invalid wallet address"
        );
        require(
            info.amount == msg.value,
            "AIPoweredWallet: Invalid transaction amount"
        );

        bytes memory result = fetchInferenceResult(_inferenceId);

        require(
            keccak256(result) == keccak256(abi.encodePacked("no")),
            "AIPoweredWallet: Suspicious transaction"
        );

        (bool success, ) = payable(receivedWallet).call{value: msg.value}("");
        require(success, "AIPoweredWallet: Transfer failed");

        context[msg.sender] = string.concat(
            context[msg.sender],
            Strings.toHexString(msg.sender),
            " transfer ",
            Strings.toString(msg.value),
            " wei to ",
            Strings.toHexString(receivedWallet),
            ". "
        );
    }

    function fetchInferenceResult(
        uint256 _inferenceId
    ) public view returns (bytes memory) {
        PromptScheduler.Inference memory inferInfo = PromptScheduler(
            promptScheduler
        ).getInferenceInfo(_inferenceId);

        if (inferInfo.assignments.length == 0) revert("Wait for inference");

        return
            PromptScheduler(promptScheduler)
                .getAssignmentInfo(inferInfo.assignments[0])
                .output;
    }
}
