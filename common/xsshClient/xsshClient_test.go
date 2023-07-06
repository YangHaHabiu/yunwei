package xsshClient

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	sshClient, err := NewSSHClient("10.10.88.215", 22, AuthConfig{User: "root", KeyFile: "../../service/ws/api/key/id_rsa"})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sshClient.Close()
	//第一次 执行命令
	execinfo, err := sshClient.Exec("ls -lha /dev --time-style=\"+%Y-%m-%d %H:%I:%S\" |sort  -k2nr  |grep -v total|grep -v grep")
	fmt.Println(execinfo.OutputString(), err)
	result := readLine(execinfo.OutputString())
	for _, v := range result {
		fmt.Println(v)
	}
	a := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(a[len(a)-1])

	//第二次执行命令
	//out1, exitcode2 := sshClient.Exec("ifconfig -a")
	//fmt.Println(out1.OutputString(), exitcode2)
	// 上传文件
	//transInfoUpload, err := sshClient.Upload("D:\\Documents\\Tencent Files\\459105919\\FileRecv\\1.jpg", "/tmp/password_upload/1.jpg")
	//fmt.Println(transInfoUpload, err)
	// 下载文件
	//transInfoDownload, err := sshClient.Download("/etc/passwd", "./passwd_download/passwd")
	//fmt.Println(transInfoDownload, err)
}

func readLine(content string) (tmp []*AssetFileData) {
	split := strings.Split(content, "\n")
	tmp = make([]*AssetFileData, 0)
	for _, v := range split {
		r := regexp.MustCompile(` -> .*`)
		v = r.ReplaceAllString(v, "")

		compile := regexp.MustCompile(`[ ]+`)
		x := compile.Split(v, -1)

		var (
			kind   string
			isLink bool
			size   string
		)

		if strings.HasPrefix(x[0], "p") {
			kind = "p"
		} else if strings.HasPrefix(x[0], "c") {
			kind = "c"
		} else if strings.HasPrefix(x[0], "d") {
			kind = "d"
		} else if strings.HasPrefix(x[0], "b") {
			kind = "b"
		} else if strings.HasPrefix(x[0], "_") {
			kind = "_"
		} else if strings.HasPrefix(x[0], "l") {
			kind = "l"
			isLink = true
		} else if strings.HasPrefix(x[0], "s") {
			kind = "s"
		} else {
			kind = "?"
		}
		if len(x) >= 7 {
			mustCompile := regexp.MustCompile(`,`)
			allString := mustCompile.FindAllString(x[4], -1)
			if len(allString) == 0 {
				size = x[4]
			}
			tmp = append(tmp, &AssetFileData{
				IsLink: isLink,
				Kind:   kind,
				Name:   x[len(x)-1],
				Date:   x[len(x)-3] + " " + x[len(x)-2],
				Size:   size,
				Code:   x[0],
			})
		}

	}

	return
}

type AssetFileData struct {
	Code   string `json:"code"`
	Date   string `json:"date"`
	IsLink bool   `json:"isLink"`
	Kind   string `json:"kind"`
	Name   string `json:"name"`
	Size   string `json:"size"`
}
