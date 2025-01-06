package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/hiro"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/magiceden"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/moralis"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/opensea"
)

func (s *Service) OpenseaCollections(ctx context.Context, f opensea.OpenSeaFilterCollections) (*opensea.CollectionsResp, error) {
	if f.Inscription != nil {
		if *f.Inscription {
			return s.GetInscriptonCollections(ctx, f)
		}
	}

	c, err := s.openseaService.OpenseaGetCollections(ctx, f)
	if err != nil {
		return nil, err
	}
	return c, nil

}

func (s *Service) GetInscriptonCollections(ctx context.Context, f opensea.OpenSeaFilterCollections) (*opensea.CollectionsResp, error) {
	result := &opensea.CollectionsResp{}
	collections := []opensea.SingleCollectionResp{}

	a := magiceden.NewMagicedenService()
	resp1, err := a.GetInscriptionHardCode()
	if err != nil {
		return nil, err
	}

	limit := 100
	if f.Limit != 0 {
		limit = f.Limit
	}
	resp := resp1[0:limit]

	for _, item := range resp {

		c := []opensea.ContractResp{
			{
				Address: item.CollectionSymbol, // or provide the actual address
				Chain:   "ordinal",             // specify the chain if needed
			},
		}

		collection := opensea.SingleCollectionResp{
			Collection:  item.CollectionSymbol,
			Name:        item.Name,
			ImageUrl:    item.Image,
			Contracts:   c,
			Description: item.Description,
		}

		collections = append(collections, collection)
	}

	result.Collections = collections
	return result, nil
}

func (s *Service) GetInscriptonInfo(ctx context.Context, address string, f opensea.OpenSeaFilterCollections) (*moralis.NFTCollectionMetadata, error) {
	var data *opensea.SingleCollectionResp
	collections, err := s.GetInscriptonCollections(ctx, opensea.OpenSeaFilterCollections{})
	if err != nil {
		return nil, err
	}
	for _, coll := range collections.Collections {
		if coll.Collection == address {
			data = &coll
		}
	}
	if data == nil {
		return nil, errors.New("invalid collection")
	} else {
		result := &moralis.NFTCollectionMetadata{
			TokenAddress: data.Collection,
			Name:         data.Name,
			Symbol:       data.Collection,
			ContractType: "ordinal_nft",
			Tokens:       []moralis.MoralisToken{},
		}
		return result, nil
	}
}

func (s *Service) GetInscriptonInfoByTokenID(ctx context.Context, address, tokenId string, f opensea.OpenSeaFilterCollections) ([]moralis.MoralisToken, error) {
	if !strings.Contains(tokenId, "i0") {
		return nil, errors.New("invalid inscription id")
	}
	var data *opensea.SingleCollectionResp
	collections, err := s.GetInscriptonCollections(ctx, opensea.OpenSeaFilterCollections{})
	if err != nil {
		return nil, err
	}
	for _, coll := range collections.Collections {
		if coll.Collection == address {
			data = &coll
		}
	}
	if data == nil {
		return nil, errors.New("invalid collection")
	}

	magic := magiceden.NewMagicedenService()
	inscriptionInfo, err := magic.GetInscriptionInfo(tokenId)
	if err == nil {
		if inscriptionInfo.Collection.Symbol != data.Collection {
			return nil, errors.New("invalid collection")
		}
	}

	serviceHiro := hiro.NewHiroService(s.conf.HiroUrl)
	info, err := serviceHiro.GetInscriptionInfo(tokenId)
	if err != nil {
		return nil, err
	}

	item := moralis.MoralisToken{
		TokenAddress:   data.Collection,
		TokenID:        info.Id,
		Name:           info.Metadata.Name,
		Owner:          info.Address,
		Symbol:         data.Collection,
		MetadataString: nil,
		Metadata: &moralis.MoralisTokenMetadata{
			Image:        fmt.Sprintf("https://ord-mirror.magiceden.dev/content/%s", info.Id),
			Name:         info.Metadata.Name,
			Description:  data.Description,
			ExternalLink: "",
			AnimationUrl: "",
			Attributes:   []moralis.Trait{},
		},
	}

	if info.MimeType == "text/html" {
		item.Metadata.AnimationUrl = item.Metadata.Image
		item.Metadata.Image = ""
	}

	if item.Metadata.Description == "" {
		collectionInfo, err := magic.GetCollectionInfo(data.Collection)
		if err == nil {
			item.Metadata.Description = collectionInfo.Description
		}
	}

	if len(info.Metadata.Attributes) > 0 {
		for _, a := range info.Metadata.Attributes {
			trait := moralis.Trait{
				TraitType:   a.TraitType,
				Value:       a.Value,
				DisplayType: nil,
				MaxValue:    nil,
				TraitCount:  0,
				Order:       nil,
				RarityLabel: "",
				Count:       0,
				Percentage:  0,
			}
			item.Metadata.Attributes = append(item.Metadata.Attributes, trait)
		}
	}
	result := []moralis.MoralisToken{item}
	return result, nil
}

func (s *Service) GetNftCollectionMetadataByContract(ctx context.Context, address, cursor, pageSize string, f moralis.MoralisFilter) (*moralis.NFTCollectionMetadata, error) {
	dfChain := moralis.MORALIS_ETH
	harcodeCollection := s.openseaService.FindHardCodeCollectionByAddress(address)
	if harcodeCollection != nil {
		dfChain = harcodeCollection.Chain
	}

	moralisService := moralis.NewMoralisNfts(dfChain, s.conf.MoralisApiKey)
	collectionMetadata, err := moralisService.GetNftCollectionMetadataByContract(address, moralis.MoralisFilter{})
	if err != nil {
		return nil, err
	}
	limit := 10
	if pageSize != "" {
		limit, err = strconv.Atoi(pageSize)
		if err != nil {
			return nil, err
		}
	}
	normalizeMetadata := true
	filter := moralis.MoralisFilter{
		Limit:             &limit,
		NormalizeMetadata: &normalizeMetadata,
	}
	if cursor != "" {
		filter.Cursor = &cursor
	}
	collectionInfo, err := moralisService.GetNftByContract(address, filter)
	if err != nil {
		return nil, err
	}
	collectionMetadata.Tokens = collectionInfo.Result
	collectionMetadata.PageSize = collectionInfo.PageSize
	collectionMetadata.Page = collectionInfo.Page
	collectionMetadata.Cursor = collectionInfo.Cursor
	return collectionMetadata, nil
}

func (s *Service) GetNftCollectionMetadataByTokenID(ctx context.Context, address, tokenID string, f moralis.MoralisFilter) ([]moralis.MoralisToken, error) {
	dfChain := moralis.MORALIS_ETH
	harcodeCollection := s.openseaService.FindHardCodeCollectionByAddress(address)
	if harcodeCollection != nil {
		dfChain = harcodeCollection.Chain
	}

	moralisService := moralis.NewMoralisNfts(dfChain, s.conf.MoralisApiKey)
	chain := dfChain

	normalizeMetadata := true
	collectionMetadata, err := moralisService.GetMultipleNfts(moralis.MoralisGetMultipleNftsFilter{
		Chain: &chain,
		ReqBody: moralis.MoralisGetMultipleNftsReqBody{
			Tokens: []moralis.NftFilter{
				{TokenAddress: address, TokenId: tokenID},
			},
			NormalizeMetadata: &normalizeMetadata,
		},
	})
	if err != nil {
		return nil, err
	}
	return collectionMetadata, nil
}
