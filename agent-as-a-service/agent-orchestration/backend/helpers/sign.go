package helpers

import (
	"fmt"
	"math/big"
	"strings"
)

func GetSignMsg(msg string) string {
	return fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
}

func AppendHexStrings(values ...string) string {
	var ret string
	for _, v := range values {
		if has0xPrefix(v) {
			v = v[2:]
		}
		ret = fmt.Sprintf("%s%s", ret, v)
	}
	return fmt.Sprintf("0x%s", strings.ToLower(ret))
}

func ParseHex2Hex(v string) string {
	if has0xPrefix(v) {
		v = v[2:]
	}
	return v
}

func ParseNumber2Hex(v string) string {
	n, _ := big.NewInt(0).SetString(v, 10)
	return ParseBigInt2Hex(n)
}

func ParseBigInt2Hex(v *big.Int) string {
	if v == nil {
		return strings.Repeat("00", 32)
	}
	val := fmt.Sprintf("%s%s", strings.Repeat("00", 32), v.Text(16))
	return val[len(val)-64:]
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func ParseAddress2Hex(v string) string {
	if has0xPrefix(v) {
		v = v[2:]
	}
	val := fmt.Sprintf("%s%s", strings.Repeat("00", 20), v)
	return val[len(val)-40:]
}
