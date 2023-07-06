package main

import (
	"flag"
	"fmt"
	"ywadmin-v3/common/interceptor/rpcserver"
	"ywadmin-v3/service/yunwei/rpc/internal/cron"
	"ywadmin-v3/service/yunwei/rpc/jobs"
	"ywadmin-v3/service/yunwei/rpc/schedule"

	"ywadmin-v3/service/yunwei/rpc/internal/config"
	"ywadmin-v3/service/yunwei/rpc/internal/server"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/yunwei.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewYunWeiServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		yunweiclient.RegisterYunWeiServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddOptions(grpc.MaxRecvMsgSize(1024 * 1024 * 1024))

	defer s.Stop()
	//rpc日志记录全局中间件
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	//开启操作队列任务
	go jobs.InitJobs(ctx)
	//开启计划队列任务
	go schedule.InitSchedule(ctx)
	//启动日常定时计划任务
	task := cron.NewCronTask(ctx)
	go task.Start()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
