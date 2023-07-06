/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: job
* @Date: 2021-6-28 17:00
 */
package jobs

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gogf/gf/util/gconv"
	"log"
	"math"
	"math/big"
	"os/exec"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"syscall"
	"time"
	"ywadmin-v3/common/send_message"
	"ywadmin-v3/common/xtime"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/model"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
)

type JobResult struct {
	OutMsg    string
	ErrMsg    string
	IsOk      bool
	IsTimeout bool
}

func NewJobFromTask(svx *svc.ServiceContext) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	taskFirstLevel, err := svx.TasksModel.FindAll(ctx, "pid__in", "-1,0", "task_status__=", "-1", "task_start_time__expect", 99)
	if err != nil {
		log.Println("主任务查询失败...", err)
		return
	}
	if len(*taskFirstLevel) > 0 {
		chs := make([]chan string, len(*taskFirstLevel))
		limitFunc := func(chLimit chan bool, ch chan string, taskId int64, timeout time.Duration) {
			Run(svx, ctx, taskId, timeout, ch)
			<-chLimit
		}
		startTime := time.Now()
		log.Println("批量任务执行开始.本次任务数量：", len(*taskFirstLevel))
		for i, v1 := range *taskFirstLevel {
			chs[i] = make(chan string, 1)
			chLimit <- true
			go limitFunc(chLimit, chs[i], v1.Id, 24*time.Hour)
		}
		for _, ch := range chs {
			log.Println(<-ch)
		}
		endTime := time.Now()
		log.Printf("批量任务执行完毕.总花费：%s，本次任务数量：%d\n", endTime.Sub(startTime), len(*taskFirstLevel))
	} else {
		time.Sleep(2 * time.Second)
		return
	}
}

