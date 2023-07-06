/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: request_test.go
* @Date: 2021-4-29 15:42
 */
package crawler

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/asmcos/requests"
	"github.com/tidwall/gjson"
)

func Test2(t *testing.T) {
	var (
		url    string
		result string
	)
	url = "https://www.zhihu.com/api/v4/columns/c_1261258401923026944/items"
	response, err := GetHttpResponse(url, false)
	if err != nil {
		return
	}
	t1, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	if gjson.Get(string(response), "data.0.created").Int() >= t1.Unix() {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(gjson.Get(string(response), "data.0.content").String()))
		if err != nil {
			return
		}
		dom.Find("p").Each(func(i int, selection *goquery.Selection) {
			if len(selection.Text()) != 0 {
				result += selection.Text() + "\n"
			}
		})
	}
	fmt.Println(result)
}

func TestDemo(t *testing.T) {

	var (
		url    string
		result string
		flag   bool
	)

	todayNewsTitle := fmt.Sprintf("%s新闻早讯，每天60秒读懂世界", time.Now().Format("1月2日"))

	for i := 1; i <= 10; i++ {
		if i == 1 {
			url = "https://iehou.com/index.htm"
		} else {
			url = fmt.Sprintf("https://iehou.com/index-%d.htm", i)
		}

		response, err := GetHttpResponse(url, false)
		if err != nil {
			return
		}

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(response)))
		if err != nil {
			return
		}
		dom.Find(".list-group li").Each(func(i int, selection *goquery.Selection) {
			name := selection.Find("a").Text()
			address, _ := selection.Find(".mr-1").Attr("href")
			//times := selection.Find(".num-font").Text()

			if name == todayNewsTitle {
				fmt.Println(address)
				responses, err := GetHttpResponse(address, false)
				if err != nil {
					return
				}
				doms, err := goquery.NewDocumentFromReader(strings.NewReader(string(responses)))
				if err != nil {
					return
				}
				doms.Find(".thread-content").Each(func(i int, selections *goquery.Selection) {
					titles := selections.Find("strong").Text()
					reg := regexp.MustCompile("\\s{2,}")
					titles = reg.ReplaceAllString(titles, "\n")

					content := selections.Find("p").Text()
					//reg = regexp.MustCompile("\\s{2,}")
					content = reg.ReplaceAllString(content, "\n")

					result = titles + "\n" + content

				})
				flag = true
			}
		})
		if flag {
			break
		}
	}
	fmt.Println(result)
}

func Test3(t *testing.T) {

	var (
		//url    string
		fkdtRes string
		yqtbRes string
	)
	todaySpan := time.Now().Format("2006-01-02")
	urlList := map[string]string{
		"fkdt": "http://wsjkw.gd.gov.cn/xxgzbdfk/fkdt/",
		"yqtb": "http://wsjkw.gd.gov.cn/xxgzbdfk/yqtb/",
	}
	for keys, url := range urlList {

		response, err := GetHttpResponse(url, false)
		if err != nil {
			return
		}

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(response)))
		if err != nil {
			return
		}
		dom.Find(".section li").Each(func(i int, selection *goquery.Selection) {
			name := selection.Find("a").Text()
			span := selection.Find("span").Text()
			address, _ := selection.Find("a").Attr("href")
			//fmt.Println(name, strings.Contains(name, "最新疫情风险等级提醒"))
			if span == todaySpan && keys == "fkdt" && strings.Contains(name, "最新疫情风险等级提醒") {
				responses, err := GetHttpResponse(address, false)
				if err != nil {
					return
				}
				doms, err := goquery.NewDocumentFromReader(strings.NewReader(string(responses)))
				if err != nil {
					return
				}
				doms.Find(".content-content").Each(func(i int, selections *goquery.Selection) {
					fkdtRes, _ = selections.Find("img").Attr("src")
				})

			} else if span == todaySpan && keys == "yqtb" && strings.Contains(name, "新冠肺炎疫情情况") {
				responses, err := GetHttpResponse(address, false)
				if err != nil {
					return
				}
				doms, err := goquery.NewDocumentFromReader(strings.NewReader(string(responses)))
				if err != nil {
					return
				}
				doms.Find(".margin_lr20").Each(func(i int, selections *goquery.Selection) {
					yqtbRes = selections.Find("h1").Text()
				})
				doms.Find(".content-content").Each(func(i int, selections *goquery.Selection) {
					yqtbRes += selections.Find("p").Text()
				})

			}
		})
	}
	if fkdtRes == "" {
		fkdtRes = "防控动态未更新，请稍后再试，详情-->http://wsjkw.gd.gov.cn/xxgzbdfk/fkdt/"
	}
	if yqtbRes == "" {
		yqtbRes = "疫情通报未更新，请稍后再试，详情-->http://wsjkw.gd.gov.cn/xxgzbdfk/yqtb/"
	}
	fmt.Println(strings.ReplaceAll(yqtbRes, "　　", "\n"))
	fmt.Println(fkdtRes)
}

