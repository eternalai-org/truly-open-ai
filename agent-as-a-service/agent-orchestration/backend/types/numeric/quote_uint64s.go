package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type QuoteUint64s []uint64

func (n *QuoteUint64s) Uint64s() []uint64 {
	return []uint64(*n)
}

func (n *QuoteUint64s) String() string {
	strs := []string{}
	for _, num := range *n {
		strs = append(strs, strconv.FormatUint(num, 10))
	}
	str := ""
	if len(strs) > 0 {
		str = strings.Join(strs, ",")
	}
	return fmt.Sprintf(`{%s}`, str)
}

func (n *QuoteUint64s) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	s = strings.Trim(s, ` `)
	s = strings.Trim(s, `"`)
	s = strings.Trim(s, `{`)
	s = strings.Trim(s, `}`)
	nums := []uint64{}
	if len(s) > 0 {
		strs := strings.Split(s, ",")
		for _, str := range strs {
			str = strings.TrimSpace(str)
			num, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}
	}
	*n = QuoteUint64s(nums)
	return nil
}

func (n *QuoteUint64s) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	strs := []string{}
	for _, num := range *n {
		strs = append(strs, strconv.FormatUint(num, 10))
	}
	str := ""
	if len(strs) > 0 {
		str = strings.Join(strs, ",")
	}
	return []byte(fmt.Sprintf(`{%s}`, str)), nil
}

func (n *QuoteUint64s) Scan(src interface{}) error {
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
	nums := []uint64{}
	if len(s) > 0 {
		strs := strings.Split(s, ",")
		for _, str := range strs {
			str = strings.TrimSpace(str)
			num, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}
	}
	*n = QuoteUint64s(nums)
	return nil
}

func (n QuoteUint64s) Value() (driver.Value, error) {
	return n.String(), nil
}
