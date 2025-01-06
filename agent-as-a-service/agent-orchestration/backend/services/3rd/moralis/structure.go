package moralis

import (
	"strings"
	"time"
)

const (
	KeyOwner                  string = "owner"
	KeyLimit                  string = "limit"
	KeyOffset                 string = "offset"
	KeyCurrsor                string = "cursor"
	KeyTokenAddresses         string = "token_addresses"
	KeyChain                  string = "chain"
	KeyFormat                 string = "format"
	KeyTotalRanges            string = "totalRanges"
	KeyRange                  string = "range"
	URLAssets                 string = "assets"
	URLNft                    string = "nft"
	WalletAddressBalance      string = "balance"
	WalletAddressTokenBalance string = "erc20"
	NormalizeMetadata         string = "normalizeMetadata"
)

var ChainToChainID = map[string]string{
	"mumbai": "80001",
	"goerli": "5",
}

type NftFilter struct {
	TokenAddress string `json:"token_address"`
	TokenId      string `json:"token_id"`
}

type MoralisGetMultipleNftsReqBody struct {
	Tokens            []NftFilter `json:"tokens"`
	NormalizeMetadata *bool       `json:"normalizeMetadata,omitempty"`
	MediaItems        *bool       `json:"media_items,omitempty"`
}

type MoralisGetMultipleNftsFilter struct {
	Chain   *string `json:"chain"`
	ReqBody MoralisGetMultipleNftsReqBody
}

type MoralisFilter struct {
	Chain             *string   `json:"chain"`
	Format            *string   `json:"format"`
	Limit             *int      `json:"limit"`
	TotalRanges       *int      `json:"totalRanges"`
	Range             *int      `json:"range"`
	Cursor            *string   `json:"cursor"`
	TokenAddresses    *[]string `json:"token_addresses"`
	NormalizeMetadata *bool     `json:"normalizeMetadata"`
}

type MoralisTokensResp struct {
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Cursor   string         `json:"cursor"`
	Result   []MoralisToken `json:"result"`
}

type MoralisBalanceResp struct {
	Balance      string `json:"balance"`
	TokenAddress string `json:"token_address"`
}

type MoralisToken struct {
	TokenAddress      string  `json:"token_address"`
	TokenID           string  `json:"token_id"`
	Amount            string  `json:"amount"`
	Owner             string  `json:"owner_of"`
	TokenHash         string  `json:"token_hash"`
	ContractType      string  `json:"contract_type"`
	Name              string  `json:"name"`
	Symbol            string  `json:"symbol"`
	TokenUri          string  `json:"token_uri"`
	MetadataString    *string `json:"metadata"`
	BlockNumberMinted string  `json:"block_number_minted"`

	// Custom
	Metadata *MoralisTokenMetadata `json:"normalized_metadata,omitempty"`
}

func (s MoralisToken) IsERC1155Type() bool {
	return strings.ToUpper(s.ContractType) == "ERC1155"
}

type InscribeBTC struct {
	Status         int    `json:"status"`
	ProjectTokenId string `json:"project_token_id"`
	InscriptionID  string `json:"inscription_id"`
}

