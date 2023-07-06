package tool

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/util/gconv"
)

func TestMd5ByString(t *testing.T) {
	s := Md5ByString("AAA")
	t.Log(s)
}

func TestKrand(t *testing.T) {
	if gconv.Int64(1665366526) < (time.Now().Unix() - 300) {
		fmt.Println("开始时间小于当前时间")
	} else {
		fmt.Println(1111)
	}

}
