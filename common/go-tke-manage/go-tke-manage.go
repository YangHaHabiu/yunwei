package main

import (
	"ywadmin-v3/common/go-tke-manage/tencentLibs"
)

func main() {

	secretId := ""
	secretKey := ""
	object := &tencentLibs.RequestModel{
		SecretId:  secretId,
		SecretKey: secretKey,
	}

	//查询集群
	//tencentLibs.OperationFunc(&tencentLibs.ClusterMsg{object}, "query")
	//生成集群配置
	//tencentLibs.OperationFunc(&tencentLibs.ClusterMsg{object}, "create")
	//查询容器
	tencentLibs.OperationFunc(&tencentLibs.ContainerMsg{object}, "query")
	//更新容器
	//tencentLibs.OperationFunc(&tencentLibs.ContainerMsg{object}, "update")

}
