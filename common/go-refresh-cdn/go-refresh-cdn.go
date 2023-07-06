package main

import (
	"flag"
	"fmt"
	"os"
	"ywadmin-v3/common/go-refresh-cdn/api"
	"ywadmin-v3/common/go-refresh-cdn/common"
	"ywadmin-v3/common/go-refresh-cdn/common/model"
)

var (
	cdnTypes  string
	cdnAction string
	KeyAuth   string
	PurgeType string
	Ips       string
)

func init() { // 每个文件会自动执行的函数
	flag.StringVar(&cdnTypes, "t", "tencent", "输入cdn供应商")
	flag.StringVar(&cdnAction, "a", "query", "输入cdn执行操作")
	flag.StringVar(&KeyAuth, "k", "accessKey", "输入key验证类型")
	flag.StringVar(&PurgeType, "p", "dirs", "输入刷新类型")
	flag.StringVar(&Ips, "i", "ips", "输入节点ip地址")
}

func help() {
	fmt.Printf(`其他参数说明：
-t cdn供应商（tencent（默认），wangsu，aliyun）
-a 操作（query（默认）:查询cdn刷新历史，refresh:刷新cdn缓存，count:统计刷新用量配额，ips:查询ip归属）
-k 验证类型（accessKey（默认）:sdk方式认证，apiKey:api方式认证）
-p 刷新类型（dirs（默认）:目录刷新，urls:url刷新）
-i 查询节点ip地址归属，多个用逗号分隔
`)
	os.Exit(1)
}

func main() {

	flag.Parse()
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	urlList := os.Getenv("urlList")
	if accessKeyId == "" || accessKeySecret == "" || urlList == "" {
		fmt.Println("当前刷新CDN版本号：", common.Version)
		fmt.Println("缺少环境变量参数：accessKeyId , accessKeySecret , urlList")
		help()
	}
	//操作类型区分
	switch cdnAction {
	case "refresh":
		fmt.Println("当前使用[", cdnTypes, "]供应商[", cdnAction, "]操作类型-刷新CDN的缓存操作")
	case "query":
		fmt.Println("当前使用[", cdnTypes, "]供应商[", cdnAction, "]操作类型-查询CDN的刷新操作记录操作")
	case "count":
		fmt.Println("当前使用[", cdnTypes, "]供应商[", cdnAction, "]操作类型-查询CDN的刷新用量配额操作")
	case "ips":
		fmt.Println("当前使用[", cdnTypes, "]供应商[", cdnAction, "]操作类型-查询IP归属操作")
	default:
		help()
	}
	//认证方式类型
	switch KeyAuth {
	case "accessKey":
		fmt.Println("当前使用[", KeyAuth, "]的sdk认证方式")
	case "apiKey":
		fmt.Println("当前使用[", KeyAuth, "]的api认证方式")
	default:
		help()
	}
	//刷新方式类型
	switch PurgeType {
	case "dirs":
		fmt.Println("当前使用[", PurgeType, "]的目录刷新方式")
	case "urls":
		fmt.Println("当前使用[", PurgeType, "]的url刷新方式")
	default:
		help()
	}

	//新建公共结构体参数
	objectCommon := &model.CommonApi{
		AccessKeySecret: accessKeySecret,
		AccessKeyId:     accessKeyId,
		UrlList:         urlList,
		Action:          cdnAction,
		KeyAuth:         KeyAuth,
		PurgeType:       PurgeType,
		Ips:             Ips,
	}

	//区分供应商
	switch cdnTypes {
	case "tencent":
		api.RefreshCdn(&api.TencentApi{objectCommon})
	case "wangsu":
		api.RefreshCdn(&api.WangSuApi{objectCommon})
	case "aliyun":
		api.RefreshCdn(&api.AliyunApi{objectCommon})
	default:
		help()
	}
}
