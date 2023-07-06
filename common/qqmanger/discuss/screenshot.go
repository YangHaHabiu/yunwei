package discuss

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"time"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xcmd"
)

func ScreenShot(urlx string) (string, error) {
	// 参数
	token := "628ae8ac54d76"
	url := url2.QueryEscape(urlx)
	width := 1280
	height := 800
	full_page := 1
	// 构造URL
	query := "https://www.screenshotmaster.com/api/v1/screenshot"
	query += fmt.Sprintf("?token=%s&url=%s&width=%d&height=%d&full_page=%d",
		token, url, width, height, full_page)

	// 调用API
	resp, err := http.Get(query)
	if err != nil {
		//panic(err)
		return "", err
	}
	defer resp.Body.Close()

	// 检查是否调用成功
	if resp.StatusCode != 200 {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error while calling api %s", errorBody)
	}

	// 保存截图
	getwd, _ := os.Getwd()
	fileNames := fmt.Sprintf("%s.png", time.Now().Format("20060102"))
	filePath := fmt.Sprintf("%s/images", getwd)
	file, err := os.Create(fmt.Sprintf("%s/%s", filePath, fileNames))
	compile, _ := regexp.Compile(`//|:`)
	split := compile.Split(globalkey.QqMsgKey["discuss_qqapi"].(string), -1)
	cmdsExec := fmt.Sprintf(`rsync -avzc --delete -e "ssh -i /root/.ssh/id_rsa -o StrictHostKeyChecking=no -o GSSAPIAuthentication=no -p 22" %s/images root@%s:/tmp/`, getwd, split[2])

	job := xcmd.NewCommandJob(1*time.Hour, cmdsExec)
	if !job.IsOk {
		return "", errors.New(job.ErrMsg)
	}
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}
	return fileNames, nil
}
