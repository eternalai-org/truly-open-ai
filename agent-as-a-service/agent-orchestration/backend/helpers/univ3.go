package helpers

import (
	"errors"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var (
	Q128        = big.NewInt(1).Lsh(big.NewInt(1), 128)
	Q96         = big.NewInt(1).Lsh(big.NewInt(1), 96)
	Q32         = big.NewInt(1).Lsh(big.NewInt(1), 32)
	MIN_TICK    = big.NewInt(-887272)
	MAX_TICK    = new(big.Int).Neg(MIN_TICK)
	MAX_UINT256 = new(big.Int).Sub(
		new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil),
		big.NewInt(1),
	)
)

func HexToBigInt(s string) *big.Int {
	s = strings.TrimPrefix(s, "0x")
	n, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("bad data")
	}
	return n
}

func GetSqrtRatioAtTick(tick int64) *big.Int {
	absTick := new(big.Int).Abs(big.NewInt(tick))
	if absTick.Cmp(MAX_TICK) > 0 {
		panic("bad data")
	}
	ratio := HexToBigInt("0x100000000000000000000000000000000")
	if new(big.Int).And(absTick, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
		ratio = HexToBigInt("0xfffcb933bd6fad37aa2d162d1a594001")
	}
	if new(big.Int).And(absTick, HexToBigInt("0x02")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xfff97272373d413259a46990580e213a"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x4")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xfff2e50f5f656932ef12357cf3c7fdcc"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x8")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xffe5caca7e10e4e61c3624eaa0941cd0"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x10")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xffcb9843d60f6159c9db58835c926644"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x20")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xff973b41fa98c081472e6896dfb254c0"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x40")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xff2ea16466c96a3843ec78b326b52861"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x80")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xfe5dee046a99a2a811c461f1969c3053"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x100")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xfcbe86c7900a88aedcffc83b479aa3a4"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x200")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xf987a7253ac413176f2b074cf7815e54"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x400")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xf3392b0822b70005940c7a398e4b70f3"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x800")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xe7159475a2c29b7443b29c7fa6e889d9"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x1000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xd097f3bdfd2022b8845ad8f792aa5825"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x2000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0xa9f746462d870fdf8a65dc1f90e061e5"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x4000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0x70d869a156d2a1b890bb3df62baf32f7"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x8000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0x31be135f97d08fd981231505542fcfa6"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x10000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0x9aa508b5b7a84e1c677de54f3e99bc9"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x20000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0x5d6af8dedb81196699c329225ee604"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x40000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0x2216e584f5fa1ea926041bedfe98"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if new(big.Int).And(absTick, HexToBigInt("0x80000")).Cmp(big.NewInt(0)) != 0 {
		ratio = new(big.Int).Mul(ratio, HexToBigInt("0x48a170391f7dc42444e8fa2"))
		ratio = new(big.Int).Div(ratio, Q128)
	}
	if tick > 0 {
		ratio = new(big.Int).Div(MAX_UINT256, ratio)
	}
	remainder := new(big.Int).Mod(ratio, new(big.Int).Mul(big.NewInt(1), Q32))
	sqrtPriceX96 := new(big.Int).Div(ratio, Q32)
	if remainder.Cmp(big.NewInt(0)) != 0 {
		sqrtPriceX96 = new(big.Int).Add(sqrtPriceX96, big.NewInt(1))
	}
	return sqrtPriceX96
}

