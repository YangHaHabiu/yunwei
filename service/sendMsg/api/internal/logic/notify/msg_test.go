package notify

import (
	"fmt"
	"testing"
)

func TestFeiShu(t *testing.T) {
	feishu := `{"secret": "c1ShVtfE1YBenlPMWgkK6d", "access_token": "16a64092-5a40-4219-adc5-813f7ae9b07f","app_id": "cli_a3031f4f9bbad00c", "app_secret": "rbOAqa0Qi5JlykIcz48BxeqC7JRYt640"}`

	feishu = ` {"app_id": "cli_a3031f4f9bbad00c", "secret": "t12LDidJcdCWYq0onfuyCe", "app_secret": "rbOAqa0Qi5JlykIcz48BxeqC7JRYt640", "access_token": "4e67b9e9-98fc-4918-9c42-99ecca3e4c74"}`
	//feishu = `{"app_id": "cli_a3031f4f9bbad00c", "app_secret": "rbOAqa0Qi5JlykIcz48BxeqC7JRYt640"}`
	feishuChan, err := SendToFeishuChan(feishu, false, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feishuChan.SendMsg("我是内容", "normal", "我是标题")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}

func TestEmail(t *testing.T) {
	json := `1111`
	SendToEmailChan("111@qq.com,222@qq.com", "111", "xxxx", json)

}

func TestSendToWechatChan(t *testing.T) {

	SendToWechatChan("1111", "", map[string]string{"content": "2222"})
}
