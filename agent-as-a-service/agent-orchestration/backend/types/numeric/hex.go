package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Hex struct {
	data []byte
}

func NewHexFromHex(hash string) *Hex {
	return NewHexFromBytes(common.HexToHash(hash).Bytes())
}

func NewHexFromBytes(data []byte) *Hex {
	return &Hex{
		data: copyData(data),
	}
}

func (n *Hex) Hex() string {
	return hexutil.Encode(n.data)
}

func (n *Hex) Bytes() []byte {
	return copyData(n.data)
}

func (n *Hex) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	data, err := hexutil.Decode(strings.Trim(s, `"`))
	if err != nil {
		return err
	}
	*n = *NewHexFromBytes(data)
	return nil
}

func (n *Hex) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, hexutil.Encode(n.data))), nil
}

func (n *Hex) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	var data string
	switch v := src.(type) {
	case []byte:
		data = hexutil.Encode(v)
	case string:
		data = v
	default:
		return errors.New("invalid data type")
	}
	dataN, err := hexutil.Decode(data)
	if err != nil {
		return err
	}
	*n = *NewHexFromBytes(dataN)
	return nil
}

func (n Hex) Value() (driver.Value, error) {
	return n.Bytes(), nil
}
