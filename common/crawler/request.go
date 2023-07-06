/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: require
* @Date: 2021-4-29 15:28
 */
package crawler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GetHttpResponse(url string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("http request fail")
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		return nil, errors.New("http response fail")
	}

	defer response.Body.Close()
	//fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New("http status code fail")
	}
	fmt.Println(response.StatusCode)
	if ok {

		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	} else {
		return ioutil.ReadAll(response.Body)
	}

}

func PostHttpResponse(url string, body string, ok bool) ([]byte, error) {
	payload := strings.NewReader(body)
	requests, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, errors.New("http request fail")
	}
	requests.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requests.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(requests)
	if err != nil {
		return nil, errors.New("http response fail")
	}

	defer response.Body.Close()
	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}
	return ioutil.ReadAll(response.Body)

}

func Post(url string, data interface{}) (error, string) {
	var marshal []byte
	switch data.(type) {
	case string:
		marshal = []byte(data.(string))
	case map[string]string:
		marshal, _ = json.Marshal(data)

	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshal))
	if err != nil {
		return err, ""
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(result))
	return err, string(result)

}

func PostNew(url string, data interface{}, header string, isGetRespHeader bool) (error, string) {
	var marshal []byte
	switch data.(type) {
	case string:
		marshal = []byte(data.(string))
	case map[string]string:
		marshal, _ = json.Marshal(data)

	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshal))
	if err != nil {
		return err, ""
	}
	req.Header.Add("Content-Type", "application/json")
	if header != "" {
		req.Header.Add("Cookie", header)

	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	if isGetRespHeader {
		return err, string(result) + "@@" + resp.Header.Get("Set-Cookie")
	}

	return err, string(result)

}
