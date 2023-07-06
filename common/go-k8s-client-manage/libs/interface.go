package libs

import "fmt"

//公共接口
type CommonInterface interface {
	Create() error
	Update() error
	Query() error
	Delete() error
	Redeployment() error
}

func OperationK8sClient(objcet CommonInterface, oper string) error {
	if oper == "create" {
		err := objcet.Create()
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else if oper == "update" {
		err := objcet.Update()
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else if oper == "query" {
		err := objcet.Query()
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else if oper == "delete" {
		err := objcet.Delete()
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else if oper == "redeployment" {
		err := objcet.Redeployment()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
