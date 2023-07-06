package xtime

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	nows := time.Now()
	date := nows.AddDate(0, 0, -7)
	currentDate := nows.Format("2006-01-02")
	beforeDate := date.Format("2006-01-02")
	fmt.Println(currentDate, beforeDate)
	return
	aaa := "/data/yunwei_script/yunwei_new/maintain_game"
	fmt.Println(filepath.Dir(aaa))
	var (
		location time.Time
		err      error
	)
	Content := `bkzj
更新时间:09月30日 14:45:222
install 5关进程
`

	compile2, _ := regexp.Compile(`\[CQ.*\]`)
	allString := compile2.ReplaceAllString(Content, "")
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z|\t`)
	allString = re.ReplaceAllString(allString, "")

	splits := strings.Split(allString, "\n")
	if strings.Contains(splits[1], "时间") {
		match, _ := regexp.Compile(`间:|间：`)
		i := match.Split(splits[1], -1)
		if len(i) > 1 {
			compile, _ := regexp.Compile(`月|日|:|-`)
			i2 := compile.Split(strings.ReplaceAll(i[1], " ", ""), -1)
			if len(i2) < 4 {
				return

			}
			month, _ := strconv.Atoi(i2[0])
			day, _ := strconv.Atoi(i2[1])
			hour, _ := strconv.Atoi(i2[2])
			minute, _ := strconv.Atoi(i2[3])
			all := fmt.Sprintf("%.2d-%.2d %.2d:%.2d", month, day, hour, minute)
			location, err = time.ParseInLocation("01-02 15:04", all, time.Local)
			if err != nil {
				return
			}

		} else {
			return
		}

	} else {
		return
	}

	inLocation, _ := time.ParseInLocation("01-02 15:04", time.Now().Format("01-02 15:04"), time.Local)
	now := time.Now()
	if inLocation.Month() > location.Month() {
		now = now.AddDate(1, 0, 0)
	}
	startTime := fmt.Sprintf("%d-%s", now.Year(), location.Format("01-02 15:04"))
	inLocationx, _ := time.ParseInLocation("2006-01-02 15:04", startTime, time.Local)
	fmt.Println(inLocationx)

}

func TestXxx(t *testing.T) {

	format := "2006-01-02 15:04"
	timeNow := time.Now().Format(format)
	fmt.Println(timeNow)
	x, err := time.ParseInLocation(format, timeNow, time.Local)
	if err != nil {
		return
	}
	nowMinute := x.Unix()

	fmt.Println(nowMinute)
}
