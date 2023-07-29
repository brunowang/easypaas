package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/easypaas/mgr/pod/service"
	"github.com/brunowang/gframe/gflog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
)

type httpServer struct {
	h *httpHandler
}

func NewHttpServer(svc service.PodService) *httpServer {
	return &httpServer{
		h: NewHttpHandler(svc),
	}
}

func (s *httpServer) Serve(httpl net.Listener) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gflog.TraceMiddleware())

	gflog.Info(nil, "Http start listen", zap.String("address", httpl.Addr().String()))

	return router.RunListener(httpl)
}
