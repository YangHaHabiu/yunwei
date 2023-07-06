package notify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FeiShuRobotAccount struct {
	AccessToken string `json:"access_token"` //密钥
	Secret      string `json:"secret"`       //密钥
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
}

type WebhookFeiShu struct {
	AccessToken string
	Secret      string
	AppId       string
	AppSecret   string
	Timestamp   int64
	EnableAt    bool
	AtAll       bool
	CardSending bool
}

type Code struct {
	Extra         string `json:"Extra"`
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
	Msg           string `json:"msg"`
}

type TokenResult struct {
	Code              int64  `json:"code"`
	Expire            int64  `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

var LevelColorMap = map[string]string{
	"average":  "grey",
	"warning":  "orange",
	"error":    "violet",
	"serious":  "carmine",
	"disaster": "red",
	"normal":   "wathet",
	"other":    "blue",
	"recover":  "green",
	"remind":   "turquoise",
}

func HourTimestamp() int64 {
	layout := "2006-01-02 15:04:05"
	timeStr := time.Now().Format(layout)
	parse, err := time.ParseInLocation(layout, timeStr, time.Local)
	if err != nil {
		return 0
	}
	return parse.Unix()
}

func SendToFeishuChan(jsonContent string, enableAt, atAll, cardSending bool) (*WebhookFeiShu, error) {
	var content FeiShuRobotAccount
	err := jsonx.Unmarshal([]byte(jsonContent), &content)
	if err != nil {
		logx.Error("解析飞书消息内容失败，原因：" + err.Error())
		return nil, err
	}

	return &WebhookFeiShu{
		AccessToken: content.AccessToken,
		Secret:      content.Secret,
		AppId:       content.AppId,
		AppSecret:   content.AppSecret,
		Timestamp:   HourTimestamp(),
		EnableAt:    enableAt,
		AtAll:       atAll,
		CardSending: cardSending,
	}, nil
}
func (w *WebhookFeiShu) SendMsg(s, level, title string, at ...string) error {
	atList := make([]string, 0)
	if w.AtAll && !w.EnableAt {
		atList = append(atList, `<at user_id = "all"></at>`)
	}
	if w.EnableAt && !w.AtAll {
		if len(at) > 0 {
			for _, v := range at {
				atList = append(atList, fmt.Sprintf(`<at user_id = "%s"></at>`, v))
			}
		} else {
			return errors.New("艾特指定人需要填入指定人帐号，禁止为空")
		}

	}
	s = s + strings.Join(atList, " ")

	var (
		colors          string
		interactiveJson string
		urls            string
		flag            bool
		token           string
		err             error
	)
	if value, ok := LevelColorMap[level]; ok {
		colors = value
	} else {
		colors = LevelColorMap["other"]
	}
	if w.CardSending {
		switch level {
		case "disaster", "serious", "error", "remind", "recover":
			msg := map[string]interface{}{
				"chat_id":  "oc_e2f37ad2ce6529f98f9edaddb8d44612",
				"msg_type": "interactive",
				"card": map[string]interface{}{
					"config": map[string]interface{}{
						"wide_screen_mode": true,
					},
					"elements": []interface{}{
						map[string]interface{}{
							"tag": "div",
							"text": map[string]string{
								"content": s,
								"tag":     "lark_md",
							},
						},
					},
					"header": map[string]interface{}{
						"template": colors,
						"title": map[string]string{
							"content": title,
							"tag":     "plain_text",
						},
					},
				},
			}
			data, _ := json.Marshal(msg)
			interactiveJson = string(data)
			urls = "https://open.feishu.cn/open-apis/message/v4/send/"
			flag = true
		default:
			sign, err := w.genSign()
			if err != nil {
				return err
			}
			msg := map[string]interface{}{
				"timestamp": w.Timestamp,
				"sign":      sign,
				"msg_type":  "text",
				"content": map[string]string{
					"text": s,
				},
			}
			data, _ := json.Marshal(msg)
			interactiveJson = string(data)
			urls = "https://open.feishu.cn/open-apis/bot/v2/hook/" + w.AccessToken
		}
	} else {
		sign, err := w.genSign()
		if err != nil {
			return err
		}
		var contentx string
		switch level {
		case "disaster", "serious", "error", "remind", "recover":
			contentx = title + "\n" + s
		default:
			contentx = s
		}
		msg := map[string]interface{}{
			"timestamp": w.Timestamp,
			"sign":      sign,
			"msg_type":  "text",
			"content": map[string]string{
				"text": contentx,
			},
		}
		data, _ := json.Marshal(msg)
		interactiveJson = string(data)
		urls = "https://open.feishu.cn/open-apis/bot/v2/hook/" + w.AccessToken
	}
	if flag {
		token, err = w.getToken()
		if err != nil {
			return err
		}
	}

	err, cs := w.postRequst(urls, interactiveJson, token, flag)
	if err != nil {
		return err
	}
	var code Code
	err = json.Unmarshal(cs, &code)
	if err != nil {
		return err
	}
	if code.StatusMessage != "success" {
		if code.Msg != "ok" {
			return errors.New(string(cs))
		}
	}
	fmt.Println("发送成功")
	return nil
}

func (w *WebhookFeiShu) genSign() (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	//timestamp 为距当前时间不超过 1 小时(3600)的时间戳
	stringToSign := fmt.Sprintf("%v", w.Timestamp) + "\n" + w.Secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func (w *WebhookFeiShu) getToken() (string, error) {
	tokenUrl := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	body := `{
    "app_id": "%s",
    "app_secret": "%s"
}
`
	body = fmt.Sprintf(body, w.AppId, w.AppSecret)
	err, s := w.postRequst(tokenUrl, body, "", false)
	if err != nil {
		return "", err
	}
	var result TokenResult
	err = json.Unmarshal(s, &result)
	if err != nil {
		return "", err
	}
	if result.Msg != "ok" {
		return "", errors.New("查询卡片发送的token失败")
	}
	return result.TenantAccessToken, nil
}

// post请求
func (w *WebhookFeiShu) postRequst(urls, interactiveJson, token string, flag bool) (error, []byte) {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}, Timeout: 5 * time.Second}
	reqest, err := http.NewRequest("POST", urls, bytes.NewBuffer([]byte(interactiveJson)))
	if err != nil {
		return err, nil
	}
	if flag {
		//t-g104c2fa7NWVAVHMWPE3YXGCPK6KJQD7L7BJU6FI
		reqest.Header.Add("Authorization", "Bearer "+token)
	}
	reqest.Header.Add("Content-Type", "application/json")
	//处理返回结果
	resp, _ := client.Do(reqest)
	if resp == nil {
		return errors.New("请求失败"), nil
	}
	defer resp.Body.Close()
	cs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	defer client.CloseIdleConnections()
	return nil, cs

}
