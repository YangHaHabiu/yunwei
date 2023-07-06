package tencentLibs

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
)

//获取公共客户端请求参数
func GetClientRequest(t *RequestModel) (*tke.Client, error) {
	//传入参数
	credential := common.NewCredential(
		t.SecretId,
		t.SecretKey,
	)

	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tke.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, err := tke.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		return nil, err
	}
	return client, nil

}
