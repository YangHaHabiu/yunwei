/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: sendMessage
* @Date: 2021-7-15 15:13
 */
package send_message

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type message struct {
	Title       string
	Sender      string
	Content     string
	StartTime   string
	EndTime     string
	ExecTime    string
	TaskStatus  int64
	ProjectCn   string
	TaskId      int64
	TaskProcess string
	TaskType    int64
	Notifier    string
	Url         string
}

//qq消息结构体
type qqMessage struct {
	Retcode int         `json:"retcode"`
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Error   interface{} `json:"error,omitempty"`
}

// 构建发送qq消息的接口
type MessageInterface interface {
	Send(groupId, msgType string) string
	SendJf(groupId, msgType string) string
	SendHot(groupId, msgType string) string
}

// 构造函数
func NewMessage(title, sender, content, startTime, endTime,
	execTime, projectCn, taskProcess string,
	taskstatus, taskId, taskType int64,
	notifier, url string) *message {
	return &message{
		Title:       title,
		Sender:      sender,
		Content:     content,
		StartTime:   startTime,
		EndTime:     endTime,
		ExecTime:    execTime,
		TaskStatus:  taskstatus,
		ProjectCn:   projectCn,
		TaskId:      taskId,
		TaskProcess: taskProcess,
		TaskType:    taskType,
		Notifier:    notifier,
		Url:         url,
	}
}

var taskMap = map[int64]string{
	0: "等待开始",
	1: "正在执行",
	2: "执行失败",
	3: "执行成功",
	4: "已取消",
}
var taskTypeMap = map[int64]string{
	1: "临时维护",
	2: "日常维护",
}

var taskHotMap = map[int64]string{
	1: "操作完成",
	2: "操作失败",
}

// 发送消息
// discuss group
func (m *message) Send(groupId, msgType string) string {
	//var bq string
	body := `%s任务状态：%s
任务类型：%s
操作者：%s
项目：%s
任务ID：%d
任务标题：%s
任务开始时间：%s
任务执行时间：%s
任务结束时间：%s
操作内容：
%s
%s
`
	if groupId == "324113338" {
		m.Notifier = ""
	}
	content := fmt.Sprintf(body, m.Notifier,
		taskMap[m.TaskStatus],
		taskTypeMap[m.TaskType],
		m.Sender, m.ProjectCn,
		m.TaskId, m.Title,
		m.StartTime, m.ExecTime,
		m.EndTime, strings.ReplaceAll(m.Content, "==>", "\n"), m.TaskProcess)
	content = strings.ReplaceAll(content, "f447b20a7fcbf53a5d5be013ea0b15af", "\"")
	keys := fmt.Sprintf("%s_id", msgType)
	data := map[string]string{
		keys:          groupId,
		"message":     content,
		"auto_escape": "false",
	}
	return m.postData(data, msgType, m.Url)

}

//热更通知
func (m *message) SendHot(groupId, msgType string) string {
	body := `%s
快捷热更：%s
操作项目：%s
日志信息：
==============================
%s
==============================
`
	content := fmt.Sprintf(body, m.Sender, taskHotMap[m.TaskStatus], m.ProjectCn, m.Content)
	//替换引号
	keys := fmt.Sprintf("%s_id", msgType)
	data := map[string]string{
		keys:          groupId,
		"message":     content,
		"auto_escape": "false",
	}
	return m.postData(data, msgType, m.Url)
}

//特殊处理jf
func (m *message) SendJf(groupId, msgType string) string {
	body := `%s任务状态：%s
任务类型：%s
操作者：%s
项目：%s
任务ID：%d
任务标题：%s
操作内容：
%s
`
	if groupId == "324113338" {
		m.Notifier = ""
	}
	content := fmt.Sprintf(body, m.Notifier, taskMap[m.TaskStatus], taskTypeMap[m.TaskType], m.Sender, m.ProjectCn, m.TaskId, m.Title, m.Content)
	//替换引号
	content = strings.ReplaceAll(content, "f447b20a7fcbf53a5d5be013ea0b15af", "\"")
	keys := fmt.Sprintf("%s_id", msgType)
	data := map[string]string{
		keys:          groupId,
		"message":     content,
		"auto_escape": "false",
	}
	return m.postData(data, msgType, m.Url)
}

func (m *message) postData(data map[string]string, msgType, url string) string {
	//var url string
	//all, err := qqMsg.SelectListAll(&qqMsg.SelectPageReq{
	//	GroupType: "group",
	//})
	//if err != nil || len(all) != 1 {
	//	fmt.Println("get qq database error", err)
	//	url = fmt.Sprintf("%s/send_%s_msg", cfg.Instance().Admin.QqAddr, msgType)
	//} else {
	//	url = fmt.Sprintf("%s/send_%s_msg", all[0].QqApi, msgType)
	//}

	_, allList := readLineType(data["message"], 70)

	for i, v := range allList {

		data["message"] = strings.Join(v, "\n")
		if i == len(allList)-1 {
			data["message"] = fmt.Sprintf("%s\n#####################\n本次通知消息共%d段，已发送完毕", data["message"], len(allList))
		}
		b, _ := json.Marshal(data)
		resp, err := http.Post(url,
			"application/json",
			bytes.NewBuffer(b))
		if err != nil {
			fmt.Println("error => ", err)
			return ""
		}
		defer resp.Body.Close()
		time.Sleep(5 * time.Second)
		//_, err = ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	return ""
		//}
		//return string(body)

	}
	return ""
}

//按照指定行数进行读取进行返回
func readLineType(read string, nums int) (error, [][]string) {
	//将字符串删除空行
	//re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	//read = re.ReplaceAllString(read, "")
	//定义几个数组
	newList := make([][]string, 0)
	lineList := make([]string, 0)
	buf := bufio.NewReader(bytes.NewBufferString(read))
	n := 1
	for {
		//匹配换行符
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			lineList = append(lineList, line)
		}
		if err != nil {
			if err == io.EOF {
				newList = append(newList, lineList)
				return nil, newList
			}
			return err, nil
		}
		//根据定义的行数加入数组中
		if n%nums == 0 {
			newList = append(newList, lineList)
			lineList = make([]string, 0)
		}
		n++
	}
}
