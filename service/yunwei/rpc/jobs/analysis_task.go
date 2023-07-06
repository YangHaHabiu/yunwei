/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: analysis_task
* @Date: 2021-8-31 17:10
 */
package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
)

func AnalysisDataByTaskId(id int64, svx *svc.ServiceContext, ctx context.Context) string {
	var statusValues string
	result := make([]string, 0)

	firstObj, _ := svx.TasksModel.FindAll(ctx, "pid__=", id)

	for _, v2 := range *firstObj {
		result = append(result, "------------\n -"+v2.Name)
		//查询层级2的数据
		secondObj, _ := svx.TasksModel.FindAll(ctx, "pid__=", v2.Id)
		//查询层级3的数据
		for _, v3 := range *secondObj {
			if v3.TaskStatus == 5 || v3.TaskStatus == 6 {
				continue
			}
			if v3.TaskStatus == 2 {
				statusValues = "[失败]"
			} else if v3.TaskStatus == 3 {
				statusValues = "[成功]"
			} else if v3.TaskStatus == 1 {
				statusValues = "[执行中]"
			} else if v3.TaskStatus == 0 {
				statusValues = "[未开始]"
			}
			result = append(result, "  --"+v3.Name+" "+statusValues)
			compile, _ := regexp.Compile("热更配置|执行服务端命令")
			if compile.MatchString(v3.Name) {
				var taskCmd TaskCommonJson
				err := json.Unmarshal([]byte(v3.Cmd), &taskCmd)
				if err != nil {
					return ""
				}
				if v3.Types == "server_hot" {
					content := fmt.Sprintf("     推送到发布%s服\n%s", taskCmd.Stable, strings.Join(contentpredadd(strings.Split(taskCmd.FileList, ",")), "\n"))
					result = append(result, content)
				} else if v3.Types == "server_cmd" {
					result = append(result, strings.Join(contentpredadd(strings.Split(taskCmd.FileList, ",")), "\n"))
				}
				if v3.TaskStatus == 2 {
					return strings.Join(result, "\n")
				}
			}

		}
	}
	return strings.Join(result, "\n")

}

func contentpredadd(src []string) (desc []string) {
	desc = make([]string, 0)
	for _, v := range src {
		desc = append(desc, "     "+v)
	}
	return
}

type TaskCommonJson struct {
	//Id            int64  `json:"id,omitempty"`
	Operation      string `json:"operation,omitempty"`
	OperationCn    string `json:"operationCn,omitempty"`
	Stable         string `json:"stable,omitempty"`
	OuterIp        string `json:"outerIp,omitempty"`
	DbUpdate       string `json:"dbUpdate,omitempty"`
	FileList       string `json:"fileList,omitempty"`
	MaintainRange  string `json:"maintainRange,omitempty"`
	CmdList        string `json:"cmdList,omitempty"`
	DbType         string `json:"dbType,omitempty"`
	PlatName       string `json:"platName,omitempty"`
	CheckSt        string `json:"checkSt,omitempty"`
	SQLCmd         string `json:"SQLCmd,omitempty"`
	Merge          string `json:"merge,omitempty"`
	ExportFileName string `json:"exportFileName,omitempty"`
	InitSetTime    string `json:"initSetTime,omitempty"`
	ExecuteSQL     string `json:"executeSQL,omitempty"`
	ExecuteFlag    string `json:"executeFlag,omitempty"`
}
