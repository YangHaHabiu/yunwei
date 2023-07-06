package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	newErr "errors"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
	"ywadmin-v3/common/go-refresh-cdn/common/model"
	comAuth "ywadmin-v3/common/go-refresh-cdn/common/wangsuSdk/auth"
)

type WangSuApi struct {
	*model.CommonApi
}

func (w *WangSuApi) DoApi() error {
	date := w.getDate()
	auth := w.encrypt(w.AccessKeyId, date, w.AccessKeySecret)
	url := "https://open.chinanetcenter.com/ccm/purge/ItemIdReceiver"
	var (
		b      []byte
		err    error
		config comAuth.AkskConfig
	)

	config.AccessKey = w.AccessKeyId
	config.SecretKey = w.AccessKeySecret
	config.EndPoint = "{endPoint}"
	config.Method = "POST"
	config.Uri = "/ccm/purge/ItemIdReceiver"
	method := "GET"
	urlList := strings.Split(w.UrlList, ",")
	var res string
	if w.Action == "query" {
		for _, v := range urlList {
			formatTemplate := "2006-01-02"
			startTime := time.Now().AddDate(0, 0, -2).Format(formatTemplate)
			endTime := time.Now().Format(formatTemplate)
			url = "https://open.chinanetcenter.com/ccm/purge/ItemIdQuery"
			config.Uri = "/ccm/purge/ItemIdQuery"
			body := make(map[string]string, 0)
			body["startTime"] = startTime + " 00:00:00"
			body["endTime"] = endTime + " 23:59:59"
			body["url"] = v
			b, err = json.Marshal(body)
			if err != nil {
				return err
			}
			if w.KeyAuth == "apiKey" {
				err = w.httpDo(url, w.getDate(), auth, "GET", "application/json", string(b))
				if err != nil {
					return err
				}
				return nil
			} else if w.KeyAuth == "accessKey" {
				response := comAuth.Invoke(config, string(b))
				jsonFormatOut(response)
				return nil
			}
		}

	} else if w.Action == "refresh" {
		body := make(map[string][]string, 0)
		if w.PurgeType == "dirs" {
			body["dirs"] = urlList
		} else if w.PurgeType == "urls" {
			body["urls"] = urlList
		}

		b, err = json.Marshal(body)
		if err != nil {
			return err
		}
		res = string(b)

	} else if w.Action == "count" {
		url = "https://open.chinanetcenter.com/ccm/upperQuery?type=purge"
		config.Uri = "/ccm/upperQuery?type=purge"
		config.Method = "GET"
	} else if w.Action == "ips" {
		if w.Ips == "" {
			return newErr.New("缺少-i ip1,ip2参数")
		}
		url = "https://open.chinanetcenter.com/api/tools/ip-info"
		config.Uri = "/api/tools/ip-info"
		method = "POST"
		body := make(map[string][]string, 0)
		body["ip"] = strings.Split(w.Ips, ",")
		b, err = json.Marshal(body)
		if err != nil {
			return err
		}
		res = string(b)
	}

	if w.KeyAuth == "apiKey" {
		err = w.httpDo(url, w.getDate(), auth, method, "application/json", res)
		if err != nil {
			return err
		}
	} else if w.KeyAuth == "accessKey" {
		response := comAuth.Invoke(config, res)
		jsonFormatOut(response)
	}
	return err
}

func (w *WangSuApi) getDate() string {
	f, err := exec.Command("date", "-R", "-u").Output()
	if err != nil {
		return ""
	}
	str := string(f)
	str = strings.TrimSpace(strings.Replace(str, "+0000", "GMT", -1)) //caution:if not trim, a space character is hidden at last position
	return str
}
func (w *WangSuApi) encrypt(accountName string, date string, apikey string) string {
	key := []byte(apikey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(date))
	value := mac.Sum(nil)
	signedApikey := base64.StdEncoding.EncodeToString(value)
	msg := base64.StdEncoding.EncodeToString([]byte(accountName + ":" + signedApikey))
	return msg
}
func (w *WangSuApi) httpDo(url string, date string, auth string, method string, accept string, requestBody string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", accept)
	req.Header.Set("Date", date)
	req.Header.Set("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	jsonFormatOut(string(body))

	return err
}
