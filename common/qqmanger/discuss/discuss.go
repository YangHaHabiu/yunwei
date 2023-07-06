/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: discuss
* @Date: 2021-7-15 14:35
 */
package discuss

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"ywadmin-v3/common/crawler"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/qqmanger/bot_send"
	"ywadmin-v3/common/xcmd"
)

type MessageCache struct {
	ID      string
	UserID  int64
	Message string
}

type wordList struct {
	KeyWord  string `json:"key_word"`
	KeyValue string `json:"key_value"`
	KeyType  string `json:"key_type"`
	Words    string `json:"words"`
}

type tokenRes struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type UnmarshalJson struct {
	GameInfo []struct {
		Content string `json:"content"`
	} `json:"game_info"`
}

type TotalJson struct {
	Data map[string]int
}

func getTongJi(url string) (result string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
		return ""
	}
	return string(body)

}

func GetDaypartingNowday(startTime string) string {
	var timeStr string
	if startTime == "" {
		timeStr = time.Now().Format("20060102")
	} else {
		timeStr = startTime
	}
	//timeStr := time.Now().Format("20060102")
	t1, err := time.ParseInLocation("20060102", timeStr, time.Local)
	tm1, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return ""
	}
	if tm1.Hour() >= 0 && tm1.Hour() < 12 {
		return t1.Format("2006-01-02") + "上午"
	} else {
		return t1.Format("2006-01-02") + "下午"
	}

}

func shiju() string {
	resp, err := http.Get("https://v2.jinrishici.com/token")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
		return ""
	}

	//fmt.Println(string(body))
	var tokenres tokenRes
	err = json.Unmarshal(body, &tokenres)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//fmt.Println(tokenres.Data)
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "https://v2.jinrishici.com/one.json", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("X-User-Token", tokenres.Data)
	//处理返回结果
	resps, err := client.Do(req)
	if err != nil {
		return ""
	}
	//关闭流
	defer resps.Body.Close()
	//检出结果集
	bodys, err := ioutil.ReadAll(resps.Body)
	if err != nil {
		//fmt.Println("ioutil.ReadAll failed ,err:%v", err)
		return ""
	}
	//fmt.Println(string(bodys))
	value := gjson.Get(string(bodys), "data.origin.content")
	//fmt.Println(value.Array())
	var results string
	for _, v := range value.Array() {
		results += v.String() + "\n"
	}
	return results

}

func getWebJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	fmt.Println("----->", url)
	if err != nil {
		// handle error
		fmt.Println("http error", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("read http body error", err.Error())
		return nil, err
	}
	return body, nil

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

func ReportDiscuss(update QqMessage, groupType string) {
	if update.Message == "hf" || update.Message == "tjhf" {
		r, _ := regexp.Compile(`\s+`)
		split := r.Split(update.Message, -1)
		var startTime string
		if len(split) > 2 {
			startTime = ""
			bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), "时间输入错误，格式：hf 20210910", groupType)
			return
		} else if len(split) == 2 {
			_, err := time.Parse("20060102", split[1])
			if err != nil {
				bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), "时间输入错误，格式：hf 20210910", groupType)

				return
			}
			startTime = split[1]
		}
		accomplish, _ := getWebJson("http://tf-alarm.gztfgame.com:85/chart/get_today_combine" + "?start_time=" + startTime)
		totals, _ := getWebJson("http://tf-alarm.gztfgame.com:85/chart/get_today_combine_total" + "?start_time=" + startTime)
		var (
			combinejson UnmarshalJson
			//combineComplish map[string]int
			combineTotal TotalJson
			results      []string
		)
		//r := `{"game_failure_count":[0,0,0],"game_list":["jnqk","wdzt","xx3d"],"game_success_count":[4,2,26]}`
		err := json.Unmarshal(accomplish, &combinejson)
		if err != nil {
			fmt.Println("json unmarshal error", err.Error())
		}
		results = append(results, GetDaypartingNowday(startTime)+"合服情况如下:")

		err = json.Unmarshal(totals, &combineTotal)
		if err != nil {
			fmt.Println("json unmarshal error", err.Error())
		}

		results = append(results, "本次需要合服总数量:")
		for k, v := range combineTotal.Data {
			res := fmt.Sprintf("%s:%d ", k, v)
			results = append(results, res)

		}

		for _, v := range combinejson.GameInfo {
			compile, _ := regexp.Compile(`,`)
			results = append(results, compile.ReplaceAllString(v.Content, "\n"))
		}
		if len(results) < 3 {
			results = make([]string, 0)
			results = append(results, GetDaypartingNowday(startTime)+"无合服计划")
		}
		result := strings.Join(results, "\n")
		bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), result, groupType)

	}

	//if update.Message == "rs" || update.Message == "热搜" {
	//	bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), crawler.Weibo(), groupType)
	//}

	if update.Message == "今日新闻" {
		bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), TodayNews(), groupType)
	}
	if update.Message == "今日疫情" {
		bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), TodayYiqing(), groupType)
	}

	if update.Message == "tfmm" || update.Message == "泰逢密码" {
		message := `===========无线===========
TaiFengGame
密码：@WIFITFGAME01

TaiFengGame1
密码：@WIFITFGAME01

TaiFengGame-2.4G
密码：Um3n3QVcea

(推荐)
COCOWAN
密码：WiFi@TFgame666


COCOWAN-2.4G
密码：WiFi@TFgame666

===========网盘===========
\\share.intranet.com\cocowan_public
cocowan
3jRTrN
`
		bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), message, groupType)

	}

	if update.Message == "摸鱼办" {
		bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), FishingReminder(), groupType)
	}
	if strings.Contains(update.Message, "悠饭") {

		r, _ := regexp.Compile(`\s+`)
		split := r.Split(update.Message, -1)
		if len(split) != 2 {
			//bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), "悠饭余额都不输，你查个寂寞", groupType)
			return
		}

		res := `您输入的悠饭金额为：%s元
本月还剩%d个工作日
您以后的日子里每餐还能吃%.2f元，加油打工人!%s
`
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()
		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		tomorryMonth := now.AddDate(0, 0, 1)
		begin, _ := time.Parse("2006-01-02", tomorryMonth.Format("2006-01-02"))
		end, _ := time.Parse("2006-01-02", lastOfMonth.Format("2006-01-02"))
		hour := CalWorkHour(begin, end)
		atoi, err := strconv.ParseFloat(split[1], 64)
		if err != nil || atoi <= 0 || atoi >= 1000 {
			bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), "你再乱写悠饭金额，我就打死你", groupType)

			return
		}
		a, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(atoi)/float64(hour*2)), 64)
		var warning string
		if a <= 10 {
			warning = "\n平均每餐不足10元,您准备喝西北风吧!!"
		} else if a >= 20 {
			warning = "\n恭喜您土豪，请我吃个烧鹅腿吧。"
		}
		var msgs string
		if hour == 0 {
			msgs = "本月已经最后一天了，无需查询了"
		} else {
			msgs = fmt.Sprintf(res, fmt.Sprintf("%.2f", atoi), hour, a, warning)
		}
		bot_send.BotSend(update.DiscussId, globalkey.QqMsgKey[update.MessageType+"_qqapi"].(string), msgs, groupType)

	}

}

func CalWorkHour(begin, end time.Time) int {
	var (
		currenTime   = begin
		workingCount int
	)
	jiejiari := map[string]bool{
		"04-02": true,
		"04-03": false,
		"04-04": false,
		"04-05": false,
		"04-24": true,
		"04-30": false,
		"05-01": false,
		"05-02": false,
		"05-03": false,
		"05-04": false,
		"05-07": true,
		"06-03": false,
		"06-04": false,
		"06-05": false,
		"09-10": false,
		"09-11": false,
		"09-12": false,
		"10-01": false,
		"10-02": false,
		"10-03": false,
		"10-04": false,
		"10-05": false,
		"10-06": false,
		"10-07": false,
		"10-08": true,
		"10-09": true,
	}
	fmt.Println(jiejiari)
	for {
		if currenTime.After(end) {
			break
		}
		if j, ok := jiejiari[currenTime.Format("01-02")]; ok {
			if j {
				workingCount++
			}

		} else {
			if currenTime.Weekday() == time.Sunday || currenTime.Weekday() == time.Saturday {
			} else {
				workingCount++
			}
		}

		currenTime = currenTime.Add(24 * time.Hour)
	}
	return workingCount
}

