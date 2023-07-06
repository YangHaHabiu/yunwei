/*
@Time : 2022-4-18 15:57
@Author : acool
@File : xsort
*/
package xsort

import "ywadmin-v3/service/admin/rpc/adminclient"

type OneList []*adminclient.MenuListTree

func (this OneList) Len() int {
	return len(this)
}

func (this OneList) Less(i, j int) bool {
	return this[i].OrderNum < this[j].OrderNum
}

func (this OneList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
