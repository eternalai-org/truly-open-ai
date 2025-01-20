package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type (
	BaseTokenSymbol   string
	TokenSetupEnum    string
	ExternalAgentType string
)

const (
	BaseTokenSymbolBVM BaseTokenSymbol = "BVM"
	BaseTokenSymbolBTC BaseTokenSymbol = "BTC"
	BaseTokenSymbolETH BaseTokenSymbol = "ETH"
	BaseTokenSymbolFB  BaseTokenSymbol = "FB"
	BaseTokenSymbolEAI BaseTokenSymbol = "EAI"
	BaseTokenSymbolSOL BaseTokenSymbol = "SOL"

	TokenSetupEnumAutoCreate     TokenSetupEnum = "auto_create"
	TokenSetupEnumAutoCreateRune TokenSetupEnum = "auto_create_rune"
	TokenSetupEnumNoToken        TokenSetupEnum = "no_token"
	TokenSetupEnumLinkExisting   TokenSetupEnum = "link_existing"

	ExternalAgentTypeFarcaster ExternalAgentType = "farcaster"
)

type TwitterInfo struct {
	gorm.Model
	TwitterID         string `gorm:"unique_index"`
	TwitterAvatar     string
	TwitterUsername   string
	TwitterName       string
	TokenType         string
	ExpiresIn         int
	AccessToken       string `gorm:"type:text"`
	Scope             string
	RefreshToken      string `gorm:"type:text"`
	ExpiredAt         *time.Time
	OauthClientId     string
	OauthClientSecret string
	Description       string `gorm:"type:longtext"`
	RefreshError      string `gorm:"type:text"`
}

type AgentInfoAgentType uint

const (
	AgentInfoAgentTypeNormal        AgentInfoAgentType = 0
	AgentInfoAgentTypeReasoning     AgentInfoAgentType = 1
	AgentInfoAgentTypeKnowledgeBase AgentInfoAgentType = 2
)

func GetAgentFee(networkID uint64) numeric.BigFloat {
	switch networkID {
	case SHARDAI_CHAIN_ID:
		{
			return numeric.NewBigFloatFromString("600")
		}
	case ETHEREUM_CHAIN_ID:
		{
			return numeric.NewBigFloatFromString("1330")
		}
	case SOLANA_CHAIN_ID:
		{
			return numeric.NewBigFloatFromString("50")
		}
	case BITTENSOR_CHAIN_ID:
		{
			return numeric.NewBigFloatFromString("100")
		}
	case TRON_CHAIN_ID:
		{
			return numeric.NewBigFloatFromString("50")
		}
	default:
		{
			return numeric.NewBigFloatFromString("1")
		}
	}
}

type (
	AssistantStatus     string
	CreateTokenModeType string
)

const (
	AssistantStatusPending  AssistantStatus = "pending"
	AssistantStatusMinting  AssistantStatus = "minting"
	AssistantStatusUpdating AssistantStatus = "updating"
	AssistantStatusReady    AssistantStatus = "ready"
	AssistantStatusFailed   AssistantStatus = "failed"

	CreateTokenModeTypeNoToken      CreateTokenModeType = "no_token"
	CreateTokenModeTypeAutoCreate   CreateTokenModeType = "auto_create"
	CreateTokenModeTypeLinkExisting CreateTokenModeType = "link_existing"
)

type TwinStatus string

const (
	TwinStatusPending     TwinStatus = "pending"
	TwinStatusRunning     TwinStatus = "running"
	TwinStatusDoneSuccess TwinStatus = "done_success"
	TwinStatusDoneError   TwinStatus = "done_error"
)

type SocialInfo struct {
	AccountName string  `json:"account_name"`
	Fee         float64 `json:"fee"`
}

