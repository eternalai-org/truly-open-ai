// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library TransferHelper {
    bytes4 private constant SELECTOR_APPROVE =
        bytes4(keccak256(bytes("approve(address,uint256)")));
    bytes4 private constant SELECTOR_TRANSFER =
        bytes4(keccak256(bytes("transfer(address,uint256)")));
    bytes4 private constant SELECTOR_TRANSFER_FROM =
        bytes4(keccak256(bytes("transferFrom(address,address,uint256)")));

    error FailedApproval();
    error FailedTransfer();

    function safeApprove(address _token, address _to, uint256 _value) internal {
        (bool success, bytes memory data) = _token.call(
            abi.encodeWithSelector(SELECTOR_APPROVE, _to, _value)
        );
        if (!success || (data.length > 0 && !abi.decode(data, (bool)))) {
            revert FailedApproval();
        }
    }

    function safeTransfer(
        address _token,
        address _to,
        uint256 _value
    ) internal {
        (bool success, bytes memory data) = _token.call(
            abi.encodeWithSelector(SELECTOR_TRANSFER, _to, _value)
        );
        if (!success || (data.length > 0 && !abi.decode(data, (bool)))) {
            revert FailedTransfer();
        }
    }

    function safeTransferFrom(
        address _token,
        address _from,
        address _to,
        uint256 _value
    ) internal {
        (bool success, bytes memory data) = _token.call(
            abi.encodeWithSelector(SELECTOR_TRANSFER_FROM, _from, _to, _value)
        );
        if (!success || (data.length > 0 && !abi.decode(data, (bool)))) {
            revert FailedTransfer();
        }
    }

    function safeTransferNative(address _to, uint256 _value) internal {
        (bool success, ) = _to.call{value: _value}("");
        if (!success) revert FailedTransfer();
    }
}
