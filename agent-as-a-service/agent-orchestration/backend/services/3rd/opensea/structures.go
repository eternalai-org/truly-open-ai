package opensea

type User struct {
	Username string  `json:"username"`
	Account  Account `json:"account"`
}

type Account struct {
	ProfileImgUrl string `json:"profile_img_url"`
}

type OpenseaGetContract struct {
	Address    string `json:"address"`
	Collection string `json:"collection"`
	Name       string `json:"name"`
}

type OpenseaGetCollectionStats struct {
	Total OpenseaGetCollectionStatsTotal `json:"total"`
}

type OpenseaGetCollectionStatsTotal struct {
	FloorPrice interface{} `json:"floor_price"`
}

type OpenSeaFilterCollections struct {
	Chain           string
	CreatorUsername string
	IncludeHidden   *bool
	Limit           int    //The number of collections to return. Must be between 1 and 100. Default: 100
	Next            string //The cursor for the next page of results. This is returned from a previous request.
	OrderBy         string
	Inscription     *bool
}

type CollectionsResp struct {
	Collections []SingleCollectionResp `json:"collections"`
	Next        string                 `json:"next"`
}

type SingleCollectionResp struct {
	Collection              string         `json:"collection"`
	Name                    string         `json:"name"`
	Description             string         `json:"description"`
	ImageUrl                string         `json:"image_url"`
	BannerImageUrl          string         `json:"banner_image_url"`
	Owner                   string         `json:"owner"`
	SafelistStatus          string         `json:"safelist_status"`
	Category                string         `json:"category"`
	IsDisabled              bool           `json:"is_disabled"`
	IsNsfw                  bool           `json:"is_nsfw"`
	TraitOffersEnabled      bool           `json:"trait_offers_enabled"`
	CollectionOffersEnabled bool           `json:"collection_offers_enabled"`
	OpenseaUrl              string         `json:"opensea_url"`
	ProjectUrl              string         `json:"project_url"`
	WikiUrl                 string         `json:"wiki_url"`
	DiscordUrl              string         `json:"discord_url"`
	TelegramUrl             string         `json:"telegram_url"`
	TwitterUsername         string         `json:"twitter_username"`
	InstagramUsername       string         `json:"instagram_username"`
	Chain                   string         `json:"chain"`
	Contracts               []ContractResp `json:"contracts"`
}

type ContractResp struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
}
