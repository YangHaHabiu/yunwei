package main

import (
	"flag"
	"fmt"
	"ywadmin-v3/service/sendMsg/api/internal/logic/jobs"
	"ywadmin-v3/service/sendMsg/api/internal/middleware"

	"ywadmin-v3/service/sendMsg/api/internal/config"
	"ywadmin-v3/service/sendMsg/api/internal/handler"
	"ywadmin-v3/service/sendMsg/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/sendmsg.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// 将nginx网关验证后的userId设置到ctx中
	server.Use(middleware.NewSetUidToCtxMiddleware(ctx).Handle)
	handler.RegisterHandlers(server, ctx)
	//开启消息计划任务
	go jobs.RunSendMsgJobs(ctx)
	//开启提醒任务
	go jobs.RemindJobs(ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
