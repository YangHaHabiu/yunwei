package group

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

//检查 slice 中某元素是否存在
func Contains(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

//读取行数
func ReadLine(fileName string) (error, [][]string) {
	f, err := os.Open(fileName)
	if err != nil {
		return err, nil
	}
	newList := make([][]string, 0)
	lineList := make([]string, 0)
	buf := bufio.NewReader(f)
	n := 1
	for {

		line, err := buf.ReadString('\n')
		lineList = append(lineList, line)
		if err != nil {
			if err == io.EOF {
				newList = append(newList, lineList)
				return nil, newList
			}
			return err, nil
		}

		if n%70 == 0 {
			newList = append(newList, lineList)
			lineList = make([]string, 0)
			n = 0
		}

		n++
	}
	return nil, nil
}

//定义返回结果JSON
type GetResult struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  string `json:"data"`
	Otype int    `json:"otype"`
}

//生成sn
func CreateSn(ts string) string {
	b := bytes.Buffer{}
	b.WriteString("appKey=")
	b.WriteString("demo")
	b.WriteString("&appSecret=")
	b.WriteString("123456")
	b.WriteString("&ts=")
	b.WriteString(ts)
	return b.String()
}

//md5加密
func Md5s(s string) string {
	data := []byte(s)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//读取文件后几行
func ReadFileLast(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return ""
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	content := make([]string, 0)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		line = strings.TrimSpace(line)
		if err == io.EOF {
			if len(line) != 0 {
				content = append(content, (line))
			}
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return ""
		}
		content = append(content, (line))
	}
	return strings.Join(content[len(content)-20:len(content)-0], "\n")
}
