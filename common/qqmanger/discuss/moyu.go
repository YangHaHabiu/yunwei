/*
@Time : 2022-3-15 16:46
@Author : acool
@File : moyu
*/
package discuss

import (
	"fmt"
	"github.com/nosixtools/solarlunar"
	"math"
	"regexp"
	"strconv"
	"time"
)

//判断上下午
func judgeMorningOrAfternoon(timeHour int) string {
	if timeHour >= 12 {
		if timeHour > 19 || timeHour < 9 {
			fmt.Println("工作日下班时间，多多休息")
		}
		return "下午好"
	} else {
		return "上午好"
	}
}

//获取相差时间
func getDayDiffer(start_time, end_time string) int64 {
	var day int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		day = diff/3600/24 + 1
		return day
	}
	return day
}

//打印信息
func FishingReminder() string {
	xingqi := []string{
		"星期日",
		"星期一",
		"星期二",
		"星期三",
		"星期四",
		"星期五",
		"星期六",
	}
	s := `【摸鱼办】提醒您：
%s，摸鱼人！
今天是%s【%s】
距离本周【周末】还有%d天
距离【清明节 %s】假期还有%d天
距离【劳动节 %s】假期还有%d天
距离【端午节 %s】假期还有%d天
距离【中秋节 %s】假期还有%d天
距离【国庆节 %s】假期还有%d天
距离【元旦   %s】假期还有%d天
距离【春节   %s】假期还有%d天
工作再累 一定不要忘记摸鱼哦！有事没事起身去茶水间去厕所去廊道走走
别老在工位上坐着钱是老板的，但命是自己的
加油打工人👷‍♀
`
	nows := time.Now()
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	hour := time.Now().Hour()
	today := time.Now().Weekday()
	day := time.Saturday - today
	year := time.Now().Year() + 1
	stringYear := strconv.Itoa(year)
	timeYear := stringYear + "-01-01"
	if today == time.Sunday || today == time.Saturday {
		s = `今天摸鱼办也休息哦
周末愉快`
	} else {
		nowxx, _ := time.Parse("2006-01-02", solarlunar.LunarToSolar(fmt.Sprintf("%d-01-01", time.Now().Year()+1), false))
		timeSpring := nowxx.Format("2006-01-02")
		qingming := fmt.Sprintf("%d-04-0%.0f", nows.Year(), math.Floor(float64(nows.Year())*0.2422+4.475)-math.Floor(float64(nows.Year())/4-15)+1)
		laodong := fmt.Sprintf("%d-05-01", nows.Year())
		duanwu := fmt.Sprintf("%s", solarlunar.LunarToSolar(fmt.Sprintf("%d-05-05", nows.Year()), false))
		zhongqiu := fmt.Sprintf("%s", solarlunar.LunarToSolar(fmt.Sprintf("%d-08-15", nows.Year()), false))
		guoqing := fmt.Sprintf("%d-10-01", nows.Year())
		s = fmt.Sprintf(s,
			judgeMorningOrAfternoon(hour),
			fmt.Sprintf("%s %s", nows.Format("2006年1月2日"), solarlunar.SolarToChineseLuanr(nows.Format("2006-01-02"))),
			xingqi[nows.Weekday()],
			day-1,
			qingming, getDayDiffer(timeNow, qingming+" 00:00:00"), //清明
			laodong, getDayDiffer(timeNow, laodong+" 00:00:00"), //劳动
			duanwu, getDayDiffer(timeNow, duanwu+" 00:00:00"), //端午
			zhongqiu, getDayDiffer(timeNow, zhongqiu+" 00:00:00"), //中秋
			guoqing, getDayDiffer(timeNow, guoqing+" 00:00:00"), //国庆
			timeYear, getDayDiffer(timeNow, timeYear+" 00:00:00"), //元旦
			timeSpring, getDayDiffer(timeNow, timeSpring+" 00:00:00"), //春节
		)
	}
	//fmt.Println(s)
	re := regexp.MustCompile(`(?m)^距离.*还有0天$[\r\n]*`)
	s = re.ReplaceAllString(s, "")

	return s
}
