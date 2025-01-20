package hiro

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
)

type HiroService struct {
	url     string
	apiKey  string //b5400021ebc9489a9fe8a4544f663b53
	headers map[string]string
}

func NewHiroService(baseUrl string) *HiroService {
	obj := &HiroService{
		url:     baseUrl,
		headers: make(map[string]string),
	}
	obj.headers["Accept"] = "*/*"
	return obj
}

func (s *HiroService) GetInscriptionInfo(inscriptionId string) (*HiroInscriptionInfo, error) {
	url := fmt.Sprintf(s.url+"/ordinals/v1/inscriptions/%s", inscriptionId)
	respBytes, _, httpStatus, err := helpers.HttpRequest(url, "GET", s.headers, nil)
	if err != nil {
		return nil, err
	}
	if httpStatus != 200 {
		return nil, errors.New("invalid")
	}
	resp := &HiroInscriptionInfo{}
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
