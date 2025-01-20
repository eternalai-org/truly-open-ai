package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/ethereum/go-ethereum/common"
)

type AddressHash struct {
	data []byte
}

func NewAddressHashFromHex(address string) *AddressHash {
	return NewAddressHashFromBytes(helpers.HexToAddress(address).Bytes())
}

func NewAddressHashFromBytes(data []byte) *AddressHash {
	return &AddressHash{
		data: copyData(data),
	}
}

func (n *AddressHash) Hex() string {
	return common.BytesToAddress(n.data).Hex()
}

func (n *AddressHash) Bytes() []byte {
	return common.BytesToAddress(n.data).Bytes()
}

func (n *AddressHash) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		n = nil
		return nil
	}
	*n = *NewAddressHashFromBytes(common.Hex2Bytes(strings.Trim(s, `"`)))
	return nil
}

func (n *AddressHash) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, common.BytesToAddress(n.data).Hex())), nil
}

func (n *AddressHash) Scan(src interface{}) error {
	if src == nil {
		n = nil
		return nil
	}
	data, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}
	*n = *NewAddressHashFromBytes(common.BytesToAddress(data).Bytes())
	return nil
}

func (n AddressHash) Value() (driver.Value, error) {
	return n.Bytes(), nil
}
