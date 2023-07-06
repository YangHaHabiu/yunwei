package xdate

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	start, end := GetWeekDate(-7)
	fmt.Println(start, end)
	dates := GetBetweenDates(start, end)
	fmt.Println(dates)

	fmt.Println("end :", b())
	fmt.Println("return:", *(c()))
	retrunAndDefer()
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func retrunAndDefer() int {
	defer deferFunc()

	return returnFunc()
}
