package tencentLibs

import (
	"encoding/json"
	"fmt"
	"ywadmin-v3/common/tool"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
)

type ClusterMsg struct {
	*RequestModel
}

func (t *ClusterMsg) Query() error {
	client, err := GetClientRequest(t.RequestModel)
	if err != nil {
		return err
	}
	request := tke.NewDescribeClustersRequest()
	response, err := client.DescribeClusters(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	fmt.Println(tool.JsonFormatOut(response.ToJsonString()))

	return nil
}

func (t *ClusterMsg) Update() error {
	return nil
}

func (t *ClusterMsg) Create() error {
	client, err := GetClientRequest(t.RequestModel)
	if err != nil {
		return err
	}
	//获取kubeconfig文件
	request := tke.NewDescribeClusterKubeconfigRequest()
	request.ClusterId = common.StringPtr("cls-m9u3jzp0")
	response, err := client.DescribeClusterKubeconfig(request)
	//根据集群id获取集群管理员角色
	//request := tke.NewAcquireClusterAdminRoleRequest()
	//request.ClusterId = common.StringPtr("cls-m9u3jzp0")
	//response, err := client.AcquireClusterAdminRole(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	var tmp KubeConfigResponse
	err = json.Unmarshal([]byte(response.ToJsonString()), &tmp)
	if err != nil {
		return err
	}
	err = tool.WriteFile("./kubeconfig", tmp.Response.Kubeconfig)
	if err != nil {
		return err
	}
	return nil
}
