/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: weibo_test
* @Date: 2021-9-15 9:46
 */
package crawler

import (
	"fmt"
	"testing"
)

func TestWeibo(t *testing.T) {
	test("3333", "1111", "222")
	//test("55555")
}

func test(m string, aa ...string) {
	fmt.Println(m, aa)
	if len(aa) > 0 {
		fmt.Println(aa[0])
	}
}
