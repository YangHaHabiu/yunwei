package bot_send

import (
	"fmt"
	"strings"
	"ywadmin-v3/common/crawler"
)

func BotSend(Id int64, url, msg, msgType string) {
	jsonStr := fmt.Sprintf(`{"%s_id":"%d","message":"%s","auto_escape":"false"}`, msgType, Id, HandelMsg(msg))
	//requests.PostJson(fmt.Sprintf("%s/send_%s_msg", url, msgType), jsonStr)
	//fmt.Println(jsonStr)
	_, result := crawler.Post(fmt.Sprintf("%s/send_%s_msg", url, msgType), jsonStr)
	//fmt.Println(result)
	if len(result) == 0 {
		jsonStr = fmt.Sprintf(`{"%s_id":"%d","message":"%s","auto_escape":"false"}`, msgType, Id, "发送消息失败，请检查消息格式")
		crawler.Post(fmt.Sprintf("%s/send_%s_msg", url, msgType), jsonStr)
	}
	//fmt.Println(err)

}

func HandelMsg(data string) string {
	data = strings.ReplaceAll(data, "\\", "\\\\")
	data = strings.ReplaceAll(data, "\t", "\\t")
	data = strings.ReplaceAll(data, "\n", "\\n")
	data = strings.ReplaceAll(data, "\"", "\\\"")
	data = strings.ReplaceAll(data, "'", "\\'")
	return data
}