func TodayNews() (result string) {
	var (
		url string
	)
	url = "https://www.zhihu.com/api/v4/columns/c_1261258401923026944/items"
	response, err := crawler.GetHttpResponse(url, false)
	if err != nil {
		return
	}
	t1, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	if gjson.Get(string(response), "data.0.created").Int() >= t1.Unix() {
		fmt.Println(string(response))
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(gjson.Get(string(response), "data.0.content").String()))
		if err != nil {
			fmt.Println(err)
			return
		}
		dom.Find("p").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
			if len(selection.Text()) != 0 {
				result += selection.Text() + "\n"
			}
		})
	} else {
		result = "今日无新闻更新，详情：https://www.zhihu.com/column/c_1261258401923026944"
	}
	return
}

func TodayYiqing() (result string) {
	var (
		//url    string
		fkdtRes string
		yqtbRes string
		//err     error
	)
	todaySpan := time.Now().Format("2006-01-02")
	urlList := map[string]string{
		"fkdt": "http://wsjkw.gd.gov.cn/xxgzbdfk/fkdt/",
		"yqtb": "http://wsjkw.gd.gov.cn/xxgzbdfk/yqtb/",
	}
	for keys, url := range urlList {

		response, err := crawler.GetHttpResponse(url, false)
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
				responses, err := crawler.GetHttpResponse(address, false)
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
				responses, err := crawler.GetHttpResponse(address, false)
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
				yqtbRes = strings.ReplaceAll(yqtbRes, "　　", "\n")
			}
		})
	}

	if yqtbRes == "" {
		yqtbRes = "疫情通报未更新，请稍后再试，详情请戳-->" + urlList["yqtb"]
	}
	getwd, _ := os.Getwd()
	fileNames := fmt.Sprintf("%s.jpg", time.Now().Format("20060102"))
	filePath := fmt.Sprintf("%s/images", getwd)
	//不存在目录则创建一个
	_, e := os.Stat(filePath)
	if e != nil {
		if os.IsNotExist(e) {
			e := os.Mkdir(filePath, os.ModePerm)
			if e != nil {
				fmt.Println(e)
			}
		}
	}
	picName := fmt.Sprintf("%s/%s", filePath, fileNames)
	results := fmt.Sprintf("[CQ:image,file=/tmp/images/%s]", fileNames)
	if fkdtRes == "" {
		results = "防控动态未更新，请稍后再试，详情请戳-->" + urlList["fkdt"]
	} else {
		exists, err := PathExists(picName)
		fmt.Println(err)
		if exists {
			return "===============疫情通报===============\n" + yqtbRes + "\n===============防控动态===============\n" + results
		}
		resp, err := http.Get(fkdtRes)
		//打开文件流
		f, errf := os.Create(picName)
		if errf != nil {
			fmt.Println("os create err:", errf)
			return
		}
		defer f.Close()

		buf := make([]byte, 4096)

		//读httpbody数据写入文件流
		for {
			n, err2 := resp.Body.Read(buf)
			if n == 0 {
				break
			}
			if err2 != nil && err2 != io.EOF {
				err = err2
				return
			}

			f.Write(buf[:n])
		}
		//推送到目标
		compile, _ := regexp.Compile(`//|:`)
		split := compile.Split(globalkey.QqMsgKey["discuss_qqapi"].(string), -1)
		cmdsExec := fmt.Sprintf(`rsync -avzc --delete -e "ssh -i /root/.ssh/id_rsa -o StrictHostKeyChecking=no -o GSSAPIAuthentication=no -p 22" %s/images root@%s:/tmp/`, getwd, split[2])
		//fmt.Println(cmdsExec)
		job := xcmd.NewCommandJob(1*time.Hour, cmdsExec)
		if !job.IsOk {
			fmt.Println(job.ErrMsg)
		}
	}
	return "===============疫情通报===============\n" + yqtbRes + "\n===============防控动态===============\n" + results
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
