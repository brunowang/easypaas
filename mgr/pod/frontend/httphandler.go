package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/easypaas/mgr/pod/dto"
	"github.com/brunowang/easypaas/mgr/pod/service"
	"github.com/brunowang/gframe/gfhttp"
	"github.com/brunowang/gframe/gflog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type httpHandler struct {
	svc service.PodService
}

func NewHttpHandler(svc service.PodService) *httpHandler {
	return &httpHandler{svc: svc}
}

func (s *httpHandler) AddPod(ctx *gin.Context) {
	var req dto.PodInfo
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler AddPod processing")
	nowt := time.Now()

	rsp, err := s.svc.AddPod(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler AddPod error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler AddPod finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) DeletePod(ctx *gin.Context) {
	var req dto.PodId
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler DeletePod processing")
	nowt := time.Now()

	rsp, err := s.svc.DeletePod(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler DeletePod error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler DeletePod finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) FindPodByID(ctx *gin.Context) {
	var req dto.PodId
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler FindPodByID processing")
	nowt := time.Now()

	rsp, err := s.svc.FindPodByID(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler FindPodByID error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler FindPodByID finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) UpdatePod(ctx *gin.Context) {
	var req dto.PodInfo
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler UpdatePod processing")
	nowt := time.Now()

	rsp, err := s.svc.UpdatePod(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler UpdatePod error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler UpdatePod finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) FindAllPod(ctx *gin.Context) {
	var req dto.Empty
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler FindAllPod processing")
	nowt := time.Now()

	rsp, err := s.svc.FindAllPod(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler FindAllPod error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler FindAllPod finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}
