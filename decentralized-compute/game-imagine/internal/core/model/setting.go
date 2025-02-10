package model

import "strings"

type Setting struct {
	Model            `bson:",inline" json:",inline"`
	OperationAddress string  `json:"operation_address" bson:"operation_address"`
	OperationPrivKey string  `json:"operation_priv_key" bson:"operation_priv_key"`
	TreasurerAddress string  `json:"treasurer_address" bson:"treasurer_address"`
	GasFeePercentage float64 `json:"gas_fee_percentage" bson:"gas_fee_percentage"`
}

func (m Setting) CollectionName() string {
	return "setting"
}

// SetOperationAddress sets the operation address
func (m *Setting) SetOperationAddress(operationAddress string) *Setting {
	m.OperationAddress = operationAddress
	return m
}

// SetOperationPrivKey sets the operation private key
func (m *Setting) SetOperationPrivKey(operationPrivKey string) *Setting {
	m.OperationPrivKey = operationPrivKey
	return m
}

// SetGasFeePercentage sets the gas fee percentage
func (m *Setting) SetGasFeePercentage(gasFeePercentage float64) *Setting {
	m.GasFeePercentage = gasFeePercentage
	return m
}

// SetTreasurerAddress sets the treasurer address
func (m *Setting) SetTreasurerAddress(treasurerAddress string) *Setting {
	m.TreasurerAddress = treasurerAddress
	return m
}

func NewSetting(treasurerAddress string) *Setting {
	setting := &Setting{
		TreasurerAddress: strings.ToLower(treasurerAddress),
	}
	setting.Init()
	return setting
}

type CreateSettingRequest struct {
	TreasurerAddress string `json:"treasurer_address" query:"treasurer_address" bson:"treasurer_address"`
}

type UpdateSettingRequest struct {
	OperationAddress string `json:"operation_address" query:"operation_address" bson:"operation_address"`
}
