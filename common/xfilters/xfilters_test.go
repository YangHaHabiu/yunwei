/*
@Time : 2022-4-1 10:06
@Author : acool
@File : xfilters_test.go
*/
package xfilters

import (
	"fmt"
	"regexp"
	"testing"
)

func TestXfilters(t *testing.T) {

	fmt.Println(Xfilters([]interface{}{
		"name__=", "dxg",
		"create_time__range", "2022-01-01,2022-02-02",
		"aa@bb__or__regexp", "111",
		"platform_en not__like", "cross",
	}...))

	a := "热更配置(jszza)"
	compile := regexp.MustCompile(`\(.*\)`)
	allString := compile.ReplaceAllString(a, "")
	fmt.Println(allString)

	for i, v := range []int{
		1, 2, 3, 4,
	} {
		fmt.Println(i, v)
	}
	//a := "/admin/user/edit1/1" //源
	//a1 := "/admin/user/edit"   //匹配
	//
	//b := "/admin/user/list1?a=1&b=2"
	//b1 := "/admin"
	//fmt.Println(strings.Contains(a, a1))
	//fmt.Println(strings.Contains(b, b1))
	//
	//compile := regexp.MustCompile(`[/|?]`)
	//split := compile.Split(b1, -1)
	//fmt.Println(strings.Join(split[:4], "/"))
	//a := `[[1,2,3],[4,5,6]]`
	//var b [][]int64
	//err := json.Unmarshal([]byte(a), &b)
	//fmt.Println(err)
	//fmt.Println(b)
	//fmt.Println(b[0][1])
	//fmt.Println(1 << 10)
	////1的二进制位 左移10位，高位抛弃，低位补0
	//fmt.Println(10 >> 1)
	////10的二进制 1010 右移1位 0101
	//_, err := crawler.GetHttpResponse("http://10.10.188.215:5558/getVersionInfo", false)
	//fmt.Println(err)
	//fmt.Printf("%v", err)

	//fmt.Println(string(s))

	//a := ""
	//fmt.Println(len(strings.Split(a, ",")))
	//fmt.Println(len(a))
	//for {
	//	resp, err := http.Get("http://10.10.88.240:8888/health")
	//	defer resp.Body.Close()
	//	b, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Printf("error:%+v", err)
	//	}
	//	log.Println(string(b))
	//	time.Sleep(100 * time.Microsecond)
	//}
	//ip := "10.10.21.59, 10.10.88.215"
	//ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	//fmt.Println(ip)
	//a := "/ad/aa/33"
	//b := strings.Split(a, "/")
	//fmt.Println(len(b))
	//result := "&_t=1655"
	//compile := regexp.MustCompile(`&_t=\d+|_t=\d+`)
	//result = compile.ReplaceAllString(result, "")
	//fmt.Println(result)

	//	Content := `cjzz
	//更新时间:8月1日 15:00
	//
	//客户端：（0728）
	//1、已推送到 h5-测试/正式
	//2、仅推送 wdsmwx
	//
	//服务端：（wx20220728）
	//1、已推送到 cjzz_发布3服
	//2、仅停服更新：wdsmxcx
	//
	//wdsmwx客户端版本：109
	//
	//
	//`
	//	compile2, _ := regexp.Compile(`\[CQ.*\]`)
	//	allString := compile2.ReplaceAllString(Content, "")
	//	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z|\t`)
	//	allString = re.ReplaceAllString(allString, "")
	//	//gameSlice := strings.Split(allString, "\n")
	//	//gameName := strings.Split(strings.TrimSpace(gameSlice[0]), " ")[0]
	//	match, _ := regexp.Compile(`间:|间：`)
	//	splits := strings.Split(allString, "\n")
	//	i := match.Split(splits[1], -1)
	//	if len(i) > 1 {
	//		r, err2 := regexp.Compile(`\s+`)
	//		fmt.Println(err2)
	//		fmt.Println(i[1])
	//		replaceAllString := r.ReplaceAllString(i[1], " ")
	//		fmt.Println(replaceAllString)
	//
	//		compile, _ := regexp.Compile(`月|日|:`)
	//		i2 := compile.Split(strings.ReplaceAll(replaceAllString, " ", ""), -1)
	//		if len(i2) != 4 {
	//			fmt.Println(111)
	//
	//		}
	//		month, _ := strconv.Atoi(i2[0])
	//		day, _ := strconv.Atoi(i2[1])
	//		hour, _ := strconv.Atoi(i2[2])
	//		minute, _ := strconv.Atoi(i2[3])
	//		all := fmt.Sprintf("%.2d-%.2d %.2d:%.2d", month, day, hour, minute)
	//		//format := time.Now().Format("01-02 15:04")
	//		//inLocation, _ := time.ParseInLocation("01-02 15:04", format, time.Local)
	//		//fmt.Println(inLocation)
	//		location, err := time.ParseInLocation("01-02 15:04", all, time.Local)
	//		fmt.Println(location.Format("01-02 15:04"))
	//		fmt.Println(err)
	//
	//	}
}
