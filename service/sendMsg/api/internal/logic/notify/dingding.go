package notify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DingDingWorkNoticeAccount struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	AgentId   string `json:"agentId"`
}

type DingDingRobotAccount struct {
	AccessToken string `json:"access_token"` //密钥
	Secret      string `json:"secret"`       //密钥
}

type Webhook struct {
	AccessToken string
	Secret      string
	EnableAt    bool
	AtAll       bool
}

type errCode struct {
	//{"errcode":300001,"errmsg":"description: robot 不存在；solution:请确认 token 是否正确；"}
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func SendToDingDingChan(jsonContent string, enableAt, atAll bool) (*Webhook, error) {
	var content DingDingRobotAccount
	err := jsonx.Unmarshal([]byte(jsonContent), &content)
	if err != nil {
		logx.Error("解析钉钉消息内容失败，原因：" + err.Error())
		return nil, err
	}

	return &Webhook{
		AccessToken: content.AccessToken,
		Secret:      content.Secret,
		EnableAt:    enableAt,
		AtAll:       atAll,
	}, nil

}

// SendMessage Function to send message
//
//goland:noinspection GoUnhandledErrorResult
func (t *Webhook) SendMessage(s string, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": s,
		},
	}
	if t.EnableAt {
		if t.AtAll {
			if len(at) > 0 {
				return errors.New("the parameter \"AtAll\" is \"true\", but the \"at\" parameter of SendMessage is not empty")
			}
			msg["at"] = map[string]interface{}{
				"isAtAll": t.AtAll,
			}
		} else {
			msg["at"] = map[string]interface{}{
				"atMobiles": at,
				"isAtAll":   t.AtAll,
			}
		}
	} else {
		if len(at) > 0 {
			return errors.New("the parameter \"EnableAt\" is \"false\", but the \"at\" parameter of SendMessage is not empty")
		}
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(t.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var code errCode
	json.Unmarshal(text, &code)
	if code.Errcode != 0 {
		return errors.New(string(text))
	}
	return nil
}

func (t *Webhook) hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (t *Webhook) getURL() string {
	wh := "https://oapi.dingtalk.com/robot/send?access_token=" + t.AccessToken
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, t.Secret)
	sign := t.hmacSha256(stringToSign, t.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
