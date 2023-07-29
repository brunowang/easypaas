package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"fmt"
	"github.com/brunowang/easypaas/mgr/pod/service"
	"github.com/brunowang/gframe/gflog"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"net"
)

type Frontend interface {
	Serve(l net.Listener) error
}

type FeWrapper struct {
	fe Frontend
	ln net.Listener
}

type server struct {
	l   net.Listener
	m   cmux.CMux
	fes []FeWrapper
	svc service.PodService
}

func MustNewServer(address string) *server {
	l, err := net.Listen("tcp", address)
	if err != nil {
		gflog.Fatal(nil, fmt.Sprintf("listening on %s failed", address), zap.Error(err))
	} else {
		gflog.Info(nil, "server start listen", zap.String("address", address))
	}
	svc := service.NewPod()
	return &server{l: l, m: cmux.New(l), svc: svc}
}

func (s *server) SetupGRPC() *server {
	grpcSrv := NewGrpcServer(s.svc)
	grpcl := s.m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	s.AddFrontend(grpcSrv, grpcl)
	return s
}

func (s *server) SetupHTTP() *server {
	httpSrv := NewHttpServer(s.svc)
	httpl := s.m.Match(cmux.HTTP1Fast())
	s.AddFrontend(httpSrv, httpl)
	return s
}

func (s *server) AddFrontend(fe Frontend, l net.Listener) *server {
	s.fes = append(s.fes, FeWrapper{fe: fe, ln: l})
	return s
}

func (s *server) Run() error {
	for _, wrap := range s.fes {
		go wrap.fe.Serve(wrap.ln)
	}
	return s.m.Serve()
}