type AgentInfo struct {
	gorm.Model
	Version              string `gorm:"default:'1'"`
	NetworkID            uint64
	NetworkName          string
	OauthClientId        string
	OauthClientSecret    string
	AgentID              string             `gorm:"unique_index"`
	AgentType            AgentInfoAgentType `gorm:"default:0"`
	TwitterInfoID        uint               `gorm:"index"`
	TwitterInfo          *TwitterInfo
	TwitterID            string `gorm:"index"`
	TwitterUsername      string `gorm:"index"`
	TwitterVerified      bool   `gorm:"default:0"`
	AgentName            string
	SystemPrompt         string `gorm:"type:longtext"`
	UserPrompt           string `gorm:"type:longtext"`
	Creator              string
	AgentContractID      string
	AgentContractAddress string
	AgentNftMinted       bool `gorm:"default:0"`
	ScanEnabled          bool `gorm:"default:1"`
	ScanLatestTime       *time.Time
	ScanLatestId         string `gorm:"type:longtext"`
	ScanError            string
	TokenMode            string
	TokenName            string
	TokenSymbol          string
	TokenAddress         string
	TokenStatus          string
	TokenSupply          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenImageUrl        string
	TokenSignature       string
	TokenInfoID          uint
	TokenInfo            *AgentTokenInfo
	TokenImageInferID    string
	TokenPositionHash    string
	TokenDesc            string `gorm:"type:longtext"`
	TokenNetworkID       uint64
	Priority             int
	ETHAddress           string `gorm:"index"`
	TronAddress          string `gorm:"index"`
	SOLAddress           string `gorm:"index"`
	SummaryLatestTime    *time.Time
	EaiBalance           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	EaiWalletBalance     numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	IsAdvance            bool             `gorm:"default:0"`
	InferLatestTime      *time.Time
	ReplyLatestTime      *time.Time
	ReplyEnabled         bool `gorm:"default:0"`
	TipEthAddress        string
	TipBtcAddress        string
	TipSolAddress        string
	IsFaucet             bool `gorm:"default:0"`
	AgentSnapshotMission []*AgentSnapshotMission
	MintFee              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	ActionDelayed        int              `gorm:"default:900"`
	TmpTwitterID         string           `gorm:"index"`
	TmpTwitterInfo       *TwitterUser     `gorm:"foreignKey:twitter_id;AssociationForeignKey:tmp_twitter_id"`
	RefTweetID           uint
	Meme                 *Meme
	ActiveLatestTime     *time.Time
	Thumbnail            string
	FarcasterID          string `gorm:"index"`
	FarcasterUsername    string `gorm:"index"`
	MintHash             string
	InferFee             numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SystemReminder       string           `gorm:"type:longtext"`
	Status               AssistantStatus
	Uri                  string
	AgentBaseModel       string
	MetaData             string `gorm:"type:longtext"`
	Minter               string
	VerifiedNftOwner     bool
	NftAddress           string
	NftTokenID           string
	NftTokenImage        string
	NftOwnerAddress      string
	NftSignature         string
	NftSignMessage       string
	NftDelegateAddress   string
	NftPublicKey         string
	Bio                  string `gorm:"type:longtext"`
	Lore                 string `gorm:"type:longtext"`
	Knowledge            string `gorm:"type:longtext"`
	MessageExamples      string `gorm:"type:longtext"`
	PostExamples         string `gorm:"type:longtext"`
	Topics               string `gorm:"type:longtext"`
	Style                string `gorm:"type:longtext"`
	Adjectives           string `gorm:"type:longtext"`
	SocialInfo           string `gorm:"type:longtext"`
	InferenceCalls       int64
	ExternalChartUrl     string
	MissionTopics        string `gorm:"type:longtext"`
	GraphData            string `gorm:"type:longtext"`

	TwinTwitterUsernames    string           `gorm:"index"` // multiple twitter usernames, split by ,
	TwinStatus              TwinStatus       `gorm:"index"`
	KnowledgeBaseID         string           `gorm:"index"`
	TwinCallProcessRequest  string           `gorm:"type:longtext"`
	TwinCallProcessResponse string           `gorm:"type:longtext"`
	TwinFee                 numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TwinStartTrainingAt     *time.Time
	TwinEndTrainingAt       *time.Time
	TwinTrainingProgress    float64 `json:"twin_training_progress"`
	TwinTrainingMessage     string  `gorm:"type:longtext"`

	EstimateTwinDoneTimestamp *time.Time `json:"estimate_twin_done_timestamp"`
	TotalMintTwinFee          float64
	TwitterName               string           `gorm:"-"`
	MemePercent               float64          `gorm:"-"`
	MemeMarketCap             numeric.BigFloat `gorm:"-"`
	Counts                    int64            `gorm:"-"`
	AgentKBId                 uint             `json:"agent_kb_id"`
	KnowledgeBase             *KnowledgeBase   `json:"knowledge_base" gorm:"foreignKey:AgentKBId;references:AgentInfoId"`
}

