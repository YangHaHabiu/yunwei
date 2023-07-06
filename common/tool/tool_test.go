/*
@Time : 2022-4-24 16:41
@Author : acool
@File : tool_test
*/
package tool

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gogf/gf/util/gconv"
)

func TestDemo(t *testing.T) {

	a := `
############操作:更新前端(关键字:仅推送)
wdzt2
更新时间:2022-09-29 15:06:09
客户端:
1.已推送到测试目录
2.仅推送cplat1
############操作:热更配置|执行服务端命令(关键字:已推送到发布1服|仅热更|执行|执行导出结果)
wdzt2
更新时间:2022-09-29 15:06:09
服务端:
1.已推送到发布1服
2.仅热更:splat1 1-5
############操作:日常更新(关键字:已推送到发布1服|仅停服更新)
wdzt2
更新时间:2022-09-29 15:06:09
服务端:
1.已推送到发布1服
2.仅停服更新splat1 1-5
############操作:闪断更新[含开关服](关键字:已推送到发布1服|只更程序)
wdzt2
更新时间:2022-09-29 15:06:09
服务端:
1.已推送到发布1服
2.只更程序splat1 1-5
############操作:同步数据库结构(关键字:同步所有库结构|同步游戏库结构|同步日志库结构)
wdzt2
1.仅更新:splat1 1-5
2.同步所有库结构
############操作:更新PHP或执行PHP命令(关键字:更新代码|更新日志库|更新代码并更新日志库|执行命令|php cli)
wdzt2
1.splat1 1-5更新代码并更新日志库
2.splat1 1-5执行命令:
php cli wdzt2 statTask,php cli wdzt2 statMarket
############操作:执行SQL文件(关键字:游戏库执行语句|游戏库执行语句导出结果|游戏库执行语句合并结果)
wdzt2
1.仅更新:splat1 1-5
2.日志库执行语句合并结果:
select * from stat_online_minute limit 1;
############操作(关键字:关停平台|清理回收|关停平台并清理回收|关进程|开进程|重启进程|加内存,重启|清档|修改时间|清档并修改时间|检查游戏进程|检查游戏进程不返回列表|备份日志库|备份游戏库|备份所有库)
wdzt2
splat1 1-5关停平台
############操作:获取服务器信息(关键字:获取服务器信息)
wdzt2
all获取服务器信息
############操作:宕机重启(关键字:启动数据库|启动数据库并开进程)---ip地址后必须有空格:
wdzt2
ip:127.0.0.1,127.0.0.2 启动数据库并开进程
############操作:添加新平台
wdzt2 添加新平台
平台ID：101
平台名称：newplat1
中文名称：wdzt2新平台1
接入后台：http://wdzt2-center.gztfgame.com
单服域名：douquwan.cn[此项可选]`
	//fmt.Println(a)
	ReadContentReturnMap(a)
}

type data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ReadContentReturnMap(contx string) []string {

	aaa := make([]*data, 0)
	for _, line := range strings.Split(contx, "############") {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			tmp := new(data)
			compile := regexp.MustCompile(`操作:(.*)\(.*`)
			findString := compile.FindAllStringSubmatch(line, -1)
			if len(findString) > 0 {
				//ct := strings.Split(findString[0][1], "|")
				//if len(ct) > 1 {
				//	for _, v := range ct {
				//		tmp.Key = v
				//		tmp.Value = line
				//	}
				//	fmt.Println(tmp)
				//} else {
				keys := strings.ReplaceAll(findString[0][1], "|", "或")
				tmp.Key = keys
				tmp.Value = line
				//}

			} else {
				compile = regexp.MustCompile(`操作:(.*)`)
				findString = compile.FindAllStringSubmatch(line, -1)
				if len(findString) > 0 {
					tmp.Key = findString[0][1]
					tmp.Value = line
				} else {
					tmp.Key = "其他"
					tmp.Value = line
				}
			}

			aaa = append(aaa, tmp)
			//break
		}

	}
	for _, v := range aaa {
		fmt.Println(v.Key, "--------------->", v.Value)
	}

	return nil
}

func TestName1(x *testing.T) {
	t := time.Now()
	beforeMonth := t.AddDate(0, 0, -7)
	format := beforeMonth.Format("2006-01-02")
	parse, _ := time.ParseInLocation("2006-01-02", format, time.Local)
	//fmt.Println(parse.Unix(), format)
	fmt.Println(parse)

}

/**
* size 随机码的位数
* kind 0    // 纯数字
       1    // 小写字母
       2    // 大写字母
       3    // 数字、大小写字母
*/
func randomString(size int, kind int) string {
	ikind, kinds, rsbytes := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		rsbytes[i] = uint8(base + rand.Intn(scope))
	}
	return string(rsbytes)
}

func TestDc(t *testing.T) {
	//fmt.Println(randomString(1, 0))
	//设置随机数种子，由于种子数值，每次启动都不一样
	//所以每次随机数都是随机的
	rand.Seed(time.Now().UnixNano())
	var (
		average   = 17
		peopleNum = 6
	)
	flag.IntVar(&average, "v", 17, "均价")
	flag.IntVar(&peopleNum, "p", 6, "人数")
	flag.Parse()
	totalMoney := average * peopleNum
	//fmt.Println(totalMoney)
	content, err := ioutil.ReadFile("./result.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	list := strings.Split(string(content), "\n")
	sum := 0
	cai := make([]string, 0)
	for i := 0; i < len(list); i++ {
		info := list[rand.Intn(len(list)-1)]
		split := strings.Split(info, "-")
		price := gconv.Int(strings.ReplaceAll(split[1], "元", ""))
		sum += price
		cai = append(cai, info)
		if len(cai) > peopleNum || sum > totalMoney {
			break
		}
	}

	strings.Join(cai, "\n")
	str := fmt.Sprintf("总花费：%d元，人均：%.2f元\n%s", sum, float32(sum)/float32(peopleNum), strings.Join(cai, "\n"))
	err = ioutil.WriteFile("./盲盒.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
