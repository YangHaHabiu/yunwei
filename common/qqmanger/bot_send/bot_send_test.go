package bot_send

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	a := `xx3d小助手更新模板`
	split := strings.Split(a, "小助手更新模板")
	fmt.Println(strings.TrimSpace(strings.Join(split, "")))
	match, _ := regexp.Compile(`间:|间：`)

	allString := `jswar
更新时间：5月12日 16:05
`
	splits := strings.Split(allString, "\n")
	i := match.Split(splits[1], -1)
	compile, _ := regexp.Compile(`月|日|:`)
	i2 := compile.Split(strings.ReplaceAll(i[1], " ", ""), -1)
	month, _ := strconv.Atoi(i2[0])
	day, _ := strconv.Atoi(i2[1])
	hour, _ := strconv.Atoi(i2[2])
	minute, _ := strconv.Atoi(i2[3])
	all := fmt.Sprintf("%.2d-%.2d %.2d:%.2d", month, day, hour, minute)
	format := time.Now().Format("01-02 15:04")
	inLocation, _ := time.ParseInLocation("01-02 15:04", format, time.Local)
	location, _ := time.ParseInLocation("01-02 15:04", all, time.Local)
	sub := inLocation.Sub(location)
	fmt.Println(sub.Seconds())

}