func (m *AgentInfo) GetCharacterArrayString(charactor string) []string {
	data := []string{}
	_ = json.Unmarshal([]byte(charactor), &data)
	return data
}

func (m *AgentInfo) GetMessageExamples() [][]struct {
	User    string `json:"user"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
} {
	data := [][]struct {
		User    string `json:"user"`
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
	}{}
	_ = json.Unmarshal([]byte(m.MessageExamples), &data)
	return data
}

func (m *AgentInfo) GetSocialInfo() []*SocialInfo {
	data := []*SocialInfo{}
	_ = json.Unmarshal([]byte(m.SocialInfo), &data)
	return data
}

func (m *AgentInfo) GetStyle() map[string][]string {
	data := map[string][]string{}
	_ = json.Unmarshal([]byte(m.Style), &data)
	return data
}

type AgentInfoWithSnapshotPostActionsResponse struct {
	ID              uint               `json:"id"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	NetworkID       uint64             `json:"network_id"`
	NetworkName     string             `json:"network_name"`
	AgentID         string             `json:"agent_id"`
	AgentType       AgentInfoAgentType `json:"agent_type"`
	TwitterInfoID   uint               `json:"twitter_info_id"`
	TwitterID       string             `json:"twitter_id"`
	TwitterUsername string             `json:"twitter_username"`
	TwitterVerified bool               `json:"twitter_verified"`
	AgentName       string             `json:"agent_name"`
	SystemPrompt    string             `json:"system_prompt"`
	UserPrompt      string             `json:"user_prompt"`
	Creator         string             `json:"creator"`

	AgentContractID         string                         `json:"agent_contract_id"`
	AgentContractAddress    string                         `json:"agent_contract_address"`
	AgentSnapshotPostAction []*AgentSnapshotPostActionResp `json:"agent_snapshot_post_action"`
}

