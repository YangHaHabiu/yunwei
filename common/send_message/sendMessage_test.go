/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: sendMessage_test.go
* @Date: 2021-7-15 15:25
 */
package send_message

import (
	"fmt"
	"testing"
	"time"
)

func TestMessage_Send(t *testing.T) {

	// var msg MessageInterface
	// msg = NewMessage("xxx游戏热更",
	// 	"admin",
	// 	"000",
	// 	"33",
	// 	"44",
	// 	"55",
	// 	"巨兽战争",
	// 	"333",
	// 	0,
	// 	111,
	// 	1,
	// 	"",
	// 	"")
	// //fmt.Println(msg.Send("797252659", "group"))
	// send := msg.Send("1544553049", "discuss")
	// fmt.Println(send, "*****")
	// fmt.Println(len(send))
	formatTemplate := "2006-01-02"
	startTime := time.Now().AddDate(0, 0, -1).Format(formatTemplate)
	endTime := time.Now().Format(formatTemplate)
	fmt.Println(startTime, endTime)
}
