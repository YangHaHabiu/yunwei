package main

import (
	"flag"
	"fmt"

	"ywadmin-v3/service/monitor/rpc/internal/config"
	"ywadmin-v3/service/monitor/rpc/internal/server"
	"ywadmin-v3/service/monitor/rpc/internal/svc"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/monitor.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		monitorclient.RegisterMonitorServer(grpcServer, server.NewMonitorServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