// 新增页面请求参数
type addReq struct {
	HotTitle    string `form:"title"`
	GameName    string `form:"gameName"`
	OperType    string `form:"operType" `
	Operator    string `form:"operator" `
	OperStatus  int    `form:"operStatus"`
	OperContent string `form:"operContent" `
}

func TestDemo1(t *testing.T) {

	ts := fmt.Sprintf("%d", time.Now().Unix())
	tt := url.Values{}
	tt.Add("title", "xx3d")
	tt.Add("gameName", "xx3d")
	tt.Add("operType", "111")
	tt.Add("operator", "jiayuanhao")
	tt.Add("operStatus", "1")
	tt.Add("types", "1")
	tt.Add("operContent", "xxxxxxxxxxxxxxxxxxxxxx")

	response, err := PostHttpResponse(fmt.Sprintf("http://10.10.88.215:8086/api/hot_update_log/add?sn=%s&ts=%s", Md5s(createSn(ts)), ts), tt.Encode(), false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(response))

}

func createSn(ts string) string {
	b := bytes.Buffer{}
	b.WriteString("appKey=")
	b.WriteString("demo")
	b.WriteString("&appSecret=")
	b.WriteString("123456")
	b.WriteString("&ts=")
	b.WriteString(ts)
	return b.String()
}

