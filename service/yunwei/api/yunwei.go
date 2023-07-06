package main

import (
	"flag"
	"fmt"
	"ywadmin-v3/service/yunwei/api/internal/config"
	"ywadmin-v3/service/yunwei/api/internal/handler"
	"ywadmin-v3/service/yunwei/api/internal/middleware"
	"ywadmin-v3/service/yunwei/api/internal/reportLogDb"
	"ywadmin-v3/service/yunwei/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/yunwei.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件
	// 将nginx网关验证后的userId设置到ctx中
	server.Use(middleware.NewSetUidToCtxMiddleware(ctx).Handle)
	// 将操作记录日志中
	server.Use(reportLogDb.NewRecordOperationLogMiddleware(ctx).Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
