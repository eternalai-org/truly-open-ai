package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

func BigFloat2Text(n *big.Float) string {
	v := n.Text('f', 64)
	if strings.Contains(v, ".") {
		v = strings.TrimRight(v, "0")
		v = strings.TrimRight(v, ".")
	}
	return v
}

func BigFloat2TextDecimals(n *big.Float, decimals int) string {
	v := n.Text('f', 64)
	if strings.Contains(v, ".") {
		v = strings.TrimRight(v, "0")
		v = strings.TrimRight(v, ".")
	}
	if strings.Contains(v, ".") {
		vs := strings.Split(v, ".")
		vd := vs[1]
		if len(vd) > decimals {
			vd = vd[0:decimals]
		}
		v = fmt.Sprintf("%s.%s", vs[0], vd)
	}
	v = strings.TrimRight(v, "0")
	return v
}

type Decimal struct {
	decimal.Decimal
}

func (n *Decimal) ToDecimal() decimal.Decimal {
	return n.Decimal
}

func (n *Decimal) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	s = strings.Trim(s, `"`)
	d, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	*n = Decimal{
		d,
	}
	return nil
}

func (n *Decimal) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	s := n.Decimal.String()
	return []byte(fmt.Sprintf(`"%s"`, s)), nil
}

func (n *Decimal) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}
	s := string(b)
	if s == "" {
		n = nil
		return nil
	}
	d, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	*n = Decimal{
		d,
	}
	return nil
}

func (n Decimal) Value() (driver.Value, error) {
	return n.Decimal.String(), nil
}

// BigFloat

func NewFloatFromString(s string) *big.Float {
	if s == "" {
		s = "0"
	}
	n := BigFloat{
		*big.NewFloat(0),
	}
	err := n.UnmarshalJSON([]byte(s))
	if err != nil {
		panic(err)
	}
	return &n.Float
}

func NewBigFloatFromString(s string) BigFloat {
	if s == "" {
		s = "0"
	}
	n := BigFloat{
		*big.NewFloat(0),
	}
	err := n.UnmarshalJSON([]byte(s))
	if err != nil {
		panic(err)
	}
	return n
}

func NewBigFloatFromFloat(value *big.Float) BigFloat {
	n := BigFloat{
		*value,
	}
	return n
}

type BigFloat struct {
	big.Float
}

func (n *BigFloat) BigFloat() *big.Float {
	return &n.Float
}

func (n *BigFloat) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	s = strings.Trim(s, `"`)
	m, ok := big.NewFloat(0).SetPrec(1024).SetString(s)
	if !ok {
		return errors.New("invalid data type")
	}
	m = m.SetPrec(1024)
	*n = BigFloat{
		*m,
	}
	return nil
}

func (n *BigFloat) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, BigFloat2Text(&n.Float))), nil
}

func (n *BigFloat) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}
	s := string(b)
	if s == "" {
		n = nil
		return nil
	}
	m, ok := big.NewFloat(0).SetPrec(1024).SetString(s)
	if !ok {
		return errors.New("invalid data type")
	}
	m.SetPrec(1024)
	*n = BigFloat{
		*m,
	}
	return nil
}

func (n BigFloat) Value() (driver.Value, error) {
	return BigFloat2Text(&n.Float), nil
}

func (n BigFloat) ToString() string {
	return BigFloat2Text(&n.Float)
}

func (n BigFloat) ToStringWithDecimals(decimals int) string {
	return BigFloat2TextDecimals(&n.Float, decimals)
}

// BigFloat

type BigInt struct {
	big.Int
}

func NewBigIntFromInt(value *big.Int) BigInt {
	n := BigInt{
		*value,
	}
	return n
}

func (n *BigInt) BigInt() *big.Int {
	return &n.Int
}

func (n *BigInt) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	s = strings.Trim(s, `"`)
	m, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return errors.New("invalid data type")
	}
	*n = BigInt{
		*m,
	}
	return nil
}

func (n *BigInt) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	s := n.String()
	return []byte(fmt.Sprintf(`"%s"`, s)), nil
}

func (n *BigInt) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}
	s := string(b)
	if s == "" {
		n = nil
		return nil
	}
	m, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return errors.New("invalid data type")
	}
	*n = BigInt{
		*m,
	}
	return nil
}

func (n BigInt) Value() (driver.Value, error) {
	return n.String(), nil
}

func ToBigInts(nums []*big.Int) []BigInt {
	rets := []BigInt{}
	for _, v := range nums {
		rets = append(rets, NewBigIntFromInt(v))
	}
	return rets
}
