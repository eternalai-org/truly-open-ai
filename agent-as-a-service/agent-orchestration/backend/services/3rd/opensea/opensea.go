package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
)

const (
	API_URL string = "https://api.opensea.io/api"
)

type OpenseaService struct {
	apiKey  string //b5400021ebc9489a9fe8a4544f663b53
	headers map[string]string
}

type HardCodeCollection struct {
	Address string
	Chain   string
	Slug    string
	ChainID int
}

func (s *OpenseaService) AdditionBaseCollections() map[string]HardCodeCollection {
	result := make(map[string]HardCodeCollection)

	result["misato-frens"] = HardCodeCollection{
		Address: "0xccb6b629f5434102e37175bdac8262722180a62f",
		Chain:   "base",
		Slug:    "misato-frens",
		ChainID: 8453,
	}

	result["chonks"] = HardCodeCollection{
		Address: "0x07152bfde079b5319e5308c43fb1dbc9c76cb4f9",
		Chain:   "base",
		Slug:    "chonks",
		ChainID: 8453,
	}

	return result
}

func (s *OpenseaService) IsInAdditionalCollections(slug string) *HardCodeCollection {
	for k, i := range s.AdditionBaseCollections() {
		if strings.EqualFold(k, slug) {
			return &i
		}
	}

	return nil
}

func (s *OpenseaService) FindHardCodeCollectionByAddress(address string) *HardCodeCollection {
	for _, i := range s.AdditionBaseCollections() {
		if strings.EqualFold(i.Address, address) {
			return &i
		}
	}

	return nil
}

func NewOpensea(apiKey string) *OpenseaService {
	if apiKey == "" {
		apiKey = "b5400021ebc9489a9fe8a4544f663b53"
	}

	headers := make(map[string]string)
	headers["accept"] = "application/json"
	headers["x-api-Key"] = apiKey
	return &OpenseaService{
		apiKey:  apiKey,
		headers: headers,
	}
}

func (s *OpenseaService) GetProfileAvatar(addr string) (string, error) {
	fullUrl := fmt.Sprintf("%s/v1/user/%s", API_URL, addr)
	resp := User{}
	_bytes, _, _, err := helpers.HttpRequest(fullUrl, "GET", s.headers, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(_bytes, &resp)
	if err != nil {
		return "", err
	}

	return resp.Account.ProfileImgUrl, nil
}

func (s *OpenseaService) OpenseaGetContract(ctx context.Context, address string, chain string) (*OpenseaGetContract, error) {
	fullUrl := fmt.Sprintf("%s/v2/chain/%s/contract/%s", API_URL, chain, address)
	_bytes, _, _, err := helpers.HttpRequest(fullUrl, "GET", s.headers, nil)
	if err != nil {
		return nil, err
	}

	resp := &OpenseaGetContract{}
	err = json.Unmarshal(_bytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *OpenseaService) OpenseaGetCollections(ctx context.Context, f OpenSeaFilterCollections) (*CollectionsResp, error) {
	params := url.Values{}
	if f.Chain != "" {
		params.Set("chain", f.Chain)
	}
	if f.CreatorUsername != "" {
		params.Set("creator_username", f.CreatorUsername)
	}
	if f.IncludeHidden != nil {
		params.Set("include_hidden", fmt.Sprintf("%v", *f.IncludeHidden))
	}
	if f.Limit != 0 {
		params.Set("limit", fmt.Sprintf("%d", f.Limit))
	}
	if f.Next != "" {
		params.Set("next", f.Next)
	}
	if f.OrderBy != "" {
		params.Set("order_by", f.OrderBy)
	} else {
		params.Set("order_by", "created_date") //default: create date
	}

	fullUrl := fmt.Sprintf("%s/v2/collections", API_URL)
	if len(params) > 0 {
		fullUrl += "?" + params.Encode()
	}
	_bytes, _, _, err := helpers.HttpRequest(fullUrl, "GET", s.headers, nil)
	if err != nil {
		return nil, err
	}

	resp := &CollectionsResp{}
	err = json.Unmarshal(_bytes, resp)
	if err != nil {
		return nil, err
	}

	//add collections
	aCollections := s.AdditionBaseCollections()
	type additionalCollection struct {
		Err  error
		Data *SingleCollectionResp
		Slug string
	}

	out := make(chan additionalCollection, len(aCollections))
	for addCollection, _ := range aCollections {
		go func(addCollection string, out chan additionalCollection) {

			ad, addErr := s.OpenseaGetSingleCollection(ctx, addCollection)
			if addErr != nil {
				out <- additionalCollection{
					Err:  addErr,
					Data: nil,
					Slug: addCollection,
				}

				return
			}

			ad.Chain = "base"
			out <- additionalCollection{
				Err:  addErr,
				Data: ad,
				Slug: addCollection,
			}

		}(addCollection, out)

	}

	for range aCollections {
		data := <-out
		if data.Err != nil {
			fmt.Println(data.Slug + "--" + data.Err.Error())
			continue
		}

		if data.Data == nil {
			fmt.Println(data.Slug + "-- nil")
			continue
		}

		fmt.Println(data.Slug + "-- is appended")
		resp.Collections = append(resp.Collections, *data.Data)
	}

	return resp, nil
}

func (s *OpenseaService) OpenseaGetSingleCollection(ctx context.Context, collectionSlug string) (*SingleCollectionResp, error) {

	fullUrl := fmt.Sprintf("%s/v2/collections/%s", API_URL, collectionSlug)

	_bytes, _, _, err := helpers.HttpRequest(fullUrl, "GET", s.headers, nil)
	if err != nil {
		return nil, err
	}

	resp := &SingleCollectionResp{}
	err = json.Unmarshal(_bytes, resp)
	if err != nil {
		return nil, err
	}

	a := s.IsInAdditionalCollections(resp.Collection)
	if a != nil {
		resp.Chain = a.Chain
	}

	return resp, nil
}