//本地执行命令
func NewCommandJob(svx *svc.ServiceContext, ctx context.Context, timeout time.Duration, command, types string, taskId int64) *JobResult {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("CMD", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	//将pid和pgid设置相同的进程
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()
	if types == "updates" {
		//先删后加
		err := svx.TasksTidPidModel.Delete(ctx, taskId)
		if err != nil {
			log.Println("删除关联的taskid失败")
		}
		_, err = svx.TasksTidPidModel.Insert(ctx, &model.TasksTidPid{
			Tid: taskId,
			Pid: int64(cmd.Process.Pid),
		})
		if err != nil {
			log.Println("插入进程关联PID失败" + err.Error())
		}

	}
	err, isTimeout := runCmdWithTimeout(cmd, timeout)
	jobresult := new(JobResult)
	jobresult.OutMsg = bufOut.String()
	jobresult.ErrMsg = bufErr.String()
	jobresult.IsOk = true
	if err != nil {
		jobresult.ErrMsg = err.Error()
		jobresult.IsOk = false
	}
	jobresult.IsTimeout = isTimeout

	return jobresult
}

func run(svx *svc.ServiceContext, ctx context.Context, taskId int64, ch chan string) {
	startTimes := time.Now().Unix()
	entity, err := svx.TasksModel.FindOne(ctx, taskId)
	if err != nil {
		log.Println("查询任务失败，原因：" + err.Error())
		ch <- fmt.Sprintf("task id %d , query error!!", taskId)
		return

	}
	svx.TasksModel.UpdateByField(ctx, []string{
		"task_exec_time",
		"task_status",
		"id",
	}, time.Now().Unix(), 1, taskId)

	if entity.TaskStatus == 3 {
		ch <- fmt.Sprintf("task id %d , 1task completed!!", taskId)
		return
	}

	firstjobs, err := svx.TasksModel.FindAll(ctx, "pid__=", taskId)
	if err != nil {
		ch <- fmt.Sprintf("task id %d , 2query firstjob error!!", taskId)
		return
	}
	var (
		operator       string
		exportfilename string
		userId         int64
		projectObj     *model.Project
		userObj        *adminclient.UserListData
		qqGroupObj     model.QqGroupList
		senders        string
		sendUrl        string
	)

	//获取基础信息
	projectObj, err = svx.TasksModel.FindProjectById(ctx, entity.ProjectId)
	if err != nil {
		ch <- fmt.Sprintf("task id %d , 3task completed!!!!", entity.Id)
		return
	}

	//获取qq发送的url
	filters, err := svx.TasksTidPidModel.FindListByMasterId(ctx, "group_type__=", "group", "is_master__=", "1")

	if err != nil || len(*filters) != 1 {
		ch <- fmt.Sprintf("task id %d , 4task completed!!", taskId)
		return
	}
	qqGroupObj = (*filters)[0]
	sendUrl = fmt.Sprintf("%s/send_group_msg", qqGroupObj.QqApi)

	if entity.UpdateBy != "0" {
		userId = gconv.Int64(entity.UpdateBy)
	} else {
		userId = gconv.Int64(entity.CreateBy)
	}
	name, err := svx.AdminRpc.UserList(ctx, &adminclient.UserListReq{
		UserId: userId,
	})
	if err != nil || len(name.List) != 1 {
		ch <- fmt.Sprintf("task id %d , 5task completed!!", taskId)
		return
	}
	userObj = name.List[0]
	splitUser := strings.Split(userObj.Email, "@")
	var msg send_message.MessageInterface
	//只有维护计划发送开始任务的qq消息
	if projectObj.GroupQq != "" {
		if entity.TaskType == "2" {
			msg = send_message.NewMessage(
				entity.Name,
				operator,
				entity.Content,
				xtime.GetTimetampByTimeMinu(entity.TaskStartTime),
				"",
				xtime.GetTimetampByTimeMinu(time.Now().Unix()),
				projectObj.ProjectCn,
				"",
				1,
				entity.Id,
				gconv.Int64(entity.TaskType),
				"",
				sendUrl)
			msg.Send(projectObj.GroupQq, projectObj.GroupType)
		}
	}
	operator = fmt.Sprintf("%s-%s", userObj.NickName, userObj.Name)
	flagone := false
	for _, v1 := range *firstjobs {

		log.Println("一级任务信息：", v1)
		if v1.TaskStatus == 3 {
			continue
		}
		secondjobs, err := svx.TasksModel.FindAll(ctx, "pid__=", v1.Id)
		if err != nil {
			ch <- fmt.Sprintf("task id %d , query secondjob error!!", v1.Id)
			return
		}

		entity2, err := svx.TasksModel.FindOne(ctx, v1.Id)
		if err != nil {
			ch <- fmt.Sprintf("task id %d , query error!!", v1.Id)
			return
		}
		//secondStep := 0
		flag := false
		for _, v2 := range *secondjobs {
			log.Printf("主任务ID [%d] -- 二级任务信息：%v\n", taskId, v2)
			// 跳过已完成的任务
			if v2.TaskStatus >= 3 {
				continue
			}

			if err != nil {
				ch <- fmt.Sprintf("task id %d , query thirdjob error!!", v2.Id)
				return
			}
			//游戏名 操作类型 孙子任务id json
			//sh $0 -g GAME_NAME -o OPERATE_TYPE -t TASK_ID -j JSON_INFO

			//执行任务开始
			cmds := fmt.Sprintf(`source /etc/profile;bash %s -g %s -o %s -t %d -u %s -j '%s'`, svx.Config.Scripts.MaintainFilePath, projectObj.ProjectEn, v2.Types, v2.Id, operator, v2.Cmd)
			//修改v2的记录
			entity3, _ := svx.TasksModel.FindOne(ctx, v2.Id)
			//修改v2和v3的任务状态为进行中
			svx.TasksModel.UpdateByField(ctx, []string{
				"task_status",
				"id",
			}, 1, entity.Id)
			svx.TasksModel.UpdateByField(ctx, []string{
				"task_status",
				"id",
			}, 1, entity2.Id)
			svx.TasksModel.UpdateByField(ctx, []string{
				"task_status",
				"id",
			}, 1, entity3.Id)

			//修改v3的记录
			jobOut := NewCommandJob(svx, ctx, 24*time.Hour, cmds, "updates", v2.Id)
			var (
				status     int64
				result     string
				timeoutStr string
			)
			result = jobOut.ErrMsg + "\n" + jobOut.OutMsg
			if jobOut.IsOk {
				status = 3
			} else {
				status = 2

			}
			if jobOut.IsTimeout {
				status = 2
				timeoutStr = "执行任务超时\n"
			}
			//写入热更日志
			hotLogHistory := new(model.HotLogHistory)
			hotLogHistory.CreateBy = gconv.String(userId)
			hotLogHistory.OperType = entity3.Types
			compile := regexp.MustCompile(`\(.*\)`)
			hotLogHistory.HotTitle = fmt.Sprintf("%s-%s-%s", projectObj.ProjectEn, compile.ReplaceAllString(entity3.Name, ""), time.Now().Format("150405"))
			hotLogHistory.ProjectId = projectObj.ProjectId
			//任务出错退出
			if status == 2 {
				times := time.Now().Unix()
				svx.TasksModel.UpdateByField(ctx, []string{
					"task_status",
					"task_end_time",
					"id",
				}, 2, times, entity.Id)

				svx.TasksModel.UpdateByField(ctx, []string{
					"task_status",
					"task_end_time",
					"id",
				}, 2, times, entity2.Id)

				svx.TasksModel.UpdateByField(ctx, []string{
					"task_status",
					"task_end_time",
					"id",
				}, 2, times, entity3.Id)

				//出错写入日志
				svx.TaskLogHistroyModel.Insert(ctx, &model.TaskLogHistroy{
					TasksId:   entity3.Id,
					TasksTime: time.Now().Unix(),
					TasksLogs: timeoutStr + result,
				})

				ids, _ := svx.TasksModel.FindOne(ctx, entity3.Id)
				compile, _ := regexp.Compile(`\(.*\)`)
				allString := compile.ReplaceAllString(ids.Name, "")
				process := fmt.Sprintf("========================\n出错的任务进度：%s-%s", entity2.Name, allString)
				fmt.Println(process)
				hotLogHistory.HotTitle = fmt.Sprintf("[ERROR]%s-%s-%s", projectObj.ProjectEn, compile.ReplaceAllString(entity3.Name, ""), time.Now().Format("150405"))
				hotLogHistory.OperStatus = 1
				hotLogHistory.OperContent = timeoutStr + result
				_, err = svx.HotLogHistoryModel.Insert(ctx, hotLogHistory, "")
				if err != nil {
					fmt.Println(err)

				}

				//依次发送qq消息开始
				var msg send_message.MessageInterface
				if projectObj.GroupQq != "" {
					if entity.TaskType == "2" {
						maintainObj, _ := svx.MaintainPlanModel.FindAll(ctx, "task_id__=", entity.Id)
						name, _ = svx.AdminRpc.UserList(ctx, &adminclient.UserListReq{
							UserId: gconv.Int64((*maintainObj)[0].CreateBy),
						})
						splitUser := strings.Split(name.List[0].Email, "@")
						senders += fmt.Sprintf("[CQ:at,qq=%s]\n", splitUser[0])
						msg = send_message.NewMessage(entity.Name,
							operator,
							entity.Content,
							xtime.GetTimetampByTimeMinu(entity.TaskStartTime),
							xtime.GetTimetampByTimeMinu(time.Now().Unix()),
							xtime.GetTimetampByTimeMinu(startTimes),
							projectObj.ProjectCn,
							process,
							2,
							entity.Id,
							gconv.Int64(entity.TaskType),
							senders,
							sendUrl,
						)
						msg.Send(projectObj.GroupQq, projectObj.GroupType)
					} else {
						senders += fmt.Sprintf("[CQ:at,qq=%s]\n", splitUser[0])
						msg = send_message.NewMessage(entity.Name,
							operator,
							AnalysisDataByTaskId(entity.Id, svx, ctx),
							"",
							"",
							"",
							projectObj.ProjectCn,
							"",
							2,
							entity.Id,
							gconv.Int64(entity.TaskType),
							senders,
							sendUrl)
						msg.SendJf(projectObj.GroupQq, projectObj.GroupType)
					}
				}
				msg = send_message.NewMessage(entity.Name,
					operator,
					entity.Content,
					xtime.GetTimetampByTimeMinu(entity.TaskStartTime),
					xtime.GetTimetampByTimeMinu(time.Now().Unix()),
					xtime.GetTimetampByTimeMinu(startTimes),
					projectObj.ProjectCn,
					process,
					2,
					entity.Id,
					gconv.Int64(entity.TaskType),
					senders,
					sendUrl)

				ch <- fmt.Sprintf("task id %d error stop3 !!!!", taskId)
				return
			}
			//成功写入导出文件名称
			if v2.Types == "export_sql" || v2.Types == "exec_sql_file" || v2.Types == "server_cmd" {
				var taskListObj model.TaskCommonJson
				json.Unmarshal([]byte(v2.Cmd), &taskListObj)
				if taskListObj.Merge != "3c9e656d1f13e0fc2cbc3af658e20243" {
					if taskListObj.Merge == "true" || taskListObj.Merge == "false" {
						exportfilename = taskListObj.ExportFileName
					}
				}
			}
			if v2.Types == "get_server_conf" {
				exportfilename = "server_info.txt"

			}
			svx.TasksModel.UpdateByField(ctx, []string{
				"task_status",
				"task_end_time",
				"id",
			}, status, time.Now().Unix(), entity3.Id)

			svx.TaskLogHistroyModel.Insert(ctx, &model.TaskLogHistroy{
				TasksId:   entity3.Id,
				TasksTime: time.Now().Unix(),
				TasksLogs: result,
			})
			hotLogHistory.OperStatus = 0
			hotLogHistory.OperContent = result
			_, err = svx.HotLogHistoryModel.Insert(ctx, hotLogHistory, "")
			if err != nil {
				fmt.Println(err)
			}
			if status != 2 {
			} else {
				flag = true
			}
		}
		if flag {
			entity2.TaskStatus = 2
			flagone = true
		} else {
			entity2.TaskStatus = 3
		}
		svx.TasksModel.UpdateByField(ctx, []string{
			"task_end_time",
			"task_status",
			"task_step",
			"id",
		}, entity2.TaskEndTime, entity2.TaskStatus, entity2.TaskStep, entity2.Id)
	}
	if flagone {
		entity.TaskStatus = 2
	} else {
		entity.TaskStatus = 3
	}
	entity.TaskEndTime = time.Now().Unix()
	if exportfilename != "" {
		entity.ExportFileName = exportfilename
	}
	svx.TasksModel.UpdateByField(ctx, []string{
		"task_end_time",
		"task_status",
		"task_step",
		"export_file_name",
		"id",
	}, entity.TaskEndTime, entity.TaskStatus, entity.TaskStep, entity.ExportFileName, entity.Id)

	//依次发送qq消息开始
	if projectObj.GroupQq != "" {
		if entity.TaskType == "2" {
			maintainObj, _ := svx.MaintainPlanModel.FindAll(ctx, "task_id__=", entity.Id)
			name, _ = svx.AdminRpc.UserList(ctx, &adminclient.UserListReq{
				UserId: gconv.Int64((*maintainObj)[0].CreateBy),
			})
			splitUser := strings.Split(name.List[0].Email, "@")
			senders = fmt.Sprintf("[CQ:at,qq=%s]\n", splitUser[0])
			msg = send_message.NewMessage(entity.Name,
				operator,
				entity.Content,
				xtime.GetTimetampByTimeMinu(entity.TaskStartTime),
				xtime.GetTimetampByTimeMinu(time.Now().Unix()),
				xtime.GetTimetampByTimeMinu(startTimes),
				projectObj.ProjectCn,
				"",
				3,
				entity.Id,
				gconv.Int64(entity.TaskType),
				senders,
				sendUrl)
			msg.Send(projectObj.GroupQq, projectObj.GroupType)
		} else {
			senders = fmt.Sprintf("[CQ:at,qq=%s]\n", splitUser[0])
			msg = send_message.NewMessage(entity.Name,
				operator,
				AnalysisDataByTaskId(entity.Id, svx, ctx),
				"",
				"",
				"",
				projectObj.ProjectCn,
				"",
				3,
				entity.Id,
				gconv.Int64(entity.TaskType),
				senders,
				sendUrl)
			msg.SendJf(projectObj.GroupQq, projectObj.GroupType)
		}
	}

	//成功后插入到临时维护计划
	//split := strings.Split(entity.Content, "==>")
	//newsList := make([]string, 0)
	//for _, v := range split {
	//	if strings.Contains(v, "前端") {
	//		continue
	//	}
	//	i := strings.Split(v, ">")
	//	for k1, v1 := range i {
	//		if k1 == 0 {
	//			continue
	//		}
	//		r := regexp.MustCompile(`\((.*?)\)`)
	//		res := r.FindStringSubmatch(v1)
	//		result := strings.TrimRight(res[1], " ")
	//		i2 := strings.Split(result, ",")
	//		for _, v := range i2 {
	//			v = strings.TrimRight(v, " ")
	//			v = strings.ReplaceAll(v, " ", "_")
	//
	//			newsList = append(newsList, v)
	//		}
	//
	//	}
	//}

	//属于日常维护才插入一条到维护计划
	//if entity.TaskType == "1" {
	//svx.MaintainPlanModel.Insert(ctx, &model.MaintainPlan{
	//	ProjectId:     entity.ProjectId,
	//	MaintainType:  "2",
	//	StartTime:     gconv.String(entity.TaskStartTime),
	//	EndTime:       gconv.String(entity.TaskEndTime),
	//	MaintainRange: strings.ReplaceAll(strings.Join(removeDup(newsList), ","), "_", " "),
	//	Title:         "临时维护",
	//	Content:       strings.ReplaceAll(entity.Content, "==>", "\n"),
	//	CreateBy:      entity.CreateBy,
	//	ClusterId:     entity.ClusterId,
	//})
	//}

	ch <- fmt.Sprintf("task id %d , exec successd!!", taskId)
	return
}

//去重列表
func removeDup(a []string) (ret []string) {
	sort.Strings(a)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func Run(svx *svc.ServiceContext, ctx context.Context, taskId int64, timeout time.Duration, ch chan string) {

	//entity := &updateModel.Entity{Id: taskId}
	//entity.FindOne()
	//entity.TaskStatus = 1
	//entity.Update([]string{
	//	"task_status",
	//}...)

	err := svx.TasksModel.UpdateByField(ctx, []string{
		"task_status", "id",
	}, 1, taskId)
	if err != nil {
		log.Println("修改任务状态为进行中失败" + err.Error())
		return
	}

	ch_run := make(chan string)
	go run(svx, ctx, taskId, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(timeout + (10 * time.Minute)):
		re := fmt.Sprintf("task id %d , timeout is %ds", taskId, timeout/time.Second)
		ch <- re

	}
}

// 手动停止任务
func Stop(svx *svc.ServiceContext, ctx context.Context, taskId int64) error {
	//根据主id查询正在执行子任务id--tasks表
	tobjct, err := svx.TasksModel.SelectRunningTaskById(ctx, taskId)
	if err != nil || tobjct == nil {
		return errors.New(fmt.Sprintf("查询正在运行任务失败，%v %v", err, tobjct))
	}
	//根据子任务id查询pid进程--tasks_pid表
	taskspidObj, err := svx.TasksTidPidModel.FindOne(ctx, tobjct.Id)
	if err != nil {
		return errors.New(fmt.Sprintf("查询对应进程id-->%d 对应进程pid-->%d错误，%v", tobjct.Id, taskspidObj.Pid, err))
	}

	//杀掉子pid进程（1）
	//command := fmt.Sprintf("kill %d", taskspidObj.Pid)
	//cmd := exec.Command("/bin/bash", "-c", command)
	//if err = cmd.Run(); err != nil {
	//	log.Println("执行命令失败，原因是：", err)
	//	return err
	//}
	//杀掉子pid进程组（2）
	if err = syscall.Kill(-int(taskspidObj.Pid), syscall.SIGTERM); err != nil {
		return errors.New(fmt.Sprintf("杀进程失败，进程ID-->%d，失败原因：%v", taskspidObj.Pid, err))
	}
	//修改对应的状态为失败，一并修改主任务状态为失败--tasks表
	if err = svx.TasksModel.UpdateByField(ctx, []string{
		"task_status",
		"task_end_time",
		"id",
	}, 2, time.Now().Unix(), taskId); err != nil {
		return err
	}

	if err = svx.TasksModel.UpdateByField(ctx, []string{
		"task_status",
		"task_end_time",
		"id",
	}, 2, time.Now().Unix(), taskspidObj.Id); err != nil {
		return err
	}

	return nil
}

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
