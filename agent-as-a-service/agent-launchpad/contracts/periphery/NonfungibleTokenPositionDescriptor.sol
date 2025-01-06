// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "../core/interfaces/IUniswapV3Pool.sol";
import "./libraries/SafeERC20Namer.sol";

import "./libraries/ChainId.sol";
import "./interfaces/INonfungiblePositionManager.sol";
import "./interfaces/INonfungibleTokenPositionDescriptor.sol";
import "./interfaces/IERC20Metadata.sol";
import "./libraries/PoolAddress.sol";
import "./libraries/NFTDescriptor.sol";
import "./libraries/TokenRatioSortOrder.sol";

/// @title Describes NFT token positions
/// @notice Produces a string containing the data URI for a JSON metadata string
contract NonfungibleTokenPositionDescriptor is
    INonfungibleTokenPositionDescriptor,
    OwnableUpgradeable
{
    address public WETH;
    /// @dev A null-terminated string
    bytes32 public nativeCurrencyLabelBytes;

    function initialize(
        address _WETH,
        bytes32 _nativeCurrencyLabelBytes
    ) external initializer {
        __Ownable_init();
        //
        WETH = _WETH;
        nativeCurrencyLabelBytes = _nativeCurrencyLabelBytes;
    }

    function setWETH(
        address _WETH,
        bytes32 _nativeCurrencyLabelBytes
    ) external onlyOwner {
        WETH = _WETH;
        nativeCurrencyLabelBytes = _nativeCurrencyLabelBytes;
    }

    /// @notice Returns the native currency label as a string
    function nativeCurrencyLabel() public view returns (string memory) {
        uint256 len = 0;
        while (len < 32 && nativeCurrencyLabelBytes[len] != 0) {
            len++;
        }
        bytes memory b = new bytes(len);
        for (uint256 i = 0; i < len; i++) {
            b[i] = nativeCurrencyLabelBytes[i];
        }
        return string(b);
    }

    /// @inheritdoc INonfungibleTokenPositionDescriptor
    function tokenURI(
        INonfungiblePositionManager positionManager,
        uint256 tokenId
    ) external view override returns (string memory) {
        (
            ,
            ,
            address token0,
            address token1,
            uint24 fee,
            int24 tickLower,
            int24 tickUpper,
            ,
            ,
            ,
            ,

        ) = positionManager.positions(tokenId);

        IUniswapV3Pool pool = IUniswapV3Pool(
            PoolAddress.computeAddress(
                positionManager.factory(),
                PoolAddress.PoolKey({token0: token0, token1: token1, fee: fee})
            )
        );

        bool _flipRatio = flipRatio(token0, token1, ChainId.get());
        address quoteTokenAddress = !_flipRatio ? token1 : token0;
        address baseTokenAddress = !_flipRatio ? token0 : token1;
        (, int24 tick, , , , , ) = pool.slot0();

        return
            NFTDescriptor.constructTokenURI(
                NFTDescriptor.ConstructTokenURIParams({
                    tokenId: tokenId,
                    quoteTokenAddress: quoteTokenAddress,
                    baseTokenAddress: baseTokenAddress,
                    quoteTokenSymbol: quoteTokenAddress == WETH
                        ? nativeCurrencyLabel()
                        : SafeERC20Namer.tokenSymbol(quoteTokenAddress),
                    baseTokenSymbol: baseTokenAddress == WETH
                        ? nativeCurrencyLabel()
                        : SafeERC20Namer.tokenSymbol(baseTokenAddress),
                    quoteTokenDecimals: IERC20Metadata(quoteTokenAddress)
                        .decimals(),
                    baseTokenDecimals: IERC20Metadata(baseTokenAddress)
                        .decimals(),
                    flipRatio: _flipRatio,
                    tickLower: tickLower,
                    tickUpper: tickUpper,
                    tickCurrent: tick,
                    tickSpacing: pool.tickSpacing(),
                    fee: fee,
                    poolAddress: address(pool)
                })
            );
    }

    function flipRatio(
        address token0,
        address token1,
        uint256 chainId
    ) public view returns (bool) {
        return
            tokenRatioPriority(token0, chainId) >
            tokenRatioPriority(token1, chainId);
    }

    function tokenRatioPriority(
        address token,
        uint256 chainId
    ) public view returns (int256) {
        if (token == WETH) {
            return TokenRatioSortOrder.DENOMINATOR;
        }
        return 0;
    }
}
