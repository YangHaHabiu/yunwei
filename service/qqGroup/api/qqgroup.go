package main

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/service/qqGroup/api/internal/config"
	"ywadmin-v3/service/qqGroup/api/internal/handler"
	"ywadmin-v3/service/qqGroup/api/internal/logic/cron"
	"ywadmin-v3/service/qqGroup/api/internal/svc"
)

var configFile = flag.String("f", "etc/qqgroup.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	//设置全局的qq负载的值
	all, _ := ctx.QqLoadBalanceModel.FindMasterByFilters()
	for _, v := range *all {
		if v.GroupType == "group" {
			globalkey.QqGroupList = append(globalkey.QqGroupList, v.Qq)
		}
		if v.GroupType == "discuss" {
			globalkey.QqDiscussList = append(globalkey.QqDiscussList, v.Qq)
		}
		if v.IsMaster == 1 {
			globalkey.QqMsgKey[v.GroupType+"_qq"] = v.Qq
			globalkey.QqMsgKey[v.GroupType+"_qqapi"] = v.QqApi
		}

	}
	//设置全局qq管理员
	for _, v := range strings.Split(c.Project.ManagerQQ, ",") {
		splitx := strings.Split(v, ":")
		if len(splitx) == 2 {
			globalkey.QqDefaltManger[gconv.Int64(splitx[0])] = splitx[1]
		}
	}
	//启动计划任务
	task := cron.NewCronTask(ctx)
	go task.Start()
	fmt.Printf("Starting server at %s:%d,Version %s...\n", c.Host, c.Port, globalkey.QqManagerVersion)
	server.Start()
}
