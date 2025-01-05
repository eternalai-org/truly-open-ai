package serializers

import "github.com/eternalai-org/eternal-ai/agent-orchestration/core/domain"

// JSONResponse ...
type JSONResponse struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}

// ResponseSuccess ...
func ResponseSuccess(data interface{}) JSONResponse {
	return JSONResponse{
		Result: data,
	}
}

// ResponseError ...
func ResponseError(err error) JSONResponse {
	switch err.(type) {
	case *domain.Error:
		{
		}
	default:
		{
			err = domain.ErrorWithMessage(domain.ErrSystemError, err.Error())
		}
	}
	return JSONResponse{
		Error: err,
	}
}
