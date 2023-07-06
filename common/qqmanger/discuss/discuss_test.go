package discuss

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//mm := "http://10.10.88.217:5555"
	//
	//compile, _ := regexp.Compile(`//|:`)
	//split := compile.Split(mm, -1)
	//
	//fmt.Println(split[2])
	//
	//mm = `[CQ:reply,id=Ih787EzEgyIAAOhzTWWRkGJ0zRkB][CQ:at,qq=1287947042,text=@公共-运营-符方湛] [CQ:at,qq=1287947042,text=@公共-运营-符方湛]   装服`
	//r, _ := regexp.Compile(`\[CQ.*\]|[ ]+`)
	//allString := r.ReplaceAllString(mm, "")
	//fmt.Println(allString)

	//today := time.Now().Weekday()
	//today = time.Friday
	//day := time.Saturday - today
	//fmt.Printf("%d", day-1)

	//startTime := time.Now().Format("2006-01-02")
	//fmt.Println(startTime)

	fmt.Println(TodayNews())

}
