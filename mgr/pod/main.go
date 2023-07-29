package main

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/easypaas/mgr/pod/conf"
	"github.com/brunowang/easypaas/mgr/pod/frontend"
	"github.com/brunowang/gframe/gflog"
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	app    = kingpin.New("pod", "pod server")
	config = app.Flag("config", "config file path").Required().String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	runtime.GOMAXPROCS(runtime.NumCPU())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	if err := conf.BindToml(*config); err != nil {
		gflog.Fatal(nil, "config bind toml failed", zap.Error(err))
	}

	go frontend.MustNewServer(":80").SetupGRPC().SetupHTTP().Run()

	<-sigs
}
