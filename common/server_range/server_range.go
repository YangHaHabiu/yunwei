package server_range

import (
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"ywadmin-v3/common/gconvx"
)

//将连续和不连续区服处理成以逗号分割可相互转换.例:31/35-37/41 <---> [31,35,36,37,41]
func HandleServerIdRange(x interface{}) (result interface{}) {
	vtype := reflect.ValueOf(x)
	switch vtype.Kind() {
	case reflect.String:
		st := x.(string)
		compile := regexp.MustCompile(`,|\\`)
		st = compile.ReplaceAllString((st), ("/"))
		stList := strings.Split(st, "/")
		resultSli := make([]int, 0, len(stList))
		for _, v := range stList {
			if strings.Contains(v, "-") {
				start := strings.Split(v, "-")
				if !gconvx.IsNumeric(start[0]) {
					return nil
				}
				if gconvx.Int(start[0]) > gconvx.Int(start[1]) {
					return nil
				}
				for i := gconvx.Int(start[0]); i <= gconvx.Int(start[1]); i++ {
					resultSli = append(resultSli, i)
				}
			} else {

				if !gconvx.IsNumeric(v) {
					return nil
				}
				resultSli = append(resultSli, gconvx.Int(v))
			}

		}
		sort.Ints(resultSli)
		return resultSli
	case reflect.Slice:
		numList := x.([]int)
		if len(numList) == 0 {
			//return []string{}
			return ""
		}
		sort.Ints(numList)
		var start = numList[0]
		var end int
		var listEnd = numList[len(numList)-1]
		retList := make([]string, 0, len(numList))
		for i := 0; i < len(numList)-1; i++ {
			if numList[i+1] > numList[i]+1 {
				end = numList[i]
				if start == end {
					retList = append(retList, strconv.Itoa(start))
				} else {
					retList = append(retList, strconv.Itoa(start)+"-"+strconv.Itoa(end))
				}
				start = numList[i+1]
			}
		}
		if start == listEnd {
			retList = append(retList, strconv.Itoa(start))
		} else {
			retList = append(retList, strconv.Itoa(start)+"-"+strconv.Itoa(listEnd))
		}
		return strings.Join(retList, "/")
	}
	return
}

//两个列表取交集
func ListIntersection(list1, list2 []int) (list3 []int) {
	if len(list1) == 0 || len(list2) == 0 {
		return
	}
	for _, v1 := range list1 {
		for _, v2 := range list2 {
			if v1 == v2 {
				list3 = append(list3, v1)
			}
		}
	}
	return
}
