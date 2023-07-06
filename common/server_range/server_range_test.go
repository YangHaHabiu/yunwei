package server_range

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"net/url"
	"strings"
	"testing"
)

//单元测试
func TestHandleServerIdRange(t *testing.T) {
	t.Log(HandleServerIdRange(gconv.SliceInt(HandleServerIdRange("1-37/39/47-49/66-67/137"))))
	//t.Log(HandleServerIdRange([]int{8, 1, 2, 3, 4, 5, 9}))
	//t.Log(strings.Join(gconvx.SliceStr([]int{1, 2, 3, 4, 5, 9}), ","))

	//t.Log(strings.Join(gconvx.SliceStr(HandleServerIdRange("50/31/35-37/41")), ","))

	//mm := []string{
	//	"1-37/47-49/66-67", "38-46/50-65/68-71/74-79/82-83", "72-73/80-81/84-103/110-113",
	//}
	//fmt.Println(strings.Join(mm, "/"))
	//
	//str := gconvx.SliceInt(HandleServerIdRange(strings.Join(mm, "/")))
	//fmt.Println(str)
	//idRange := HandleServerIdRange(str)
	//fmt.Println(idRange)
	//
	////t.Log(gconv.SlicemysqlStr([]int{1, 2, 3, 4, 5, 9}))
	//
	////t.Log(ListIntersection([]int{1, 2, 3, 4, 5, 9}, []int{11, 12}))
}

//性能基准测试1
func BenchmarkHandleServerIdRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HandleServerIdRange([]int{1, 2, 3, 4, 5, 9})

	}
}

func TestListIntersection2(t *testing.T) {
	var urlStr string = "张清贤"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:", escapeUrl)

	enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Println("解码:", enEscapeUrl)
}

//性能基准测试2
func BenchmarkHandleServerIdRange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HandleServerIdRange("50/31/35-37/41")

	}
}

func TestAbc(t *testing.T) {

	intersection := ListIntersection(gconv.SliceInt([]int64{1, 2, 3, 4}), gconv.SliceInt([]int64{1, 7}))
	fmt.Println(intersection)
	fmt.Println(len(intersection))

}

func TestListIntersection(t *testing.T) {

	var tmp string
	for _, v := range strings.Split("jszza 1-2,jsszb 3-9,jsszc,jszzm 1-11/13", ",") {
		plantinfo := strings.Split(v, " ")
		if len(plantinfo) == 1 {
			tmp += fmt.Sprintf("( platform='%s' ) or ", plantinfo[0])
		} else {
			inputRange := HandleServerIdRange(plantinfo[1])
			if inputRange == nil || len(inputRange.([]int)) == 0 {
				return
			}
			tmp += fmt.Sprintf("( platform='%s' and sid in (%s) ) or ", plantinfo[0], strings.Join(gconv.SliceStr(inputRange), ","))
		}

	}
	fmt.Println(fmt.Sprintf("and (%s)", tmp[:strings.LastIndex(tmp, "or")]))
}
