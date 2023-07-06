package api

import "fmt"

type CommonInterface interface {
	DoApi() error
}

func RefreshCdn(obj CommonInterface) {
	err := obj.DoApi()
	if err != nil {
		fmt.Println(err)
		return
	}
}
