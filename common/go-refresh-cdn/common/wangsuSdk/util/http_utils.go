package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"ywadmin-v3/common/go-refresh-cdn/common/wangsuSdk/model"
)

func Call(requestMsg model.HttpRequestMsg) string {
	client := &http.Client{}
	//fmt.Println(requestMsg)
	req, err := http.NewRequest(requestMsg.Method, requestMsg.Url, strings.NewReader(requestMsg.Body))
	if err != nil {
		// handle error
		fmt.Println("http request", err)
		return ""
	}
	for key := range requestMsg.Headers {
		req.Header.Set(key, requestMsg.Headers[key])
	}
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("resp body", err)
		return ""
	}
	//fmt.Println(resp)
	return string(body)
}
