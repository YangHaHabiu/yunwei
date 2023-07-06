package api

import (
	"bytes"
	"encoding/json"
	newErr "errors"
	"fmt"
	"strings"
	"time"
	"ywadmin-v3/common/go-refresh-cdn/common/model"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type TencentApi struct {
	*model.CommonApi
}

func (t *TencentApi) DoApi() error {

	credential := common.NewCredential(
		t.AccessKeyId,
		t.AccessKeySecret,
	)
	urlList := strings.Split(t.UrlList, ",")
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的

	client, err := cdn.NewClient(credential, "", cpf)
	if err != nil {
		return err
	}
	if t.Action == "query" {
		request := cdn.NewDescribePurgeTasksRequest()
		formatTemplate := "2006-01-02"
		startTime := time.Now().AddDate(0, 0, -2).Format(formatTemplate)
		endTime := time.Now().Format(formatTemplate)
		request.StartTime = common.StringPtr(startTime + " 00:00:00")
		request.EndTime = common.StringPtr(endTime + " 23:59:59")
		for _, v := range urlList {
			request.Keyword = common.StringPtr(v)
			response, err := client.DescribePurgeTasks(request)
			if _, ok := err.(*errors.TencentCloudSDKError); ok {
				return err
			}
			if err != nil {
				return err
			}
			jsonFormatOut(response.ToJsonString())
		}

	} else if t.Action == "refresh" {
		var (
			result string
		)
		// 实例化一个请求对象,每个接口都会对应一个request对象
		//request := cdn.NewPurgeUrlsCacheRequest()
		//request.Urls = common.StringPtrs(urlList)
		if t.PurgeType == "dirs" {
			request := cdn.NewPurgePathCacheRequest()
			request.Paths = common.StringPtrs(urlList)
			request.FlushType = common.StringPtr("flush")
			// 返回的resp是一个PurgeUrlsCacheResponse的实例，与请求对象对应
			response, err := client.PurgePathCache(request)
			if _, ok := err.(*errors.TencentCloudSDKError); ok {
				return err
			}
			if err != nil {
				return err
			}
			result = response.ToJsonString()
		} else if t.PurgeType == "urls" {
			request := cdn.NewPurgeUrlsCacheRequest()
			request.Urls = common.StringPtrs(urlList)
			response, err := client.PurgeUrlsCache(request)
			if _, ok := err.(*errors.TencentCloudSDKError); ok {
				return err
			}
			if err != nil {
				return err
			}
			result = response.ToJsonString()
		}

		jsonFormatOut(result)
	} else if t.Action == "count" {
		request := cdn.NewDescribePurgeQuotaRequest()
		// 返回的resp是一个PurgeUrlsCacheResponse的实例，与请求对象对应
		response, err := client.DescribePurgeQuota(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			return err
		}
		if err != nil {
			return err
		}
		jsonFormatOut(response.ToJsonString())
	} else if t.Action == "ips" {
		if t.Ips == "" {
			return newErr.New("缺少-i ip1,ip2参数")
		}
		request := cdn.NewDescribeCdnIpRequest()
		request.Ips = common.StringPtrs(strings.Split(t.Ips, ","))
		response, err := client.DescribeCdnIp(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			return err
		}
		if err != nil {
			return err
		}
		jsonFormatOut(response.ToJsonString())

	}
	return err
}

// 输出json格式的字符串回包
func jsonFormatOut(result string) {
	var str bytes.Buffer
	err := json.Indent(&str, []byte(result), "", "    ")
	if err != nil {
		fmt.Println(result)
		return
	}
	fmt.Printf("%s\n", str.String())
}
