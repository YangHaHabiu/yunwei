package schedule

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"

	"github.com/gogf/gf/util/gconv"
	"github.com/zeromicro/go-zero/core/logx"
)

//初始化计划任务
func InitSchedule(svx *svc.ServiceContext) {
	for {
		//go ScheduleCrond(svx)
		ScheduleCrond(svx)
		time.Sleep(time.Second)
	}

}

type typeMap struct {
	Ids      []string `json:"ids"`
	Rangeids []string `json:"rangeids"`
}

//计划任务管理
func ScheduleCrond(svx *svc.ServiceContext) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tsq, err := svx.TasksScheduleQueueModel.FindAll(ctx, "schedule_status__=", "1", "schedule_start_time__expect", 99)
	if err != nil {
		logx.Error("查询计划队列任务失败了，原因：" + err.Error())
		return
	}
	dictList, err := svx.AdminRpc.DictList(ctx, &adminclient.DictListReq{
		Pid:      -2,
		Types:    "schedule_types",
		Current:  0,
		PageSize: 0,
	})
	if err != nil {
		logx.Error("查询计划队列字典失败了，原因：" + err.Error())
		return
	}

	tmp := make(map[string]typeMap, 0)
	for _, v := range dictList.List {
		t := new(typeMap)
		tmp[v.Value] = *t
	}
	if len(*tsq) > 0 {
		for k1 := range tmp {
			t := new(typeMap)
			t.Ids = make([]string, 0)
			t.Rangeids = make([]string, 0)
			for _, v2 := range *tsq {
				if k1 == v2.ScheduleType {
					t.Ids = append(tmp[k1].Ids, gconv.String(v2.Id))
					t.Rangeids = append(tmp[k1].Rangeids, strings.Split(v2.ScheduleRangeIds, ",")...)
				}
			}
			tmp[k1] = *t
		}
		logx.Error("封装计划队列字典信息：", tmp)
		for k1, v1 := range tmp {
			fmt.Println(k1)
			fmt.Println(v1)
			if k1 == "1" && len(v1.Ids) > 0 {
				logx.Error("待装计划...")
				installPath := filepath.Dir(svx.Config.Scripts.InstallFilePath)
				lock := installPath + "/installx.lock"
				if tool.IsExist(lock) {
					logx.Error("存在装服的锁文件")
					return
				}
				file, err := os.Create(lock)
				if err != nil {
					logx.Error("创建锁文件失败")
					return
				}
				defer file.Close()

				//修改所有的状态的,防止下一分钟又重复执行
				result := make([]string, 0)
				err = svx.TasksScheduleQueueModel.BatchUpdateStatus(ctx, strings.Join(v1.Ids, ","), "", 2, 0)
				if err != nil {
					logx.Error("批量修改状态失败了，原因：" + err.Error())
					os.Remove(lock)
					return
				}
				//开始执行任务
				viol, err := svx.TasksScheduleQueueModel.FindAllViewInstallOpenList(ctx, "view_open_plan_autoid__in", strings.Join(v1.Rangeids, ","), "view_install_status__=", "2")
				if err != nil {
					logx.Error("查询安装列表视图失败了，原因：" + err.Error())
					os.Remove(lock)
					return
				}
				if len(*viol) > 0 {
					for _, v := range *viol {
						result = append(result, v.ViewInstallList)
					}
					str := strings.Join(result, "\n")

					err = ioutil.WriteFile(filepath.Join(installPath, "Tmp/tmp/list.txt"), []byte(str+"\n"), 0666)
					if err != nil {
						os.Remove(lock)
						return
					}
					//执行脚本
					scriptCmd := fmt.Sprintf("cd %s;sh main.sh -w WEB -s all", installPath)
					go func(v1 typeMap) {
						ctx, cancel := context.WithCancel(context.Background())
						defer cancel()
						scheduleStatus := 3
						job := xcmd.NewCommandJob(10*time.Minute, scriptCmd)
						msgFilePath := fmt.Sprintf("%s/Log/once_run_main.log", installPath)
						content, _ := ioutil.ReadFile(msgFilePath)
						outContent := "==========装服执行成功，路径：" + scriptCmd + "==========\n" + string(content)
						if !job.IsOk {
							scheduleStatus = 4
							outContent = fmt.Sprintf("==========装服执行失败==========\n%s", string(content))
						}
						if job.IsTimeout {
							scheduleStatus = 4
							outContent = fmt.Sprintf("==========装服脚本执行超时==========")
						}

						//b, _ := json.Marshal([]byte(outContent))

						err := svx.TasksScheduleQueueModel.BatchUpdateStatus(ctx, strings.Join(v1.Ids, ","), outContent, scheduleStatus, time.Now().Unix())
						if err != nil {
							logx.Error("修改计划队列状态失败")
							os.Remove(lock)
							return
						}
						logx.Error(outContent)
						os.Remove(lock)
						if scheduleStatus == 4 {
							//修改开服计划状态
							err := svx.OpenPlanModel.BatchUpdateStatus(ctx, strings.Join(v1.Rangeids, ","), -1)
							if err != nil {
								logx.Error("修改开服计划状态失败", v1.Ids)
							}
						}
					}(v1)
				} else {
					err := svx.TasksScheduleQueueModel.BatchUpdateStatus(ctx, strings.Join(v1.Ids, ","), "查询装服视图计划任务列表为空", 4, time.Now().Unix())
					if err != nil {
						logx.Error("修改计划队列状态失败")
						os.Remove(lock)
						return
					}
					os.Remove(lock)
				}
			} else if k1 == "2" && len(v1.Ids) > 0 {
				logx.Error("待合计划...")
				//直接修改状态成功，其他操作待定
				err := svx.TasksScheduleQueueModel.BatchUpdateStatus(ctx, strings.Join(v1.Ids, ","), "", 3, time.Now().Unix())
				if err != nil {
					logx.Error("修改计划队列状态失败")
					return
				}
			}
		}
	}

}
