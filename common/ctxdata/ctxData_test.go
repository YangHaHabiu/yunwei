package ctxdata_test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestXxx(t *testing.T) {
	lists := []string{`【中国农业银行】李进于02月22日10:10向您尾号2579账户完成转账交易人民币200.00，余额1620.42。备注信息： 人民币1000.00`,
		`【邮储银行】23年02月10日15:12赖贵强账户8436向您尾号778账户他行汇入,收入金额100.00元,余额5023.02元。`,
		`您的借记卡账户长城电子借记卡，于01月19日手机银行支取(跨行支付)人民币500.00元,交易后余额4493.95【中国银行】`,
		`您尾号8175的储蓄卡账户8月29日20时3分转帐支取支出人民币10000.00元,活期余额16713.30元。[建设银行]`,
		`30日23:04账户*7961*汇款汇出支出10000.00元，余额83.79元。对方户名:林国斌，对方银行:中国建设银行总行(不受理个人业务)，对方账号:*7717*。[兴业银行]`,
	}

	//匹配金额
	r := regexp.MustCompile("(人民币|金额|支出)[1-9]\\d*.\\d*")
	//匹配人名
	r2 := regexp.MustCompile("】([\u4e00-\u9fa5]+)于|:\\d+([\u4e00-\u9fa5]+)账户|对方户名:([\u4e00-\u9fa5]+)，")
	for _, s := range lists {
		fmt.Println(s)
		s2 := r2.FindAllStringSubmatch(s, -1)
		b := r.FindAllString(s, 1)
		for _, v := range b {
			r3 := regexp.MustCompile(`人民币|金额|支出`)
			s3 := r3.Split(v, -1)
			if len(s3) > 1 {
				var name string
				if len(s2) > 0 {
					name = s2[0][2]
					if name == "" {
						name = s2[0][1]
						if name == "" {
							name = s2[0][3]
						}
					}
				} else {
					name = "未知"
				}
				fmt.Println(name, "支付金额：", s3[1])
			}
			fmt.Println("------------------------------")
		}

	}
}
