package api

import (
	newErr "errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"ywadmin-v3/common/go-refresh-cdn/common/model"

	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliyunApi struct {
	*model.CommonApi
}

func (t *AliyunApi) DoApi() error {
	client, _err := t.CreateClient(tea.String(t.AccessKeyId), tea.String(t.AccessKeySecret))
	if _err != nil {
		return _err
	}
	var (
		tryErr error
	)

	if t.Action == "refresh" {
		urlList := strings.ReplaceAll(t.UrlList, ",", "\n")
		if t.PurgeType == "dirs" {
			refreshObjectCachesRequest := &cdn20180510.RefreshObjectCachesRequest{
				ObjectPath: tea.String(urlList),
				ObjectType: tea.String("Directory"),
			}
			runtime := &util.RuntimeOptions{}
			tryErr = func() (_e error) {
				defer func() {
					if r := tea.Recover(recover()); r != nil {
						_e = r
					}
				}()
				// 复制代码运行请自行打印 API 的返回值
				response, _err := client.RefreshObjectCachesWithOptions(refreshObjectCachesRequest, runtime)
				if _err != nil {
					return _err
				}
				fmt.Printf("%v", response.Body.String())

				return nil
			}()
		} else if t.PurgeType == "urls" {
			pushObjectCacheRequest := &cdn20180510.PushObjectCacheRequest{
				ObjectPath: tea.String(urlList),
			}
			runtime := &util.RuntimeOptions{}
			tryErr = func() (_e error) {
				defer func() {
					if r := tea.Recover(recover()); r != nil {
						_e = r
					}
				}()
				// 复制代码运行请自行打印 API 的返回值
				response, _err := client.PushObjectCacheWithOptions(pushObjectCacheRequest, runtime)
				if _err != nil {
					return _err
				}
				fmt.Printf("%v", response.Body.String())
				return nil
			}()

		}
	} else if t.Action == "query" {
		formatTemplate := "2006-01-02"
		startTime := time.Now().AddDate(0, 0, -2).Format(formatTemplate)
		endTime := time.Now().Format(formatTemplate)
		r := regexp.MustCompile(`http://|https://|/`)
		s := r.ReplaceAllString(t.UrlList, "")
		for _, v := range strings.Split(s, ",") {
			describeRefreshTasksRequest := &cdn20180510.DescribeRefreshTasksRequest{
				StartTime:  tea.String(startTime + "T00:00:00Z"),
				EndTime:    tea.String(endTime + "T23:59:59Z"),
				DomainName: tea.String(v),
				ObjectType: tea.String("directory"),
			}
			runtime := &util.RuntimeOptions{}
			tryErr = func() (_e error) {
				defer func() {
					if r := tea.Recover(recover()); r != nil {
						_e = r
					}
				}()
				// 复制代码运行请自行打印 API 的返回值
				response, _err := client.DescribeRefreshTasksWithOptions(describeRefreshTasksRequest, runtime)
				if _err != nil {
					return _err
				}
				fmt.Printf("%v\n", response.Body.String())

				return nil
			}()
		}

	} else if t.Action == "count" {
		describeCdnUserQuotaRequest := &cdn20180510.DescribeCdnUserQuotaRequest{}
		runtime := &util.RuntimeOptions{}
		tryErr = func() (_e error) {
			defer func() {
				if r := tea.Recover(recover()); r != nil {
					_e = r
				}
			}()
			// 复制代码运行请自行打印 API 的返回值
			response, _err := client.DescribeCdnUserQuotaWithOptions(describeCdnUserQuotaRequest, runtime)
			if _err != nil {
				return _err
			}
			fmt.Printf("%v\n", response.Body.String())

			return nil
		}()
	} else if t.Action == "ips" {
		if t.Ips == "" {
			return newErr.New("缺少-i ip1,ip2参数")
		}
		for _, v := range strings.Split(t.Ips, ",") {
			fmt.Println("IP:", v)
			describeIpInfoRequest := &cdn20180510.DescribeIpInfoRequest{IP: tea.String(v)}
			runtime := &util.RuntimeOptions{}
			tryErr = func() (_e error) {
				defer func() {
					if r := tea.Recover(recover()); r != nil {
						_e = r
					}
				}()
				// 复制代码运行请自行打印 API 的返回值
				response, _err := client.DescribeIpInfoWithOptions(describeIpInfoRequest, runtime)
				if _err != nil {
					return _err
				}
				fmt.Printf("%v\n", response.Body.String())
				return nil
			}()
		}
		return nil
	}
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_result, _err := util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
		fmt.Printf("%#v\n", *_result)
	}
	return _err
}

func (t *AliyunApi) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result = &cdn20180510.Client{}
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}
