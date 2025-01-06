package magiceden

type MagicedenInscriptionInfo struct {
	Id                            string   `json:"id"`
	ContentURI                    string   `json:"contentURI"`
	ContentType                   string   `json:"contentType"`
	ContentPreviewURI             string   `json:"contentPreviewURI"`
	Sat                           int64    `json:"sat"`
	SatName                       string   `json:"satName"`
	SatRarity                     string   `json:"satRarity"`
	SatBlockHeight                int      `json:"satBlockHeight"`
	SatBlockTime                  string   `json:"satBlockTime"`
	Satributes                    []string `json:"satributes"`
	GenesisTransaction            string   `json:"genesisTransaction"`
	GenesisTransactionBlockTime   string   `json:"genesisTransactionBlockTime"`
	GenesisTransactionBlockHeight int      `json:"genesisTransactionBlockHeight"`
	GenesisTransactionBlockHash   string   `json:"genesisTransactionBlockHash"`
	InscriptionNumber             int      `json:"inscriptionNumber"`
	Meta                          struct {
		Name       string `json:"name"`
		Attributes []struct {
			TraitType string `json:"trait_type"`
			Value     string `json:"value"`
		} `json:"attributes"`
	} `json:"meta"`
	Chain            string `json:"chain"`
	Owner            string `json:"owner"`
	CollectionSymbol string `json:"collectionSymbol"`
	Collection       struct {
		Symbol                         string        `json:"symbol"`
		Name                           string        `json:"name"`
		ImageURI                       string        `json:"imageURI"`
		Chain                          string        `json:"chain"`
		InscriptionIcon                string        `json:"inscriptionIcon"`
		Description                    string        `json:"description"`
		Supply                         int           `json:"supply"`
		TwitterLink                    string        `json:"twitterLink"`
		DiscordLink                    string        `json:"discordLink"`
		WebsiteLink                    string        `json:"websiteLink"`
		CreatedAt                      string        `json:"createdAt"`
		OverrideContentType            string        `json:"overrideContentType"`
		DisableRichThumbnailGeneration bool          `json:"disableRichThumbnailGeneration"`
		Labels                         []interface{} `json:"labels"`
		CreatorTipsAddress             string        `json:"creatorTipsAddress"`
		EnableCollectionOffer          bool          `json:"enableCollectionOffer"`
	} `json:"collection"`
	ItemType                   string `json:"itemType"`
	Location                   string `json:"location"`
	LocationBlockHeight        int    `json:"locationBlockHeight"`
	LocationBlockTime          string `json:"locationBlockTime"`
	LocationBlockHash          string `json:"locationBlockHash"`
	OutputValue                int    `json:"outputValue"`
	Output                     string `json:"output"`
	Listed                     bool   `json:"listed"`
	ListedAt                   string `json:"listedAt"`
	ListedPrice                int    `json:"listedPrice"`
	ListedMakerFeeBp           int    `json:"listedMakerFeeBp"`
	ListedSellerReceiveAddress string `json:"listedSellerReceiveAddress"`
	ListedForMint              bool   `json:"listedForMint"`
	SacAddress                 string `json:"sacAddress"`
	SacMerkleTreeSize          int    `json:"sacMerkleTreeSize"`
	DisplayName                string `json:"displayName"`
	LastSalePrice              int    `json:"lastSalePrice"`
	UpdatedAt                  string `json:"updatedAt"`
}

type CollectionInfo struct {
	Symbol                         string        `json:"symbol"`
	Name                           string        `json:"name"`
	ImageURI                       string        `json:"imageURI"`
	Chain                          string        `json:"chain"`
	InscriptionIcon                string        `json:"inscriptionIcon"`
	Description                    string        `json:"description"`
	Supply                         int           `json:"supply"`
	TwitterLink                    string        `json:"twitterLink"`
	DiscordLink                    string        `json:"discordLink"`
	WebsiteLink                    string        `json:"websiteLink"`
	CreatedAt                      string        `json:"createdAt"`
	OverrideContentType            string        `json:"overrideContentType"`
	DisableRichThumbnailGeneration bool          `json:"disableRichThumbnailGeneration"`
	Labels                         []interface{} `json:"labels"`
	CreatorTipsAddress             string        `json:"creatorTipsAddress"`
	EnableCollectionOffer          bool          `json:"enableCollectionOffer"`
}
