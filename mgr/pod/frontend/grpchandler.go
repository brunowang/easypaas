package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"context"
	"github.com/brunowang/easypaas/mgr/pod/dto"
	"github.com/brunowang/easypaas/mgr/pod/service"
	"github.com/brunowang/easypaas/pbgen/pod"
	"github.com/brunowang/gframe/gflog"
	"go.uber.org/zap"
	"time"
)

type grpcHandler struct {
	pod.UnimplementedPodServer
	svc service.PodService
}

func NewGrpcHandler(svc service.PodService) *grpcHandler {
	return &grpcHandler{svc: svc}
}

func (g *grpcHandler) AddPod(ctx context.Context, req *pod.PodInfo) (*pod.Response, error) {
	params := &dto.PodInfo{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler AddPod processing")
	nowt := time.Now()

	result, err := g.svc.AddPod(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler AddPod error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) DeletePod(ctx context.Context, req *pod.PodId) (*pod.Response, error) {
	params := &dto.PodId{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler DeletePod processing")
	nowt := time.Now()

	result, err := g.svc.DeletePod(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler DeletePod error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) FindPodByID(ctx context.Context, req *pod.PodId) (*pod.PodInfo, error) {
	params := &dto.PodId{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler FindPodByID processing")
	nowt := time.Now()

	result, err := g.svc.FindPodByID(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler FindPodByID error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) UpdatePod(ctx context.Context, req *pod.PodInfo) (*pod.Response, error) {
	params := &dto.PodInfo{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler UpdatePod processing")
	nowt := time.Now()

	result, err := g.svc.UpdatePod(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler UpdatePod error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) FindAllPod(ctx context.Context, req *pod.Empty) (*pod.AllPod, error) {
	params := &dto.Empty{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler FindAllPod processing")
	nowt := time.Now()

	result, err := g.svc.FindAllPod(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler FindAllPod error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}
