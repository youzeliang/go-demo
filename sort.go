package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonWrapper struct {					//注意此处
	people [] Person
	by func(p, q * Person) bool
}

func (pw PersonWrapper) Len() int {    		// 重写 Len() 方法
	return len(pw.people)
}
func (pw PersonWrapper) Swap(i, j int){     // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper) Less(i, j int) bool {    // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}


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


	intList := [] int {2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}
