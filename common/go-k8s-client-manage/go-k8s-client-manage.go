package main

import (
	"flag"
	"fmt"
	"os"
	"ywadmin-v3/common/go-k8s-client-manage/comm"
	"ywadmin-v3/common/go-k8s-client-manage/libs"

	"github.com/go-playground/validator/v10"
)

func init() {
	flag.StringVar(&comm.Actions, "a", "", "输入action操作")
	flag.StringVar(&comm.FileConfig, "f", "", "输入fileconfig文件")
	flag.StringVar(&comm.Works, "w", "", "输入工作负载类型")
	flag.StringVar(&comm.YmalFile, "y", "", "输入yaml文件")
	flag.StringVar(&comm.NameSpace, "n", "", "输入命名空间")
	flag.StringVar(&comm.ServiceName, "s", "", "输入服务名称")
}

func main() {
	flag.Parse()
	object := comm.InputArg{
		Actions:     comm.Actions,
		FileConfig:  comm.FileConfig,
		Works:       comm.Works,
		YamlFile:    comm.YmalFile,
		NameSpace:   comm.NameSpace,
		ServiceName: comm.ServiceName,
	}
	validate := validator.New()
	err := validate.Struct(object)
	errinfo := comm.ProcessErr(object, err)
	if len(errinfo) != 0 {
		comm.Help()
		fmt.Println(errinfo)
		os.Exit(1)
	}

	//获取客户端请求
	c, err := libs.ClientSetFunc(object.FileConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//根据工作负载进行对应操作
	switch object.Works {
	case "deployment":
		//操作无状态部署
		err := libs.OperationK8sClient(&libs.DeploymentMsg{
			&libs.ClientSet{
				Clientset:   c,
				NameSpace:   object.NameSpace,
				ServiceName: object.ServiceName,
				Works:       object.Works,
			},
		}, object.Actions)
		if err != nil {
			os.Exit(3)
		}
	case "statefullset":
		//操作有状态部署
		err := libs.OperationK8sClient(&libs.StatefullsetMsg{
			&libs.ClientSet{
				Clientset:   c,
				NameSpace:   object.NameSpace,
				ServiceName: object.ServiceName,
				Works:       object.Works,
			},
		}, object.Actions)
		if err != nil {
			os.Exit(4)
		}
	}

}
