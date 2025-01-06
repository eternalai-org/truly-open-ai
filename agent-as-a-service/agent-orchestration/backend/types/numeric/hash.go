package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Hash struct {
	data []byte
}

func NewHashFromHex(hash string) *Hash {
	return NewHashFromBytes(common.HexToHash(hash).Bytes())
}

func NewHashFromBytes(data []byte) *Hash {
	return &Hash{
		data: copyData(data),
	}
}

func (n *Hash) Hex() string {
	return common.BytesToHash(n.data).Hex()
}

func (n *Hash) Bytes() []byte {
	return common.BytesToHash(n.data).Bytes()
}

func (n *Hash) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	*n = *NewHashFromBytes(common.Hex2Bytes(strings.Trim(s, `"`)))
	return nil
}

func (n *Hash) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, common.BytesToHash(n.data).Hex())), nil
}

func (n *Hash) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	data, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}
	*n = *NewHashFromBytes(common.BytesToHash(data).Bytes())
	return nil
}

func (n Hash) Value() (driver.Value, error) {
	return n.Bytes(), nil
}
