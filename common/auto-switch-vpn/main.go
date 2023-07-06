package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
	"ywadmin-v3/common/crawler"
	"ywadmin-v3/common/qqmanger/group"

	"github.com/robfig/cron/v3"
)

type AikuaiResp struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
	Data   struct {
		Total int `json:"total"`
		Data  []struct {
			Interface string `json:"interface"`
			Name      string `json:"name"`
			Timestamp int    `json:"timestamp"`
		} `json:"data"`
		Interface [][]string `json:"interface"`
	}
}
type NewAikuaiResp struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
	Data   struct {
		IfaceCheck []struct {
			ID              int    `json:"id"`
			Interface       string `json:"interface"`
			ParentInterface string `json:"parent_interface"`
			IPAddr          string `json:"ip_addr"`
			Gateway         string `json:"gateway"`
			Internet        string `json:"internet"`
			Updatetime      string `json:"updatetime"`
			AutoSwitch      string `json:"auto_switch"`
			Result          string `json:"result"`
			Errmsg          string `json:"errmsg"`
			Comment         string `json:"comment"`
		} `json:"iface_check"`
	} `json:"Data"`
}

// 爱快登录函数
func AikuaiLoginFunc(url string) (string, error) {
	datas := `{
		"username": "admin",
		"passwd": "c63e88cbfc879f3610a220a36f34dd3f",
		"pass": "c2FsdF8xMXlteGgoMTIzKQ==",
		"remember_password": "true"
	  }`
	err, s := crawler.PostNew(url+"/Action/login", datas, "", true)
	if err != nil {
		return "", err
	}
	var tmp AikuaiResp
	header := strings.Split(s, "@@")[1]
	dataHeader := strings.Split(s, "@@")[0]
	json.Unmarshal([]byte(dataHeader), &tmp)
	if tmp.ErrMsg != "Succeess" {
		return "", err
	}
	setCookie := "username=admin; login=1; " + header
	return setCookie, nil
}

// 切换线路
func AikuaiSwitchLineFunc(url, setCookie string) (string, error) {
	datas := `{
		"func_name": "l2tp_client",
		"action": "show",
		"param": {
		  "TYPE": "total,data,interface"
		}
	  }`
	err, s2 := crawler.PostNew(url+"/Action/call", datas, setCookie, false)
	if err != nil {
		return "", err
	}
	var tmp2 AikuaiResp
	json.Unmarshal([]byte(s2), &tmp2)
	if tmp2.ErrMsg != "Success" {
		return "", errors.New("切换失败")
	}

	lineList := make([]string, 0)
	for _, v := range tmp2.Data.Interface {
		if !group.Contains([]string{"wan1", "auto", "vwan100", "adsl5"}, v[0]) {
			lineList = append(lineList, v[0])
		}
	}
	index := 0
	for i, v := range lineList {
		if v == tmp2.Data.Data[1].Interface {
			index = i
		}
	}
	index = (index + 1) % len(lineList)
	datas = fmt.Sprintf(`
	{
		"func_name": "l2tp_client",
		"action": "edit",
		"param": {
		  "name": "l2tp_tfgame2",
		  "comment": "",
		  "server": "office.gztfgame.com",
		  "gateway": "172.30.30.1",
		  "server_port": 1701,
		  "username": "heystone",
		  "passwd": "yw@dqw.com",
		  "ipsec_secret": "TFgame",
		  "interface": "%s",
		  "leftid": "",
		  "rightid": "",
		  "mru": 1400,
		  "timing_rst_switch": 0,
		  "timing_rst_week": "1234567",
		  "timing_rst_time": "00:00,,",
		  "cycle_rst_time": 0,
		  "qos_switch": 0,
		  "updatetime": %d,
		  "dns2": "",
		  "dns1": "",
		  "mppe": "",
		  "ip_addr": "172.30.30.4",
		  "id": 2,
		  "enabled": "yes",
		  "mtu": 1400,
		  "week": "1234567",
		  "mon9": 0,
		  "date1": "00:00",
		  "date2": "",
		  "date3": ""
		}
	  }
	`, lineList[index], time.Now().Unix())
	err, s3 := crawler.PostNew(url+"/Action/call", datas, setCookie, false)
	if err != nil {
		return "", err
	}
	var tmp3 AikuaiResp
	json.Unmarshal([]byte(s3), &tmp3)
	if tmp3.ErrMsg != "Success" {
		return "", errors.New("切换失败")
	}
	return fmt.Sprintf("VPN线路由原%s，切换到%s", tmp2.Data.Data[1].Interface, lineList[index]), nil
}