func Md5s(s string) string {
	data := []byte(s)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func TestName(t *testing.T) {
	msg := `[CQ:reply,id=E1GTuhtdZn8AALdXlEfdEmJ03bUB][CQ:at,qq=459105919,text=@运维-贾源皓] [CQ:at,qq=459105919]
操作状态：执行完成
游戏名：jswar
操作者：jiayuanhao
完成时间：2022-05-06 17:33:22
操作内容：
------------
2022-05-06 17:33:16
任务1:[成功]
    操作类型: 热更配置
    发布服: 2
    操作范围: jszza
    文件列表:
        config/Act91012Cfg.lua
        config/RenWuListCfg.lua
任务2:[成功]
    操作类型: 执行服务端命令
    操作范围: jszza
    命令列表:
        admin -s GameServer "updateCfg Act91012Cfg"
        admin -s GameServer "updateCfg RenWuListCfg"

`
	//msg = strings.ReplaceAll(strings.ReplaceAll(msg, "\\", "\\\\"), "\n", "\\n")
	//compile, _ := regexp.Compile(`"`)
	//msg = compile.ReplaceAllString(msg, "\\\"")
	//msg = strings.ReplaceAll(msg, "\"", "\\\"")
	//msg = strings.ReplaceAll(msg, "'", "\\'")
	//jsonStr := fmt.Sprintf(`{"%s_id":"%d","message":"%s","auto_escape":"false"}`, msgType, Id, msg)

	BotSend(515603955, "http://10.10.88.217:5555", msg, "group")

	//msg = strings.ReplaceAll(strings.ReplaceAll(msg, "\\", "\\\\"), "\n", "\\n")
	//msg = strings.ReplaceAll(msg, "\"", "\\\"")
	//msg = strings.ReplaceAll(msg, "'", "\\'")
	//jsonStr := fmt.Sprintf(`{"%s_id":"%d","message":"%s","auto_escape":"false"}`, "group", 324113338, msg)
	//
	//Post("http://10.10.88.217:5555/send_group_msg", jsonStr)
}

func BotSend(Id int64, url, msg, msgType string) {
	msg = HandelMsg(msg)
	jsonStr := fmt.Sprintf(`{"%s_id":"%d","message":"%s","auto_escape":"false"}`, msgType, Id, msg)
	fmt.Println(jsonStr)
	requests.PostJson(fmt.Sprintf("%s/send_%s_msg", url, msgType), jsonStr)
	//Post(fmt.Sprintf("%s/send_%s_msg", url, msgType), jsonStr)

}

func HandelMsg(data string) string {
	data = strings.ReplaceAll(data, "\\", "\\\\")
	data = strings.ReplaceAll(data, "\n", "\\n")
	data = strings.ReplaceAll(data, "\"", "\\\"")
	data = strings.ReplaceAll(data, "'", "\\'")
	return data
}

type qqMsgStatus struct {
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

func TestWeibo2(t *testing.T) {
	content, errx := GetHttpResponse(fmt.Sprintf("%s/getStatus", "http://10.10.88.217:5555"), false)

	var qqTmpStatus qqMsgStatus
	if errx == nil {
		err := json.Unmarshal(content, &qqTmpStatus)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if errx != nil || qqTmpStatus.Retcode != 0 {
		fmt.Println(33333333)
	} else {
		fmt.Println(111111111)
	}

}

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

func TestAikkuai(t *testing.T) {

	data := `{
		"username": "admin",
		"passwd": "c63e88cbfc879f3610a220a36f34dd3f",
		"pass": "c2FsdF8xMXlteGgoMTIzKQ==",
		"remember_password": "true"
	  }`
	err, s := PostNew("http://244.mengzuofang.com:8081/Action/login", data, "", true)

	if err != nil {
		return
	}
	var tmp AikuaiResp
	header := strings.Split(s, "@@")[1]
	dataHeader := strings.Split(s, "@@")[0]
	json.Unmarshal([]byte(dataHeader), &tmp)
	if tmp.ErrMsg != "Succeess" {
		return
	}
	setCookie := "username=admin; login=1; " + header
	data = `{
		"func_name": "l2tp_client",
		"action": "show",
		"param": {
		  "TYPE": "total,data,interface"
		}
	  }`
	err2, s2 := PostNew("http://244.mengzuofang.com:8081/Action/call", data, setCookie, false)
	if err2 != nil {
		return
	}
	var tmp2 AikuaiResp
	json.Unmarshal([]byte(s2), &tmp2)
	if tmp2.ErrMsg != "Success" {
		return
	}
	fmt.Println(s2)
	lineList := make([]string, 0)
	for _, v := range tmp2.Data.Interface {
		if v[0] != "auto" {
			lineList = append(lineList, v[0])
		}
	}
	fmt.Println(lineList)

	// lineList := []string{
	// 	"wan1", "wan2", "wan3", "wan4", "adsl1", "adsl2",
	// }
	index := 0
	for i, v := range lineList {
		if v == tmp2.Data.Data[1].Interface {
			index = i
		}
	}
	index = (index + 1) % len(lineList)
	//fmt.Println(lineList[index])
	data = fmt.Sprintf(`
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

	// err3, s3 := PostNew("http://244.mengzuofang.com:8081/Action/call", data, setCookie, false)
	// if err3 != nil {
	// 	return
	// }
	// var tmp3 AikuaiResp
	// json.Unmarshal([]byte(s3), &tmp3)
	// if tmp3.ErrMsg != "Success" {
	// 	return
	// }
	datas := `{
		"func_name": "monitor_iface",
		"action": "show",
		"param": {
		  "TYPE": "iface_check,iface_stream,ether_info,snapshoot"
		}
	  }`
	err, s3 := PostNew("http://244.mengzuofang.com:8081/Action/call", datas, setCookie, false)
	if err != nil {
		return
	}
	var tmp3 NewAikuaiResp
	json.Unmarshal([]byte(s3), &tmp3)
	if tmp2.ErrMsg != "Success" {
		return
	}

	for _, v := range tmp3.Data.IfaceCheck {
		fmt.Println(v.Interface, v.Errmsg, v.Result)
		if v.Interface == "l2tp_tfgame2" {
			if v.Result != "success" {
				fmt.Println(1111)
				return
			}
		}

	}

}

func TestPost(t *testing.T) {

	err, _ := Post("http://127.0.0.1:5556/sendDiscussMsg", map[string]string{
		"group_id":    "629049295",
		"message":     "1111",
		"auto_escape": "false",
	})
	fmt.Println(err)
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

func TestAA(t *testing.T) {
	out := `{"post_type":"message","message_id":"KwjMRRtdZn8AABCsWHHoUWQiT38B","user_id":459105919,"time":1679970175,"seq":4268,"rand":1483860049,"font":"微软雅黑","message":"你好小助手","raw_message":"你好小助手","message_type":"group","sender":{"user_id":459105919,"nickname":"acool","card":"","sex":"unknown","age":0,"area":"","level":1,"role":"member","title":""},"group_id":721996869,"group_name":"梦作坊运维群","block":false,"sub_type":"normal","anonymous":null,"atme":false,"atall":false,"self_id":614449311}`
	fmt.Println(out)
	var data QqMessage
	err := json.Unmarshal([]byte(out), &data)
	fmt.Println(err)
	fmt.Println(data)

}

type QqMessage struct {
	SelfID      int64   `json:"self_id"`
	Time        int     `json:"time"`
	PostType    string  `json:"post_type"`
	NoticeType  string  `json:"notice_type"`
	MessageType string  `json:"message_type"`
	SubType     string  `json:"sub_type"`
	MessageID   string  `json:"message_id"`
	DiscussId   int64   `json:"discuss_id"`
	GroupID     int64   `json:"group_id"`
	GroupName   string  `json:"group_name"`
	UserID      int64   `json:"user_id"`
	Anonymous   string  `json:"anonymous"`
	Message     string  `json:"message"`
	RawMessage  string  `json:"raw_message"`
	Atme        bool    `json:"atme"`
	Block       bool    `json:"block"`
	Seqid       int     `json:"seqid"`
	Font        string  `json:"font"`
	Sender      Senders `json:"sender"`
}
type Senders struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Level    int    `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}
