// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconvx

import "strings"

// SliceStr is alias of Strings.
func SliceStr(i interface{}) []string {
	return Strings(i)
}

// SlicemysqlStr
func SlicemysqlStr(x []int, master ...int) (result string) {
	tmp := make([]string, 0, len(x))
	for _, v := range x {
		if master[0] == v {
			continue
		}
		tmp = append(tmp, "'"+String(v)+"'")
	}

	return strings.Join(tmp, ",")
}

// SlicemysqlStr
func SliToMysqlStr(str string) string {
	x := strings.Split(str, ",")
	tmp := make([]string, 0, len(x))
	for _, v := range x {
		tmp = append(tmp, "'"+String(v)+"'")
	}

	return strings.Join(tmp, ",")
}

// Strings converts <i> to []string.
func Strings(i interface{}) []string {
	if i == nil {
		return nil
	}
	if r, ok := i.([]string); ok {
		return r
	} else {
		var array []string
		switch value := i.(type) {
		case []int:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int8:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int16:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int32:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int64:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint8:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint16:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint32:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint64:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []bool:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []float32:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []float64:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []interface{}:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case [][]byte:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		default:
			return []string{String(i)}
		}
		return array
	}
}
