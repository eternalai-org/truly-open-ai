package btcapi

import (
	"fmt"
	"net/http"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

func (c *Client) buildMempoolUrl(resourcePath string) string {
	if resourcePath != "" {
		return c.MempoolUrl + "/" + resourcePath
	}
	return c.MempoolUrl
}

type MiningPoolResp struct {
	PoolID       uint             `json:"poolId"`
	Name         string           `json:"name"`
	Link         string           `json:"link"`
	BlockCount   int              `json:"blockCount"`
	Rank         int              `json:"rank"`
	EmptyBlocks  int              `json:"emptyBlocks"`
	Slug         string           `json:"slug"`
	AvgMatchRate numeric.BigFloat `json:"avgMatchRate"`
	AvgFeeDelta  numeric.BigFloat `json:"avgFeeDelta"`
	PoolUniqueID uint             `json:"poolUniqueId"`
}

func (c *Client) GetMiningPools() ([]*MiningPoolResp, error) {
	resp := struct {
		Pools []*MiningPoolResp
	}{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildMempoolUrl("api/v1/mining/pools/3m"),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Pools, nil
}

type BlockAuditScoreResp struct {
	Hash           string           `json:"hash"`
	MatchRate      numeric.BigFloat `json:"matchRate"`
	ExpectedFees   int64            `json:"expectedFees"`
	ExpectedWeight int64            `json:"expectedWeight"`
}

func (c *Client) GetBlockAuditScore(hash string) (*BlockAuditScoreResp, error) {
	resp := BlockAuditScoreResp{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildMempoolUrl(fmt.Sprintf("api/v1/mining/blocks/audit/score/%s", hash)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetBlocksAuditScore(startHeight uint64) ([]*BlockAuditScoreResp, error) {
	resp := []*BlockAuditScoreResp{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildMempoolUrl(fmt.Sprintf("api/v1/mining/blocks/audit/scores/%d", startHeight)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type MiningBlockResp struct {
	Id string `json:"id"`
}

func (c *Client) GetMiningPoolBlocks(slug string) ([]*MiningBlockResp, error) {
	resp := []*MiningBlockResp{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildMempoolUrl(fmt.Sprintf("api/v1/mining/pool/%s/blocks", slug)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type BlockResp struct {
	Id     string `json:"id"`
	Extras struct {
		MatchRate      numeric.BigFloat `json:"matchRate"`
		ExpectedFees   int64            `json:"expectedFees"`
		ExpectedWeight int64            `json:"expectedWeight"`
		Pool           struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"pool"`
	} `json:"extras"`
}

func (c *Client) GetBlock(hash string) (*BlockResp, error) {
	resp := BlockResp{}
	err := c.methodJSON(
		http.MethodGet,
		c.buildMempoolUrl(fmt.Sprintf("api/v1/block/%s", hash)),
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
