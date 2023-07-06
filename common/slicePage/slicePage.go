package slicePage

import "math"

func SlicePage(page, pageSize, nums int64) (sliceStart, sliceEnd int64) {
	// 定义page和size的默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	// 如果pageSize大于num（切片长度）, 那么sliceEnd直接返回num的值
	if pageSize > nums {
		return 0, nums
	}
	// 总页数计算，math.Ceil 返回不小于计算值的最小整数（的浮点值）
	pageCount := int64(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize
	// 如果页总数比sliceEnd小，那么就把总数赋值给sliceEnd
	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}
