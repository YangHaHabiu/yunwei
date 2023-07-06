package tool

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"
)

//批量创建目录
func CreateMutiDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//当前路径下获取文件
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			s = append(s, fi.Name())
		}
	}
	return s, nil
}

//写文件
func WriteFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content) //将数据先写入缓存
	if err != nil {
		return err
	}
	writer.Flush() //将缓存中的内容写入文件
	return nil
}

//追加文件
func WriteFileAppend(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//fmt.Println("open file failed, err:", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	//time.Now().Format("2006-01-02 15:04:05") + "\n"
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}
	err = writer.Flush() //将缓存中的内容写入文件
	return err
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}
