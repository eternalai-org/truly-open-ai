// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IDAOToken} from "./IDAOToken.sol";
import {IWorkerHub} from "../interfaces/IWorkerHub.sol";

contract DAOToken is IDAOToken, ERC20Upgradeable, OwnableUpgradeable {
    IWorkerHub public workerHub;
    uint256 public MAX_SUPPLY_CAP;

    uint256[100] private __gap;

    modifier onlyWorkerHub() {
        require(
            msg.sender == address(workerHub),
            "Caller is not the workerHub"
        );
        _;
    }

    function initialize(
        string memory _name,
        string memory _symbol,
        uint256 _MAX_SUPPLY_CAP
    ) public initializer {
        require(_MAX_SUPPLY_CAP > 0, "Invalid supply cap");

        __ERC20_init(_name, _symbol);
        __Ownable_init();
        MAX_SUPPLY_CAP = _MAX_SUPPLY_CAP;
    }

    function mint(address to, uint256 amount) public onlyWorkerHub {
        require(
            totalSupply() + amount <= MAX_SUPPLY_CAP,
            "Max supply exceeded"
        );
        _mint(to, amount);
    }

    function mintBatch(
        address[] memory to,
        uint256[] memory amount
    ) external onlyWorkerHub {
        require(to.length == amount.length, "Length mismatch");

        for (uint256 i = 0; i < to.length; i++) {
            require(
                totalSupply() + amount[i] <= MAX_SUPPLY_CAP,
                "Max supply exceeded"
            );
            _mint(to[i], amount[i]);
        }
    }

    function validateSupplyIncrease(
        uint256 _amount
    ) external view returns (bool) {
        return totalSupply() + _amount <= MAX_SUPPLY_CAP;
    }

    function updateWorkerHub(address _workerHub) external onlyOwner {
        require(_workerHub != address(0), "WorkerHub is the zero address");
        workerHub = IWorkerHub(_workerHub);
    }

    function getMaxSupply() external view returns (uint256) {
        return MAX_SUPPLY_CAP;
    }
}
