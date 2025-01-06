package hiro

type HiroInscriptionInfo struct {
	Id                 string      `json:"id"`
	Number             int         `json:"number"`
	Address            string      `json:"address"`
	GenesisAddress     string      `json:"genesis_address"`
	GenesisBlockHeight int         `json:"genesis_block_height"`
	GenesisBlockHash   string      `json:"genesis_block_hash"`
	GenesisTxId        string      `json:"genesis_tx_id"`
	GenesisFee         string      `json:"genesis_fee"`
	GenesisTimestamp   int64       `json:"genesis_timestamp"`
	TxId               string      `json:"tx_id"`
	Location           string      `json:"location"`
	Output             string      `json:"output"`
	Value              string      `json:"value"`
	Offset             string      `json:"offset"`
	SatOrdinal         string      `json:"sat_ordinal"`
	SatRarity          string      `json:"sat_rarity"`
	SatCoinbaseHeight  int         `json:"sat_coinbase_height"`
	MimeType           string      `json:"mime_type"`
	ContentType        string      `json:"content_type"`
	ContentLength      int         `json:"content_length"`
	Timestamp          int64       `json:"timestamp"`
	CurseType          interface{} `json:"curse_type"`
	Recursive          bool        `json:"recursive"`
	RecursionRefs      interface{} `json:"recursion_refs"`
	Parent             interface{} `json:"parent"`
	Metadata           struct {
		Attributes []struct {
			TraitType string `json:"trait_type"`
			Value     string `json:"value"`
		} `json:"attributes"`
		Name string `json:"name"`
	} `json:"metadata"`
}
