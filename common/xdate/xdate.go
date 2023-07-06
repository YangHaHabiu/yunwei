package xdate

import "time"

var (
	timeFormatTpl = "20060102"
)

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func GetWeekDate(arr ...int) (start, end string) {
	var count int
	if len(arr) != 0 {
		count = arr[0]
	} else {
		count = -7
	}

	now := time.Now()
	date := now.AddDate(0, 0, count)
	format := date.Format(timeFormatTpl)
	if count < 0 {
		start = format
		end = now.Format(timeFormatTpl)
	} else {
		end = format
		start = now.Format(timeFormatTpl)
	}

	return
}
