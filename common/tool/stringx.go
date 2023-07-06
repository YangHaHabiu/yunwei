package tool

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"
	"unicode"

	"github.com/axgle/mahonia"
)

var (
	//匹配前三位
	TopThreeArray = []string{
		"/admin/log/syslog",
		"/admin/log/loginlog",
		"/admin/user/login",
		"/admin/user/logout",
		"/admin/user/currentUser",
		"/admin/user/reSetPassword",
		"/admin/user/updatePersonalPasswordData",
		"/admin/captcha/checkCaptcha",
		"/admin/captcha/getCaptcha",
	}

	//匹配第三位
	ThreeArray = []string{
		"list",
	}
)

//判断某个字符串在数组中
func StrInArr(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// 通过map主键唯一的特性过滤重复元素
func RemoveDuplicate(arr []string) []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		//判断主键为val的map是否存在
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = nil
		}
	}
	return resArr
}

func RandCreator(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strList := []byte(str)
	result := []byte{}
	i := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
	return string(result)
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

//GBK转utf8的方法
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 输出json格式的字符串回包
func JsonFormatOut(result string) string {
	var str bytes.Buffer
	err := json.Indent(&str, []byte(result), "", "    ")
	if err != nil {
		return result
	}
	return str.String()
}
