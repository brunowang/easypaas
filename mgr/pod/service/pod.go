package service

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"context"
	"github.com/brunowang/easypaas/mgr/pod/dto"
)

type PodService interface {
	AddPod(ctx context.Context, req *dto.PodInfo) (*dto.Response, error)

	DeletePod(ctx context.Context, req *dto.PodId) (*dto.Response, error)

	FindPodByID(ctx context.Context, req *dto.PodId) (*dto.PodInfo, error)

	UpdatePod(ctx context.Context, req *dto.PodInfo) (*dto.Response, error)

	FindAllPod(ctx context.Context, req *dto.Empty) (*dto.AllPod, error)
}

type PodServiceImpl struct{}

func NewPod() PodService {
	return &PodServiceImpl{}
}

func (s *PodServiceImpl) AddPod(ctx context.Context, req *dto.PodInfo) (*dto.Response, error) {
	return new(dto.Response), nil
}

func (s *PodServiceImpl) DeletePod(ctx context.Context, req *dto.PodId) (*dto.Response, error) {
	return new(dto.Response), nil
}

func (s *PodServiceImpl) FindPodByID(ctx context.Context, req *dto.PodId) (*dto.PodInfo, error) {
	return new(dto.PodInfo), nil
}

func (s *PodServiceImpl) UpdatePod(ctx context.Context, req *dto.PodInfo) (*dto.Response, error) {
	return new(dto.Response), nil
}

func (s *PodServiceImpl) FindAllPod(ctx context.Context, req *dto.Empty) (*dto.AllPod, error) {
	return new(dto.AllPod), nil
}