type AgentInfoResponse struct {
	ID              uint               `json:"id"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	NetworkID       uint64             `json:"network_id"`
	NetworkName     string             `json:"network_name"`
	AgentID         string             `json:"agent_id"`
	AgentType       AgentInfoAgentType `json:"agent_type"`
	TwitterInfoID   uint               `json:"twitter_info_id"`
	TwitterID       string             `json:"twitter_id"`
	TwitterUsername string             `json:"twitter_username"`
	TwitterVerified bool               `json:"twitter_verified"`
	AgentName       string             `json:"agent_name"`
	SystemPrompt    string             `json:"system_prompt"`
	UserPrompt      string             `json:"user_prompt"`
	Creator         string             `json:"creator"`

	AgentContractID       string                      `json:"agent_contract_id"`
	AgentContractAddress  string                      `json:"agent_contract_address"`
	AgentSnapshotMissions []*AgentSnapshotMissionResp `json:"agent_snapshot_missions"`
}

func (m *AgentInfo) GetAgentContractID() int64 {
	res, err := strconv.ParseInt(m.AgentContractID, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

type (
	AgentTwitterPostStatus string
	AgentTwitterPostType   string
)

const (
	AgentTwitterPostTypePost     AgentTwitterPostType = "post"
	AgentTwitterPostTypeReview   AgentTwitterPostType = "review"
	AgentTwitterPostTypeUnReview AgentTwitterPostType = "unreview"

	AgentTwitterPostStatusNew              AgentTwitterPostStatus = "new"
	AgentTwitterPostStatusInvalid          AgentTwitterPostStatus = "invalid"
	AgentTwitterConversationInvalid        AgentTwitterPostStatus = "conversation_invalid"
	AgentTwitterPostStatusValid            AgentTwitterPostStatus = "valid"
	AgentTwitterPostStatusInferNew         AgentTwitterPostStatus = "infer_new"
	AgentTwitterPostStatusInferSubmitted   AgentTwitterPostStatus = "infer_submitted"
	AgentTwitterPostStatusInferError       AgentTwitterPostStatus = "infer_error"
	AgentTwitterPostStatusInferFailed      AgentTwitterPostStatus = "infer_failed"
	AgentTwitterPostStatusInferResolved    AgentTwitterPostStatus = "infer_resolved"
	AgentTwitterPostStatusReplied          AgentTwitterPostStatus = "replied"
	AgentTwitterPostStatusRepliedError     AgentTwitterPostStatus = "replied_error"
	AgentTwitterPostStatusRepliedCancelled AgentTwitterPostStatus = "replied_cancelled"
	AgentTwitterPostStatusReposted         AgentTwitterPostStatus = "reposted"
	AgentTwitterPostStatusRepostedError    AgentTwitterPostStatus = "reposted_error"
)

type AgentTwitterPost struct {
	gorm.Model
	NetworkID             uint64
	AgentInfoID           uint `gorm:"index"`
	AgentInfo             *AgentInfo
	TwitterID             string
	TwitterUser           *TwitterUser `gorm:"foreignKey:twitter_id;AssociationForeignKey:twitter_id"`
	TwitterUsername       string
	TwitterName           string
	TwitterPostID         string           `gorm:"unique_index"`
	TwitterConversationId string           `gorm:"index"`
	TwitterParentPostID   string           `gorm:"index"`
	TwitterParentPost     *UserTwitterPost `gorm:"foreignKey:twitter_parent_post_id;AssociationForeignKey:twitter_post_id"`
	Type                  AgentTwitterPostType
	PostType              AgentSnapshotPostActionType
	PostAt                *time.Time `gorm:"index"`
	Content               string     `gorm:"type:longtext"`
	ExtractContent        string     `gorm:"type:longtext"`
	InferData             string     `gorm:"type:longtext"`
	ReplyContent          string     `gorm:"type:longtext"`
	ReplyPostId           string     `gorm:"index"`
	ReplyPostIds          string     `gorm:"type:text"`
	RePostId              string
	ImageUrl              string
	InferTxHash           string
	InferAt               *time.Time
	InferNum              uint                   `gorm:"default:0"`
	Status                AgentTwitterPostStatus `gorm:"index"`
	Prompt                string                 `gorm:"type:longtext"`
	Error                 string                 `gorm:"type:longtext"`
	FollowerCount         uint                   `gorm:"default:0"`
	ReplyPostAt           *time.Time
	ReplyPostReply        int
	ReplyPostView         int
	ReplyPostFavorite     int
	ReplyPostBookmark     int
	ReplyPostQuote        int
	ReplyPostRetweet      int
	RePostAt              *time.Time
	InscribeTxHash        string
	BitcoinTxHash         string
	ReplyScheduleAt       *time.Time
	Fee                   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	IsMigrated            bool             `gorm:"default:0"`
	TokenName             string
	TokenSymbol           string
	TokenAddress          string
	TokenImageUrl         string
	TokenDesc             string `gorm:"type:longtext"`
	TokenImageInferID     string
	TokenSignature        string
	IsCreateAgent         bool `gorm:"default:0"`
	AgentChain            string
	OwnerUsername         string
	OwnerTwitterID        string
}

func (m *AgentTwitterPost) GetAgentOnwerName() string {
	if m.OwnerUsername != "" {
		return m.OwnerUsername
	}
	return m.TwitterUsername
}

func (m *AgentTwitterPost) GetOwnerTwitterID() string {
	if m.OwnerTwitterID != "" {
		return m.OwnerTwitterID
	}
	return m.TwitterID
}

type TweetParseInfo struct {
	TokenName     string
	TokenDesc     string
	TokenSymbol   string
	TokenImageUrl string
	ChainName     string
	Owner         string
	Personality   string
	Type          string
	IsCreateToken bool
	IsIntellect   bool
	IsCreateAgent bool
	Description   string
}

type UserTwitterPost struct {
	gorm.Model
	TwitterID       string
	TwitterUsername string
	TwitterName     string
	TwitterPostID   string `gorm:"unique_index"`
	PostAt          *time.Time
	Content         string `gorm:"type:longtext"`
}

type AgentSummaryPostStatus string

const (
	AgentSummaryPostStatusNew            AgentSummaryPostStatus = "new"
	AgentSummaryPostStatusInvalid        AgentSummaryPostStatus = "invalid"
	AgentSummaryPostStatusValid          AgentSummaryPostStatus = "valid"
	AgentSummaryPostStatusInferNew       AgentSummaryPostStatus = "infer_new"
	AgentSummaryPostStatusInferSubmitted AgentSummaryPostStatus = "infer_submitted"
	AgentSummaryPostStatusInferError     AgentSummaryPostStatus = "infer_error"
	AgentSummaryPostStatusInferFailed    AgentSummaryPostStatus = "infer_failed"
	AgentSummaryPostStatusInferResolved  AgentSummaryPostStatus = "infer_resolved"
	AgentSummaryPostStatusPosted         AgentSummaryPostStatus = "posted"
	AgentSummaryPostStatusPostedError    AgentSummaryPostStatus = "posted_error"
)

type AgentTokenInfo struct {
	gorm.Model
	NetworkID       uint64
	NetworkName     string
	AgentInfoID     uint `gorm:"unique_index"`
	AgentInfo       *AgentInfo
	PriceUsd        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Price           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	BaseTokenSymbol string
	BaseTokenPrice  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	PoolAddress     string           `gorm:"index"`
	PriceLast24h    numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	VolumeLast24h   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalVolume     numeric.BigFloat `gorm:"index;type:decimal(36,18);default:0"`
	TipAmount       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	WalletBalance   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	DexUrl          string
	TotalSupply     int64
	UsdMarketCap    float64
	RaydiumPool     string
	PriceChange     float64
	DexId           string
	//
	Percent   float64          `gorm:"-"`
	MarketCap numeric.BigFloat `gorm:"-"`
}

type AgentTradeHistory struct {
	gorm.Model
	NetworkID         uint64
	TxHash            string `gorm:"index"`
	ContractAddress   string `gorm:"index"`
	EventId           string `gorm:"unique_index:pump_trade_histories_main_idx"`
	TxAt              time.Time
	RecipientAddress  string `gorm:"index"`
	RecipientUserID   uint   `gorm:"index"`
	RecipientUser     *User
	Amount0           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Amount1           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SqrtPriceX96      string
	Liquidity         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Tick              int64
	Price             numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	AgentTokenAddress string           `gorm:"index"`
	AgentInfoID       uint             `gorm:"index"`
	AgentInfo         *AgentInfo
	TokenInAddress    string
	AmountIn          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenOutAddress   string           `gorm:"index"`
	AmountOut         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	BaseTokenSymbol   string           `gorm:"default:'BTC'"`
	BaseTokenPrice    numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	BaseAmount        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenAmount       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	IsBuy             bool
}

type (
	AgentEaiTopupType   string
	AgentEaiTopupStatus string
)

const (
	AgentEaiTopupTypeDeposit         AgentEaiTopupType = "deposit"
	AgentEaiTopupTypeFaucet          AgentEaiTopupType = "faucet"
	AgentEaiTopupTypeSpent           AgentEaiTopupType = "spent"
	AgentEaiTopupTypeRefund          AgentEaiTopupType = "refund"
	AgentEaiTopupTypeRefundTrainFail AgentEaiTopupType = "refund_train_fail"

	AgentEaiTopupStatusNew        AgentEaiTopupStatus = "new"
	AgentEaiTopupStatusProcessing AgentEaiTopupStatus = "processing"
	AgentEaiTopupStatusDone       AgentEaiTopupStatus = "done"
	AgentEaiTopupStatusError      AgentEaiTopupStatus = "error"
	AgentEaiTopupStatusCancelled  AgentEaiTopupStatus = "cancelled"
)

type AgentEaiTopup struct {
	gorm.Model
	NetworkID      uint64
	AgentInfoID    uint `gorm:"index"`
	AgentInfo      *AgentInfo
	EventId        string            `gorm:"unique_index"`
	Type           AgentEaiTopupType `gorm:"default:'deposit'"`
	DepositAddress string
	ToAddress      string
	DepositTxHash  string
	TopupTxHash    string
	Amount         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Status         AgentEaiTopupStatus
	InscribeTxHash string
	Error          string `gorm:"type:longtext"`
	Toolset        string
}

type AuthCode struct {
	gorm.Model
	PublicCode string `gorm:"index"`
	SecretCode string `gorm:"index"`
	ETHAddress string `gorm:"index"`
	Expired    time.Time
}

type (
	AgentTipHistoryStatus string
	AgentTipHistorySymbol string
)

const (
	AgentTipHistoryStatusDone AgentTipHistoryStatus = "done"

	AgentTipHistorySymbolBTC AgentTipHistorySymbol = "BTC"
	AgentTipHistorySymbolETH AgentTipHistorySymbol = "ETH"
	AgentTipHistorySymbolSOL AgentTipHistorySymbol = "SOL"
)

func (m *AgentInfo) GetHeadSystemPrompt() string {
	var headSystemPrompt string
	if m.AgentName != "" && m.TwitterUsername != "" {
		headSystemPrompt = headSystemPrompt + "Your Twitter name is <twitter_name>. Your Twitter username is @<twitter_username>. People refer to you as <twitter_name>, @<twitter_username>, <twitter_username>.\n\n"
	}
	if m.TokenSymbol != "" && m.TokenAddress != "" {
		headSystemPrompt = headSystemPrompt + "You have a token. Your token name is <token_name>. Your token ticker is $<token_ticker>. People refer to your token as <token_name> or $<token_ticker>. Your token address is <token_address>. Your token was deployed on Solana.\n\n"
	}
	headSystemPrompt = strings.TrimSpace(headSystemPrompt)
	headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<twitter_name>", m.AgentName)
	headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<twitter_username>", m.TwitterUsername)
	headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<token_name>", m.TokenName)
	headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<token_ticker>", m.TokenSymbol)
	headSystemPrompt = strings.ReplaceAll(headSystemPrompt, "<token_address>", m.TokenAddress)
	return headSystemPrompt
}

func (m *AgentInfo) GetSystemPrompt() string {
	return fmt.Sprintf(`%s\n\n%s`, m.GetHeadSystemPrompt(), m.SystemPrompt)
}

type AgentTradeToken struct {
	gorm.Model
	NetworkID    uint64
	NetworkName  string
	TokenSymbol  string
	TokenName    string
	TokenAddress string
	CmcId        string
	Enabled      bool `gorm:"default:1"`
}

type AgentUriData struct {
	Name string `json:"name"`
}

type AgentExternalInfo struct {
	gorm.Model
	NetworkID        uint64            `gorm:"unique_index:agent_external_main_idx"`
	Type             ExternalAgentType `gorm:"unique_index:agent_external_main_idx"`
	AgentInfoID      uint              `gorm:"unique_index:agent_external_main_idx"`
	AgentInfo        *AgentInfo
	ExternalID       string
	ExternalUsername string
	ExternalName     string
}

type AgentChainFee struct {
	gorm.Model
	NetworkID uint64           `gorm:"unique_index"`
	InferFee  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	MintFee   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenFee  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}

type AgentStudioChildren struct {
	ID          uint                   `json:"id"`
	Idx         string                 `json:"idx"`
	CategoryIdx string                 `json:"categoryIdx"`
	Title       string                 `json:"title"`
	Data        map[string]interface{} `json:"data"`
}

type AgentStudio struct {
	ID          uint                   `json:"id"`
	Idx         string                 `json:"idx"`
	CategoryIdx string                 `json:"categoryIdx"`
	Title       string                 `json:"title"`
	Data        map[string]interface{} `json:"data"`
	Children    []*AgentStudioChildren `json:"children"`
}
