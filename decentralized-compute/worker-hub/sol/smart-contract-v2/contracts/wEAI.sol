pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract WrappedEAI is ERC20, Ownable {
    // Event emitted when EAI is wrapped
    event EAIWrapped(address indexed user, uint256 amount);

    // Event emitted when WEAI is unwrapped
    event EAIUnwrapped(address indexed user, uint256 amount);

    error FailedTransfer();

    constructor() ERC20("Wrapped EAI", "WEAI") {}

    receive() external payable {
        wrap();
    }

    // Allows users to deposit EAI and mint an equivalent amount of WEAI
    function wrap() public payable {
        _mint(msg.sender, msg.value);
        emit EAIWrapped(msg.sender, msg.value);
    }

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to,amount);
    }

    // Allows users to burn WEAI and withdraw an equivalent amount of EAI
    function unwrap(uint256 amount) public {
        _burn(msg.sender, amount);

        (bool success, ) = msg.sender.call{value: amount}("");
        if (!success) revert FailedTransfer();

        emit EAIUnwrapped(msg.sender, amount);
    }
}
