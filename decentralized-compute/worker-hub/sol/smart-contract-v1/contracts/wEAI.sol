pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract WrappedEAI is ERC20, Ownable {
    error FailedTransfer();
    // Event emitted when EAI is wrapped
    event EAIWrapped(address indexed user, uint256 amount);

    // Event emitted when WEAI is unwrapped
    event EAIUnwrapped(address indexed user, uint256 amount);

    constructor() ERC20("Wrapped EAI", "WEAI") {}

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    // Allows users to deposit EAI and mint an equivalent amount of WEAI
    function wrap() public payable {
        _mint(msg.sender, msg.value);
        emit EAIWrapped(msg.sender, msg.value);
    }

    // Allows users to burn WEAI and withdraw an equivalent amount of EAI
    function unwrap(uint256 amount) public {
        _burn(msg.sender, amount);

        (bool success, ) = payable(msg.sender).call{value: amount}("");
        if (!success) revert FailedTransfer();

        emit EAIUnwrapped(msg.sender, amount);
    }

    // Fallback function to receive EAI when the contract is called directly
    receive() external payable {
        wrap();
    }
}

/// @title Wrapped Token Contract for Testing
/// @notice Represents a wrapped token used in cross-chain asset bridging scenarios.
/// @dev This contract is intended for testing purposes only.
contract WrappedToken is ERC20, Ownable {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {}

    /// @notice Mints new tokens to a specified address.
    /// @dev Can only be called by the contract owner when testing. In the cross-chain asset bridging scenarios, we can use onlyBridge instead of onlyOwner.
    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    /// @notice Burns a specific amount of tokens from the caller's account.
    function burn(uint256 amount) public {
        _burn(msg.sender, amount);
    }
}
