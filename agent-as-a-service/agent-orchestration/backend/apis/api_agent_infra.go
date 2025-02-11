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

func (s *Server) proxyAgentStoreMiddleware(prefixPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		StoreId := s.stringFromContextParam(c, "Store_id")
		// apiKey := c.GetHeader("api-key")
		Store, err := s.nls.GetAgentStore(context.Background(), StoreId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
			return
		}
		if Store.Status != models.AgentStoreStatusActived {
			c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
			return
		}
		// agentStoreInstall, err := s.nls.ValidateUserStoreFee(context.Background(), apiKey)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		// 	return
		// }
		var urlPath string
		director := func(req *http.Request) {
			hostURL, err := url.Parse(Store.ApiUrl)
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
			prefixPath = prefixPath + "/" + StoreId
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
		// _ = s.nls.ChargeUserStoreInstall(context.Background(), agentStoreInstall.ID, urlPath, w.Status())
	}
}

func (s *Server) ScanAgentInfraMintHash(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	networkID, err := s.uint64FromContextQuery(c, "network_id")
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err = s.nls.ScanAgentInfraMintHash(
		ctx,
		userAddress,
		networkID,
		s.stringFromContextQuery(c, "tx_hash"),
		s.uintFromContextParam(c, "agent_store_id"),
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}
