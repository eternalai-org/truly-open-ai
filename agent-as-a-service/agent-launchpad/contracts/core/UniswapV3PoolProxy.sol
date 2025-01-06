// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Proxy} from "@openzeppelin/contracts/proxy/Proxy.sol";
import {IUniswapV3Factory} from "./interfaces/IUniswapV3Factory.sol";

/**
 * @title InitializedProxy
 */
contract UniswapV3PoolProxy is Proxy {
    //
    IUniswapV3Factory public immutable factory;

    // ======== Constructor =========
    constructor() {
        factory = IUniswapV3Factory(msg.sender);
    }

    /**
     * @dev This is a virtual function that should be overridden so it returns the address to which the fallback function
     * and {_fallback} should delegate.
     */
    function _implementation()
        internal
        view
        virtual
        override
        returns (address impl)
    {
        return factory.getUniswapV3PoolImplementation();
    }
}
