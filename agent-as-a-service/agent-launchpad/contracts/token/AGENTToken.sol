// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol";
import {ERC20Votes, EIP712} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";

contract AGENTToken is ERC20, ERC20Permit, ERC20Votes {
    constructor(
        string memory name,
        string memory symbol,
        uint256 amount,
        address recipient
    ) ERC20(name, symbol) ERC20Permit(name) {
        _mint(recipient, amount);
    }

    function _mint(
        address account,
        uint256 amount
    ) internal virtual override(ERC20, ERC20Votes) {
        return ERC20Votes._mint(account, amount);
    }

    function _burn(
        address account,
        uint256 amount
    ) internal virtual override(ERC20, ERC20Votes) {
        ERC20Votes._burn(account, amount);
    }

    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override(ERC20, ERC20Votes) {
        ERC20Votes._afterTokenTransfer(from, to, amount);
    }
}
