/*
@Time : 2022-3-31 17:21
@Author : acool
@File : xfilters
*/
package xfilters

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/util/gconv"
)

// 查询条件语句
func Xfilters(filters ...interface{}) string {
	l := len(filters)
	resList := make([]string, 0)
	if l%2 == 0 && l > 0 {
		var tmps string
		for k := 0; k < l; k += 2 {
			biaodashi := strings.Split(filters[k].(string), "__")
			if len(biaodashi) < 1 {
				return ""
			}
			switch filters[k+1].(type) {
			case int, int64, string:
				tmp := strings.TrimSpace(gconv.String(filters[k+1]))
				if len(tmp) != 0 && tmp != "0" {
					if biaodashi[1] == "like" {
						tmps = fmt.Sprintf("%s %s '%%%s%%'", biaodashi[0], biaodashi[1], tmp)
					} else if biaodashi[1] == "rlike" {
						tmps = fmt.Sprintf("%s %s '%%%s'", biaodashi[0], "like", tmp)
					} else if biaodashi[1] == "nrlike" {
						tmps = fmt.Sprintf("%s %s '%%%s'", biaodashi[0], "not like", tmp)
					} else if biaodashi[1] == "regexp" {
						tmp = strings.ReplaceAll(tmp, ",", "|")
						tmps = fmt.Sprintf("%s %s '%s'", biaodashi[0], biaodashi[1], tmp)
					} else if biaodashi[1] == "in" || biaodashi[1] == "not in" {
						tmps = fmt.Sprintf("%s %s ('%s')", biaodashi[0], biaodashi[1], strings.Join(strings.Split(tmp, ","), "','"))
					} else if biaodashi[1] == "or" {
						mm := make([]string, 0)
						var tt string
						if len(biaodashi) > 2 {
							tt = biaodashi[2]
							if biaodashi[2] == "regexp" {
								tmp = strings.ReplaceAll(tmp, ",", "|")
							} else if biaodashi[2] == "like" {
								tmp = "'%%%s%%'"
							}
							for _, v := range strings.Split(biaodashi[0], "@") {
								mm = append(mm, fmt.Sprintf("%s %s '%s'", v, tt, tmp))
							}
							tmps = "(" + strings.Join(mm, " or ") + ")"
						}

					} else if biaodashi[1] == "expect" {
						tmps = fmt.Sprintf(" ( unix_timestamp(now()) - %s >=0 and %s <> 0 )", biaodashi[0], biaodashi[0])
					} else if biaodashi[1] == "range" {
						if len(strings.Split(tmp, ",")[0]) != 0 {
							tmps = fmt.Sprintf("(DATE_FORMAT(%s,'%%Y-%%m-%%d') BETWEEN '%s' and '%s')", biaodashi[0], strings.Split(tmp, ",")[0], strings.Split(tmp, ",")[1])
						}
					} else if biaodashi[1] == "frange" {
						if len(strings.Split(tmp, ",")[0]) != 0 {
							tmps = fmt.Sprintf("(FROM_UNIXTIME(%s,'%%Y-%%m-%%d') BETWEEN '%s' and '%s')", biaodashi[0], strings.Split(tmp, ",")[0], strings.Split(tmp, ",")[1])
						}
					} else if biaodashi[1] == "xrange" {
						if len(strings.Split(tmp, ",")[0]) != 0 {
							tmps = fmt.Sprintf("(%s BETWEEN '%s' and '%s')", biaodashi[0], strings.Split(tmp, ",")[0], strings.Split(tmp, ",")[1])
						}
					} else {
						switch filters[k+1].(type) {
						case int64, int:
							tmps = fmt.Sprintf("%s %s %d", biaodashi[0], biaodashi[1], gconv.Int64(filters[k+1]))
						default:
							tmps = fmt.Sprintf("%s %s '%s'", biaodashi[0], biaodashi[1], tmp)

						}

					}
				} else {
					tmps = ""
				}
			}
			if len(strings.TrimSpace(tmps)) != 0 {
				resList = append(resList, tmps)
			}
		}
		return strings.Join(resList, " and ")
	} else {
		return ""
	}

}
