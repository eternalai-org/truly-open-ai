// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import './ABDKMath32x32.sol';

/// @notice The signed 32.32-binary fixed-point number representation, which can have up to 32 binary digits and up to 32
/// binary decimals. The values of this are bound by the minimum and the maximum values permitted by the underlying Solidity
/// type int64.

type Float32x32 is int64;

function fromInt(int256 x) pure returns (Float32x32) {
  unchecked {
    require (x >= -0x80000000 && x <= 0x7FFFFFFF);
    return Float32x32.wrap(int64 (x << 32));
  }
}

function toInt(Float32x32 x) pure returns (int64) {
  unchecked {
    return int64(Float32x32.unwrap(x) >> 32);
  }
}

function add(Float32x32 x, Float32x32 y) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.add(Float32x32.unwrap(x), Float32x32.unwrap(y)));
  }
}

function sub(Float32x32 x, Float32x32 y) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.sub(Float32x32.unwrap(x), Float32x32.unwrap(y)));
  }
}

function mul(Float32x32 x, Float32x32 y) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.mul(Float32x32.unwrap(x), Float32x32.unwrap(y)));
  }
}

function div(Float32x32 x, Float32x32 y) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.div(Float32x32.unwrap(x), Float32x32.unwrap(y)));
  }
}

function sqrt(Float32x32 x) pure returns (Float32x32) {
  unchecked {
    return x;
  }
}

function neg(Float32x32 x) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.neg(Float32x32.unwrap(x)));
  }
}

function abs(Float32x32 x) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.abs(Float32x32.unwrap(x)));
  }
}

function exp(Float32x32 x) pure returns (Float32x32) {
  unchecked {
    return Float32x32.wrap(ABDKMath32x32.exp(Float32x32.unwrap(x)));
  }
}  

function eq(Float32x32 x, Float32x32 y) pure returns (bool) {
  unchecked {
    return Float32x32.unwrap(x) == Float32x32.unwrap(y);
  }
}

function gt(Float32x32 x, Float32x32 y) pure returns (bool) {
  unchecked {
    return Float32x32.unwrap(x) > Float32x32.unwrap(y);
  }
}

function gte(Float32x32 x, Float32x32 y) pure returns (bool) {
  unchecked {
    return Float32x32.unwrap(x) >= Float32x32.unwrap(y);
  }
}

function lt(Float32x32 x, Float32x32 y) pure returns (bool) {
  unchecked {
    return Float32x32.unwrap(x) < Float32x32.unwrap(y);
  }
}

function lte(Float32x32 x, Float32x32 y) pure returns (bool) {
  unchecked {
    return Float32x32.unwrap(x) <= Float32x32.unwrap(y);
  }
}

function neq(Float32x32 x, Float32x32 y) pure returns (bool) {
  unchecked {
    return Float32x32.unwrap(x) != Float32x32.unwrap(y);
  }
}

function min(Float32x32 x, Float32x32 y) pure returns (Float32x32) {
  unchecked {
    return lt(x, y) ? x : y;
  }
}

function max(Float32x32 x, Float32x32 y) pure returns (Float32x32) {
  unchecked {
    return gt(x, y) ? x : y;
  }
}

using {
  add,
  div,
  mul,
  sub,
  neg,
  eq,
  gt,
  gte,
  lt,
  lte,
  neq,
  exp,
  min,
  max
} for Float32x32 global;

using {
  add as +,
  div as /,
  mul as *,
  sub as -,
  neg as -,
  eq as ==,
  gt as >,
  gte as >=,
  lt as <,
  lte as <=,
  neq as !=
} for Float32x32 global;
