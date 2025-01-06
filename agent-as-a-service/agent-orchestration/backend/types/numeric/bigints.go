package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type BigInts []*big.Int

func (n *BigInts) Uint64s() []*big.Int {
	return []*big.Int(*n)
}

func (n *BigInts) String() string {
	strs := []string{}
	for _, num := range *n {
		strs = append(strs, num.Text(10))
	}
	str := ""
	if len(strs) > 0 {
		str = strings.Join(strs, ",")
	}
	return fmt.Sprintf(`{%s}`, str)
}

func (n *BigInts) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	s = strings.Trim(s, ` `)
	s = strings.Trim(s, `"`)
	s = strings.Trim(s, `[`)
	s = strings.Trim(s, `]`)
	s = strings.Trim(s, `{`)
	s = strings.Trim(s, `}`)
	nums := []*big.Int{}
	if len(s) > 0 {
		strs := strings.Split(s, ",")
		for _, str := range strs {
			str = strings.TrimSpace(str)
			str = strings.Trim(str, `"`)
			num, ok := big.NewInt(0).SetString(str, 10)
			if !ok {
				return errors.New("invalid data type")
			}
			nums = append(nums, num)
		}
	}
	*n = BigInts(nums)
	return nil
}

func (n *BigInts) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	strs := []string{}
	for _, num := range *n {
		strs = append(strs, fmt.Sprintf(`"%s"`, num.Text(10)))
	}
	str := ""
	if len(strs) > 0 {
		str = strings.Join(strs, ",")
	}
	return []byte(fmt.Sprintf(`[%s]`, str)), nil
}

func (n *BigInts) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	data, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}
	s := string(data)
	s = strings.Trim(s, ` `)
	s = strings.Trim(s, `"`)
	s = strings.Trim(s, `{`)
	s = strings.Trim(s, `}`)
	nums := []*big.Int{}
	if len(s) > 0 {
		strs := strings.Split(s, ",")
		for _, str := range strs {
			str = strings.TrimSpace(str)
			num, ok := big.NewInt(0).SetString(str, 10)
			if !ok {
				return errors.New("invalid data type")
			}
			nums = append(nums, num)
		}
	}
	*n = BigInts(nums)
	return nil
}

func (n BigInts) Value() (driver.Value, error) {
	return n.String(), nil
}
