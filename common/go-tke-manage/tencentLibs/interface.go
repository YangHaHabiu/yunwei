package tencentLibs

import "fmt"

type CommonInterface interface {
	Query() error
	Update() error
	Create() error
}

// 操作
func OperationFunc(object CommonInterface, opt string) {
	if opt == "query" {
		err := object.Query()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if opt == "update" {
		err := object.Update()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if opt == "create" {
		err := object.Create()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
