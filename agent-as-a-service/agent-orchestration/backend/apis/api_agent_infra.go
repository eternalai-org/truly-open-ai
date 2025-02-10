package apis

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) proxyAgentInfraMiddleware(prefixPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		infraId := s.stringFromContextParam(c, "infra_id")
		// apiKey := c.GetHeader("api-key")
		infra, err := s.nls.GetAgentInfra(context.Background(), infraId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		if infra.Status != models.AgentInfraStatusActived {
			c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
			return
		}
		// agentInfraInstall, err := s.nls.ValidateUserInfraFee(context.Background(), apiKey)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		// 	return
		// }
		var urlPath string
		director := func(req *http.Request) {
			hostURL, err := url.Parse(infra.ApiUrl)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
				return
			}
			if hostURL.Scheme != "https" {
				c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
				return
			}
			if !strings.Contains(hostURL.Host, ".") {
				c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
				return
			}
			prefixPath = prefixPath + "/" + infraId
			req.URL.Scheme = hostURL.Scheme
			req.URL.Host = hostURL.Host
			req.Host = hostURL.Host
			subPath := strings.TrimPrefix(req.URL.Path, prefixPath)
			req.URL.Path = strings.TrimSuffix(hostURL.Path, "/") + "/" + strings.TrimPrefix(subPath, "/")
			query := req.URL.Query()
			req.URL.RawQuery = query.Encode()
			for k := range r.Header {
				v := c.GetHeader(k)
				req.Header.Set(k, v)
			}
			urlPath = req.URL.String()
			if os.Getenv("DEV") == "true" {
				fmt.Printf("%s -> %s\n", r.URL.String(), urlPath)
			}
		}
		proxy := &httputil.ReverseProxy{
			Director: director,
		}
		proxy.ServeHTTP(w, r)
		// _ = s.nls.ChargeUserInfraInstall(context.Background(), agentInfraInstall.ID, urlPath, w.Status())
	}
}

func (s *Server) CreateOrUpdateAgentInfra(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AgentInfraReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	obj, err := s.nls.CreateOrUpdateAgentInfra(ctx, userAddress, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfraResp(obj)})
}

func (s *Server) GetListAgentInfra(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	res, count, err := s.nls.GetListAgentInfra(ctx, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfraRespArray(res), Count: &count})
}

func (s *Server) GetListAgentInfraByUser(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	res, count, err := s.nls.GetListAgentInfraByUser(ctx, userAddress, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfraRespArray(res), Count: &count})
}

func (s *Server) GetAgentInfraDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	res, err := s.nls.GetAgentInfraDetail(ctx, s.uintFromContextParam(c, "id"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfraResp(res)})
}

func (s *Server) GetAgentInfraInstallCode(c *gin.Context) {
	ctx := s.requestContext(c)
	agentInfraID := s.uintFromContextParam(c, "id")
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	res, err := s.nls.CreateAgentInfraInstallCode(ctx, userAddress, agentInfraID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: res.Code})
}

func (s *Server) GetListAgentInfraInstallByUser(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	res, count, err := s.nls.GetListAgentInfraInstallByUser(ctx, userAddress, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentInfraInstallRespArray(res), Count: &count})
}