type MoralisTokenMetadata struct {
	Image        string  `json:"image"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ExternalLink string  `json:"external_link"`
	AnimationUrl string  `json:"animation_url"`
	Attributes   []Trait `json:"attributes"`
}

type Trait struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
	RarityLabel string      `json:"rarity_label"`
	Count       int         `json:"count"`
	Percentage  float64     `json:"percentage"`
}

// Covalent structures
type CovalentNftTransactionFilter struct {
	Chain           *string
	ContractAddress string
	TokenID         string
}

type CovalentGetTokenHolderRequest struct {
	ContractAddress string
	Chain           *string
	Page            int32
	Limit           int32
}

type CovalentGetAllTokenHolderRequest struct {
	ContractAddress string
	Chain           *string
	Limit           int32
}

type CovalentGetTokenHolderData struct {
	UpdatedAt time.Time `json:"updated_at"`
	Items     []struct {
		ContractDecimals     int         `json:"contract_decimals"`
		ContractName         string      `json:"contract_name"`
		ContractTickerSymbol string      `json:"contract_ticker_symbol"`
		ContractAddress      string      `json:"contract_address"`
		SupportsErc          interface{} `json:"supports_erc"`
		LogoURL              string      `json:"logo_url"`
		Address              string      `json:"address"`
		Balance              string      `json:"balance"`
		TotalSupply          string      `json:"total_supply"`
		BlockHeight          int         `json:"block_height"`
	} `json:"items"`
	Pagination struct {
		HasMore    bool        `json:"has_more"`
		PageNumber int         `json:"page_number"`
		PageSize   int         `json:"page_size"`
		TotalCount interface{} `json:"total_count"`
	} `json:"pagination"`
}

type CovalentGetTokenHolderResponse struct {
	Data         CovalentGetTokenHolderData `json:"data"`
	Error        bool                       `json:"error"`
	ErrorMessage interface{}                `json:"error_message"`
	ErrorCode    interface{}                `json:"error_code"`
}

type CovalentGetNftTransactionResponse struct {
	Data         CovalentGetNftTransactionData `json:"data"`
	Error        bool                          `json:"error"`
	ErrorMessage interface{}                   `json:"error_message"`
	ErrorCode    interface{}                   `json:"error_code"`
}

type CovalentGetNftTransactionData struct {
	UpdatedAt time.Time `json:"updated_at"`
	Items     []struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsErc          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
		Type                 string   `json:"type"`
		NftTransactions      []struct {
			BlockSignedAt    time.Time   `json:"block_signed_at"`
			BlockHeight      int         `json:"block_height"`
			TxHash           string      `json:"tx_hash"`
			TxOffset         int         `json:"tx_offset"`
			Successful       bool        `json:"successful"`
			FromAddress      string      `json:"from_address"`
			FromAddressLabel interface{} `json:"from_address_label"`
			ToAddress        string      `json:"to_address"`
			ToAddressLabel   interface{} `json:"to_address_label"`
			Value            string      `json:"value"`
			ValueQuote       interface{} `json:"value_quote"`
			GasOffered       int         `json:"gas_offered"`
			GasSpent         int         `json:"gas_spent"`
			GasPrice         int         `json:"gas_price"`
			FeesPaid         string      `json:"fees_paid"`
			GasQuote         interface{} `json:"gas_quote"`
			GasQuoteRate     interface{} `json:"gas_quote_rate"`
			LogEvents        []struct {
				BlockSignedAt              time.Time   `json:"block_signed_at"`
				BlockHeight                int         `json:"block_height"`
				TxOffset                   int         `json:"tx_offset"`
				LogOffset                  int         `json:"log_offset"`
				TxHash                     string      `json:"tx_hash"`
				RawLogTopics               []string    `json:"raw_log_topics"`
				SenderContractDecimals     interface{} `json:"sender_contract_decimals"`
				SenderName                 interface{} `json:"sender_name"`
				SenderContractTickerSymbol interface{} `json:"sender_contract_ticker_symbol"`
				SenderAddress              string      `json:"sender_address"`
				SenderAddressLabel         interface{} `json:"sender_address_label"`
				SenderLogoURL              interface{} `json:"sender_logo_url"`
				RawLogData                 string      `json:"raw_log_data"`
				Decoded                    interface{} `json:"decoded"`
			} `json:"log_events"`
		} `json:"nft_transactions"`
	} `json:"items"`
	Pagination interface{} `json:"pagination"`
}

type MoralisMessage struct {
	Message string `json:"message"`
	Err     error
}

type NFTCollectionMetadata struct {
	TokenAddress          string    `json:"token_address"`
	Name                  string    `json:"name"`
	SyncedAt              time.Time `json:"synced_at"`
	Symbol                string    `json:"symbol"`
	ContractType          string    `json:"contract_type"`
	PossibleSpam          bool      `json:"possible_spam"`
	VerifiedCollection    bool      `json:"verified_collection"`
	CollectionLogo        string    `json:"collection_logo"`
	CollectionBannerImage string    `json:"collection_banner_image"`
	CollectionCategory    string    `json:"collection_category"`
	ProjectUrl            string    `json:"project_url"`
	WikiUrl               string    `json:"wiki_url"`
	DiscordUrl            string    `json:"discord_url"`
	TelegramUrl           string    `json:"telegram_url"`
	TwitterUsername       string    `json:"twitter_username"`
	InstagramUsername     string    `json:"instagram_username"`

	Tokens   []MoralisToken `json:"tokens"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Cursor   string         `json:"cursor"`
}
