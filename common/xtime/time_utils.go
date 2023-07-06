package xtime

import (
	"github.com/gogf/gf/util/gconv"
	"time"
)

var (
	//全局时间格式带秒
	FormatTimeWithSecond = "2006-01-02 15:04:05"

	//全局时间格式带分
	FormatTimeWithMinu = "2006-01-02 15:04"
)

//获取相差时间
func GetHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation(FormatTimeWithSecond, start_time, time.Local)
	t2, err := time.ParseInLocation(FormatTimeWithSecond, end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

//将字符串时间转换时间戳
func GetTimestampByStrTime(times string) (result time.Time) {
	result, err := time.ParseInLocation(FormatTimeWithMinu, times, time.Local)
	if err != nil {
		return
	}
	return
}

func GetTimeByDay(times string) int64 {
	result, err := time.ParseInLocation("2006-01-02", times, time.Local)
	if err != nil {
		return 0
	}
	return result.Unix()
}

//将时间戳转换字符串
func GetTimetampByTimeInt(timeInt int64) string {
	return time.Unix(timeInt, 0).Format(FormatTimeWithSecond)
}

//将时间戳转换字符串
func GetTimetampByTimeMinu(timeInt int64) string {
	return time.Unix(timeInt, 0).Format(FormatTimeWithMinu)
}

//将秒时间转换分钟时间戳
func GetMinTimeChangeSecondTime(times string) int64 {
	r := time.Unix(gconv.Int64(times), 0).Format(FormatTimeWithMinu)
	return (GetTimestampByStrTime(r)).Unix()
}

//将字符时间转分钟时间戳
func GetTimeStampByStrTime(strTime string) int64 {
	var ts time.Time
	ts, _ = time.ParseInLocation(FormatTimeWithSecond, strTime, time.Local)
	if ts.Unix() <= 0 {
		ts, _ = time.ParseInLocation(FormatTimeWithMinu, strTime, time.Local)
	}
	r := time.Unix(ts.Unix(), 0).Format(FormatTimeWithMinu)
	return GetTimestampByStrTime(r).Unix()
}

func GetDateTime(date string) (int64, int64) {
	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}
