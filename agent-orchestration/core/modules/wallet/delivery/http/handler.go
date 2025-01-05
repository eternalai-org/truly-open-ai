package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/domain"
	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/serializers"
	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/server"
)

// WalletHandler  represent the httphandler for wallet
type WalletHandler struct {
	server        *server.HTTPServer
	walletUsecase domain.IWalletUsecase
}

// NewWalletHandler ...
func NewWalletHandler(s *server.HTTPServer, uu domain.IWalletUsecase) {
	handler := &WalletHandler{
		server:        s,
		walletUsecase: uu,
	}
	wallet := handler.server.Engine().Group("wallet")
	wallet.POST("/update", handler.GenerateUpdate)
	wallet.POST("/get-private-key", handler.GetPrivateKey)
}

// Generate ...
func (h *WalletHandler) GenerateUpdate(c *gin.Context) {
	var req domain.WalletUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.ResponseError(domain.ErrInvalidArgument))
		return
	}
	ctx := h.server.GetContext(c)
	wallet, err := h.walletUsecase.GenerateUpdate(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.ResponseError(err))
		return
	}
	for a, k := range req.Batch {
		_, err := h.walletUsecase.GenerateUpdate(ctx, domain.WalletUpdateRequest{
			WalletId:   req.WalletId,
			WalletType: req.WalletType,
			Address:    a,
			PrivateKey: k,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, serializers.ResponseError(err))
			return
		}
	}
	c.JSON(http.StatusOK, serializers.ResponseSuccess(wallet))
}

// GetPrivateKey ...
func (h *WalletHandler) GetPrivateKey(c *gin.Context) {
	var req domain.WalletGetPrivateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.ResponseError(domain.ErrInvalidArgument))
		return
	}
	ctx := h.server.GetContext(c)
	privateKey, err := h.walletUsecase.GetPrivateKey(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.ResponseError(err))
		return
	}
	c.JSON(http.StatusOK, serializers.ResponseSuccess(privateKey))
}