// 利用日志判断切换线路的条件之一
func AikuaiJudgeLine(url, setCookie string) bool {
	datas := `{
		"func_name": "syslog-wanpppoe",
		"action": "show",
		"param": {
		  "TYPE": "total,data,interface",
		  "ORDER_BY": "id",
		  "ORDER": "desc",
		  "limit": "0,20",
		  "FINDS": "content",
		  "KEYWORDS": "%s",
		  "FILTER1": "interface,==,l2tp_tfgame2"
		}
	  }`
	var noramlTime, exitTime int
	for i, v := range []string{"Connect:%20l2tp_tfgame2", "Exit."} {
		data := fmt.Sprintf(datas, v)
		err, s2 := crawler.PostNew(url+"/Action/call", data, setCookie, false)
		if err != nil {
			return false
		}
		var tmp2 AikuaiResp
		json.Unmarshal([]byte(s2), &tmp2)
		if tmp2.ErrMsg != "Success" {
			return false
		}
		time.Sleep(time.Second)
		if i == 0 {
			noramlTime = tmp2.Data.Data[0].Timestamp
		} else {
			exitTime = tmp2.Data.Data[0].Timestamp
		}
	}
	if noramlTime < exitTime {
		return true
	}
	return false
}

// 利用拨号判断切换线路的条件之一
func AikuaiDialJudgeLine(url, setCookie string) bool {
	datas := `{
		"func_name": "monitor_iface",
		"action": "show",
		"param": {
		  "TYPE": "iface_check,iface_stream,ether_info,snapshoot"
		}
	  }`
	err, s2 := crawler.PostNew(url+"/Action/call", datas, setCookie, false)
	if err != nil {
		return false
	}
	var tmp2 NewAikuaiResp
	err = json.Unmarshal([]byte(s2), &tmp2)
	if err != nil {
		return false
	}
	if tmp2.ErrMsg != "Success" {
		return false
	}
	for _, v := range tmp2.Data.IfaceCheck {
		if v.Interface == "l2tp_tfgame2" && v.Errmsg != "线路检测成功" {
			return true
		}
	}
	return false
}

var (
	//url       string
	setCookie string
	urlx      = "http://244.mengzuofang.com:8081"
)

func main() {
	s, err := AikuaiLoginFunc(urlx)
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " crond error:" + err.Error())
		return
	}
	setCookie = s

	c := cron.New(cron.WithSeconds())
	spec := "*/30 * * * * ?"
	_, err = c.AddFunc(spec, erveryMinuteCheckLine)
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " crond error:" + err.Error())
		return
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " crond start")
	c.Start()
	select {} //阻塞主线程不退出

}

func erveryMinuteCheckLine() {
	//每天登陆下后台保持cookie最新
	t := time.Now()
	if t.Hour() == 23 && t.Minute() == 59 {
		s, err := AikuaiLoginFunc(urlx)
		if err != nil {
			fmt.Println("登陆路由后台失败：", err)
			return
		}
		setCookie = s
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" 开始登录路由后台了--->", setCookie)
	}
	//利用日志或者拨号记录判断
	if AikuaiJudgeLine(urlx, setCookie) || AikuaiDialJudgeLine(urlx, setCookie) {
		//开始切换
		coookie, _ := AikuaiLoginFunc(urlx)
		var msg string
		s, err := AikuaiSwitchLineFunc(urlx, coookie)
		if err != nil {
			msg = "切换线路失败,原因：" + err.Error()
		} else {
			msg = s
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" 操作-->", msg)
	}

}