func GetLiquidityForAmount0(sqrtRatioAX96, sqrtRatioBX96, amount0 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	intermediate := new(big.Int).Mul(sqrtRatioAX96, sqrtRatioBX96)
	intermediate.Div(intermediate, Q96)
	result := new(big.Int).Mul(amount0, intermediate)
	result.Div(result, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
	return result
}

func GetLiquidityForAmount1(sqrtRatioAX96, sqrtRatioBX96, amount1 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	result := new(big.Int).Mul(amount1, Q96)
	result.Div(result, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
	return result
}

func GetLiquidityForAmounts(sqrtRatioCurrentX96, sqrtRatioAX96, sqrtRatioBX96, amount0, amount1 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	if sqrtRatioCurrentX96.Cmp(sqrtRatioAX96) < 0 {
		return GetLiquidityForAmount0(sqrtRatioAX96, sqrtRatioBX96, amount0)
	} else if sqrtRatioCurrentX96.Cmp(sqrtRatioBX96) < 0 {
		liquidity0 := GetLiquidityForAmount0(sqrtRatioCurrentX96, sqrtRatioBX96, amount0)
		liquidity1 := GetLiquidityForAmount1(sqrtRatioAX96, sqrtRatioCurrentX96, amount1)
		if liquidity0.Cmp(liquidity1) < 0 {
			return liquidity0
		} else {
			return liquidity1
		}
	} else {
		return GetLiquidityForAmount1(sqrtRatioAX96, sqrtRatioBX96, amount1)
	}
}

func GetAmount0ForLiquidity(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	result := new(big.Int).Mul(liquidity, Q96)
	result.Mul(result, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
	result.Div(result, sqrtRatioBX96)
	result.Div(result, sqrtRatioAX96)
	return result
}

func GetAmount1ForLiquidity(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	result := new(big.Int).Mul(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
	result.Div(result, Q96)
	return result
}

func GetAmountsForLiquidity(sqrtRatioCurrentX96, sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int) (*big.Int, *big.Int) {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	amountOut0 := big.NewInt(0)
	amountOut1 := big.NewInt(0)
	if sqrtRatioCurrentX96.Cmp(sqrtRatioAX96) < 0 {
		amountOut0 = GetAmount0ForLiquidity(sqrtRatioAX96, sqrtRatioBX96, liquidity)
	} else if sqrtRatioCurrentX96.Cmp(sqrtRatioBX96) > 0 {
		amountOut1 = GetAmount1ForLiquidity(sqrtRatioAX96, sqrtRatioBX96, liquidity)
	} else {
		if sqrtRatioCurrentX96.Cmp(sqrtRatioBX96) < 0 {
			amountOut0 = GetAmount0ForLiquidity(sqrtRatioCurrentX96, sqrtRatioBX96, liquidity)
		}
		if sqrtRatioAX96.Cmp(sqrtRatioCurrentX96) < 0 {
			amountOut1 = GetAmount1ForLiquidity(sqrtRatioAX96, sqrtRatioCurrentX96, liquidity)
		}
	}
	return amountOut0, amountOut1
}

func ParseGasFeeError(errText string) (*big.Int, error) {
	if strings.Contains(errText, "insufficient") && strings.Contains(errText, "cost") {
		errTexts := strings.Split(errText, " ")
		for idx, v := range errTexts {
			if v == "cost" {
				valueText := strings.Trim(errTexts[idx+1], ",")
				gasFee, ok := big.NewInt(0).SetString(valueText, 10)
				if !ok {
					return nil, errors.New(errText)
				}
				return gasFee, nil
			}
		}
	}
	return nil, errors.New(errText)
}

func PriceToTick(price float64) int64 {
	return int64(math.Log(price) / math.Log(1.0001))
}

func GetNumberFromString(strText string) int {
	re := regexp.MustCompile("[0-9]+")
	arr := re.FindAllString(strText, -1)
	if len(arr) > 0 {
		i, err := strconv.Atoi(arr[len(arr)-1])
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

func HexToAddress(address string) common.Address {
	addr := common.HexToAddress(address)
	if !strings.EqualFold(addr.Hex(), address) {
		panic("wrong address")
	}
	return addr
}

func HexToHash(hashHex string) common.Hash {
	hash := common.HexToHash(hashHex)
	if !strings.EqualFold(hash.Hex(), hashHex) {
		panic("wrong hash")
	}
	return hash
}
