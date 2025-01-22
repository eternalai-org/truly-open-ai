// SPDX-License-Identifier: BSD-4-Clause
/*
 * Modified from ABDK Math 64.64 Smart Contract Library (https://github.com/abdk-consulting/abdk-libraries-solidity)
 * ABDK Math 64.64 Smart Contract Library.  Copyright © 2019 by ABDK Consulting.
 * Author: Mikhail Vladimirov <mikhail.vladimirov@gmail.com>
 */
pragma solidity ^0.8.0;

/**
 * Smart contract library of mathematical functions operating with signed
 * 32.32-bit fixed point numbers.  Signed 32.32-bit fixed point number is
 * basically a simple fraction whose numerator is signed 64-bit integer and
 * denominator is 2^32.  As long as denominator is always the same, there is no
 * need to store it, thus in Solidity signed 32.32-bit fixed point numbers are
 * represented by int64 type holding only the numerator.
 */
library ABDKMath32x32 {
  /*
   * Minimum value signed 32.32-bit fixed point number may have. 
   */
  int128 private constant MIN_32x32 = -0x8000000000000000;

  /*
   * Maximum value signed 32.32-bit fixed point number may have. 
   */
  int128 private constant MAX_32x32 = 0x7FFFFFFFFFFFFFFF;

  /**
   * Convert signed 256-bit integer number into signed 32.32-bit fixed point
   * number.  Revert on overflow.
   *
   * @param x signed 256-bit integer number
   * @return signed 32.32-bit fixed point number
   */
  function fromInt (int256 x) internal pure returns (int64) {
    unchecked {
      require (x >= -0x80000000 && x <= 0x7FFFFFFF);
      return int64 (x << 32);
    }
  }

  /**
   * Convert signed 32.32 fixed point number into signed 32-bit integer number
   * rounding down.
   *
   * @param x signed 32.32-bit fixed point number
   * @return signed 32-bit integer number
   */
  function toInt (int64 x) internal pure returns (int32) {
    unchecked {
      return int32 (x >> 32);
    }
  }

  /**
   * Calculate x + y.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @param y signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function add (int64 x, int64 y) internal pure returns (int64) {
    unchecked {
      int128 result = int128(x) + y;
      require (result >= MIN_32x32 && result <= MAX_32x32);
      return int64 (result);
    }
  }

  /**
   * Calculate x - y.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @param y signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function sub (int64 x, int64 y) internal pure returns (int64) {
    unchecked {
      int128 result = int128(x) - y;
      require (result >= MIN_32x32 && result <= MAX_32x32);
      return int64 (result);
    }
  }

  /**
   * Calculate x * y rounding down.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @param y signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function mul (int64 x, int64 y) internal pure returns (int64) {
    unchecked {
      int128 result = int128(x) * y >> 32;
      require (result >= MIN_32x32 && result <= MAX_32x32);
      return int64 (result);
    }
  }

  /**
   * Calculate x / y rounding towards zero.  Revert on overflow or when y is
   * zero.
   *
   * @param x signed 32.32-bit fixed point number
   * @param y signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function div (int64 x, int64 y) internal pure returns (int64) {
    unchecked {
      require (y != 0);
      int128 result = (int128 (x) << 32) / y;
      require (result >= MIN_32x32 && result <= MAX_32x32);
      return int64 (result);
    }
  }

  /**
   * Calculate -x.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function neg (int64 x) internal pure returns (int64) {
    unchecked {
      require (x != MIN_32x32);
      return -x;
    }
  }

  /**
   * Calculate |x|.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function abs (int64 x) internal pure returns (int64) {
    unchecked {
      require (x != MIN_32x32);
      return x < 0 ? -x : x;
    }
  }

  /**
   * Calculate binary exponent of x.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function exp_2 (int64 x) internal pure returns (int64) {
    unchecked {
      require (x < 0x2000000000); // Overflow

      if (x < -0x2000000000) return 0; // Underflow

      uint128 result = 0x8000000000000000;

      if (x & 0x80000000 > 0)
        result = result * 0x16A09E667F3BCC908 >> 64;
      if (x & 0x40000000 > 0)
        result = result * 0x1306FE0A31B7152DE >> 64;
      if (x & 0x20000000 > 0)
        result = result * 0x1172B83C7D517ADCD >> 64;
      if (x & 0x10000000 > 0)
        result = result * 0x10B5586CF9890F629 >> 64;
      if (x & 0x8000000 > 0)
        result = result * 0x1059B0D31585743AE >> 64;
      if (x & 0x4000000 > 0)
        result = result * 0x102C9A3E778060EE6 >> 64;
      if (x & 0x2000000 > 0)
        result = result * 0x10163DA9FB33356D8 >> 64;
      if (x & 0x1000000 > 0)
        result = result * 0x100B1AFA5ABCBED61 >> 64;
      if (x & 0x800000 > 0)
        result = result * 0x10058C86DA1C09EA1 >> 64;
      if (x & 0x400000 > 0)
        result = result * 0x1002C605E2E8CEC50 >> 64;
      if (x & 0x200000 > 0)
        result = result * 0x100162F3904051FA1 >> 64;
      if (x & 0x100000 > 0)
        result = result * 0x1000B175EFFDC76BA >> 64;
      if (x & 0x80000 > 0)
        result = result * 0x100058BA01FB9F96D >> 64;
      if (x & 0x40000 > 0)
        result = result * 0x10002C5CC37DA9491 >> 64;
      if (x & 0x20000 > 0)
        result = result * 0x1000162E525EE0547 >> 64;
      if (x & 0x10000 > 0)
        result = result * 0x10000B17255775C04 >> 64;
      if (x & 0x8000 > 0)
        result = result * 0x1000058B91B5BC9AE >> 64;
      if (x & 0x4000 > 0)
        result = result * 0x100002C5C89D5EC6C >> 64;
      if (x & 0x2000 > 0)
        result = result * 0x10000162E43F4F831 >> 64;
      if (x & 0x1000 > 0)
        result = result * 0x100000B1721BCFC99 >> 64;
      if (x & 0x800 > 0)
        result = result * 0x10000058B90CF1E6D >> 64;
      if (x & 0x400 > 0)
        result = result * 0x1000002C5C863B73F >> 64;
      if (x & 0x200 > 0)
        result = result * 0x100000162E430E5A1 >> 64;
      if (x & 0x100 > 0)
        result = result * 0x1000000B172183551 >> 64;
      if (x & 0x80 > 0)
        result = result * 0x100000058B90C0B48 >> 64;
      if (x & 0x40 > 0)
        result = result * 0x10000002C5C8601CC >> 64;
      if (x & 0x20 > 0)
        result = result * 0x1000000162E42FFF0 >> 64;
      if (x & 0x10 > 0)
        result = result * 0x10000000B17217FBA >> 64;
      if (x & 0x8 > 0)
        result = result * 0x1000000058B90BFCD >> 64;
      if (x & 0x4 > 0)
        result = result * 0x100000002C5C85FE3 >> 64;
      if (x & 0x2 > 0)
        result = result * 0x10000000162E42FF0 >> 64;
      if (x & 0x1 > 0)
        result = result * 0x100000000B17217F8 >> 64;


      result >>= uint128 (int128 (31 - (x >> 32)));
      require (result <= uint128 (int128 (MAX_32x32)));

      return int64 (int128 (result));
    }
  }

  /**
   * Calculate natural exponent of x.  Revert on overflow.
   *
   * @param x signed 32.32-bit fixed point number
   * @return signed 32.32-bit fixed point number
   */
  function exp (int64 x) internal pure returns (int64) {
    unchecked {
      require (x < 0x2000000000); // Overflow

      if (x < -0x2000000000) return 0; // Underflow

      return exp_2 (
          int64 (int128 (x) * 0x171547652B82FE177 >> 64));
    }
  }
}