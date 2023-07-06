package tencentLibs

import (
	newErr "errors"
	"fmt"
	"ywadmin-v3/common/tool"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
)

type ContainerMsg struct {
	*RequestModel
}

func (t *ContainerMsg) Query() error {
	client, err := GetClientRequest(t.RequestModel)
	if err != nil {
		return err
	}

	//查询容器
	request := tke.NewDescribeEKSContainerInstancesRequest()
	response, err := client.DescribeEKSContainerInstances(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	fmt.Println(tool.JsonFormatOut(response.ToJsonString()))

	return nil
}

func (t *ContainerMsg) Update() error {
	if t.RequestModel.EksCiId == "" {
		return newErr.New("缺少eksciid")
	}
	client, err := GetClientRequest(t.RequestModel)
	if err != nil {
		return err
	}

	request := tke.NewUpdateEKSContainerInstanceRequest()
	request.EksCiId = common.StringPtr(t.EksCiId)
	response, err := client.UpdateEKSContainerInstance(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	fmt.Println(tool.JsonFormatOut(response.ToJsonString()))
	return nil
}

func (t *ContainerMsg) Create() error {
	return nil
}
