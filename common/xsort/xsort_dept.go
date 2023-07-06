/*
@Time : 2022-4-18 15:57
@Author : acool
@File : xsort
*/
package xsort

import (
	"ywadmin-v3/service/admin/rpc/adminclient"
)

type DeptList []*adminclient.DeptListData

func (this DeptList) Len() int {
	return len(this)
}

func (this DeptList) Less(i, j int) bool {
	return this[i].OrderNum < this[j].OrderNum
}

func (this DeptList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
