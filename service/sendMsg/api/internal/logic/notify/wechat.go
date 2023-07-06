package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	//发送消息使用导的url
	sendurl = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
	//获取token使用导的url
	getToken = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
)

var (
	requestError = errors.New("request error,check url or network")
	WechatChan   chan *sendMsg
)

type WechatAccount struct {
	AppID          string `json:"app_id"`     // appid
	AppSecret      string `json:"app_secret"` // appsecret
	AgentId        int    `json:"agent_id"`
	Token          string `json:"token"`            // token
	EncodingAesKey string `json:"encoding_aes_key"` // EncodingAesKey
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//定义一个简单的文本消息格式
type sendMsg struct {
	Touser      string            `json:"touser"`
	Toparty     string            `json:"toparty"`
	Totag       string            `json:"totag"`
	Msgtype     string            `json:"msgtype"`
	jsonContent string            `json:"jsonContent"`
	Agentid     int               `json:"agentid"`
	Text        map[string]string `json:"text"`
	Safe        int               `json:"safe"`
}

type SendMsgError struct {
	Errcode int    `json:"errcode`
	Errmsg  string `json:"errmsg"`
}

func SendToWechatChan(touser, jsonContent string, text map[string]string) error {
	touser = strings.ReplaceAll(touser, ",", "|")
	sms := &sendMsg{
		Touser:      touser,
		Msgtype:     "text",
		jsonContent: jsonContent,
		Text:        text,
		Safe:        0,
		Toparty:     "",
		Totag:       "",
	}
	if err := sms.SendSms(); err != nil {
		logx.Errorf("发送微信信息失败，原因是：%s\n", err.Error())
		return err
	}
	return nil

}

func (s *sendMsg) SendSms() error {

	var content WechatAccount
	err := jsonx.Unmarshal([]byte(s.jsonContent), &content)
	if err != nil {
		return err
	}

	token, err := GetToken(content.AppID, content.AppSecret)
	if err != nil {
		return err
	}
	m := sendMsg{
		Touser: s.Touser,
		//Toparty: s.Toparty,
		Msgtype: s.Msgtype,
		Agentid: content.AgentId,
		//Totag:   s.Totag,
		Text: s.Text,
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = SendMsg(token.AccessToken, buf)
	if err != nil {
		println(err.Error())
		return err
	} else {
		return nil
	}

}

//发送消息.msgbody 必须是 API支持的类型
func SendMsg(AccessToken string, msgBody []byte) error {
	body := bytes.NewBuffer(msgBody)
	resp, err := http.Post(sendurl+AccessToken, "application/json", body)
	if resp.StatusCode != 200 {
		return requestError
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var e SendMsgError
	err = json.Unmarshal(buf, &e)
	if err != nil {
		return err
	}
	if e.Errcode != 0 && e.Errmsg != "ok" {
		return errors.New(string(buf))
	}
	return nil
}

//通过corpid 和 corpsecret 获取token
func GetToken(corpid, corpsecret string) (at AccessToken, err error) {
	resp, err := http.Get(getToken + corpid + "&corpsecret=" + corpsecret)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = requestError
		return
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &at)
	if at.AccessToken == "" {
		err = errors.New("corpid or corpsecret error.")
	}
	return
}
