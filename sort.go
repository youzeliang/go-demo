package main

import (
	"fmt"
	"sort"
)

func main() {

	test1()

}

func test1() {

	//整数排序
	number := []int{1, 23, 9, 4}
	sort.Ints(number)
	fmt.Println("int:", number)

	//检测切片是否已经排序完成
	done := sort.IntsAreSorted(number)
	fmt.Println("sorted:", done)
}
