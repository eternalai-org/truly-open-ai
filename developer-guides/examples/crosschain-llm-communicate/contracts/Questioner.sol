// SPDX-License-Identifier: MIT

pragma solidity ^0.8.22;

import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';
import {OApp, MessagingFee, Origin} from '@layerzerolabs/oapp-evm/contracts/oapp/OApp.sol';
import {MessagingReceipt} from '@layerzerolabs/oapp-evm/contracts/oapp/OAppSender.sol';
import {IAIKernel} from './interfaces/IAIKernel.sol';
import {IPromptScheduler} from './interfaces/IPromptScheduler.sol';
import {RequestBuilder} from './RequestBuilder.sol';

/**
 * @title Questioner
 * @dev Cross-chain AI question maker contract using LayerZero
 */
contract Questioner is OApp {
        // ============ Events ============

    event QuestionAskedLocal(
        uint256 indexed inferId,
        address indexed sender,
        string question
    );
    event QuestionAskedCrossChain(
        uint32 indexed _dstEid,
        address indexed sender,
        string question
    );

    string constant version = '1.0.0';
    address public promptScheduler;
    address public aiKernel;
    string public modelName;
    uint256 public inferenceId;

    constructor(
        address _endpoint,
        address _delegate,
        address _promptScheduler,
        address _aiKernel,
        string memory _modelName
    ) OApp(_endpoint, _delegate) Ownable(_delegate) {
        if (_promptScheduler == address(0) || _aiKernel == address(0)) {
            revert('Invalid prompt scheduler address');
        }

        promptScheduler = _promptScheduler;
        aiKernel = _aiKernel;
        modelName = _modelName;
    }

    function askLocal(string memory _question) public returns (uint256) {
        string memory request = RequestBuilder.buildRequest(
            _question,
            modelName
        );

        uint256 inferId = IAIKernel(aiKernel).infer(
            abi.encodePacked(request),
            msg.sender
        );
        inferenceId = inferId;
        emit QuestionAskedLocal(inferId, msg.sender, _question);

        return inferId;
    }

    function askByL0(
        uint32 _dstEid,
        string memory _question,
        bytes calldata _options
    ) external payable returns (MessagingReceipt memory receipt) {
        bytes memory _payload = abi.encode(_question);
        receipt = _lzSend(
            _dstEid,
            _payload,
            _options,
            MessagingFee(msg.value, 0),
            payable(msg.sender)
        );

        emit QuestionAskedCrossChain(_dstEid, msg.sender, _question);
    }

    function quote(
        uint32 _dstEid,
        string memory _message,
        bytes memory _options,
        bool _payInLzToken
    ) public view returns (MessagingFee memory fee) {
        bytes memory payload = abi.encode(_message);
        fee = _quote(_dstEid, payload, _options, _payInLzToken);
    }

    function _lzReceive(
        Origin calldata /*_origin*/,
        bytes32 /*_guid*/,
        bytes calldata payload,
        address /*_executor*/,
        bytes calldata /*_extraData*/
    ) internal override {
        string memory decodeData = abi.decode(payload, (string));
        askLocal(decodeData);
    }

    function setModelName(string calldata _modelName) external onlyOwner {
        modelName = _modelName;
    }

    function fetchInferenceResult(
        uint256 _inferenceId
    ) public view returns (bytes memory) {
        IPromptScheduler.Inference memory inferInfo = IPromptScheduler(
            promptScheduler
        ).getInferenceInfo(_inferenceId);

        if (inferInfo.output.length == 0) revert('Wait for inference');

        return inferInfo.output;
    }
}
