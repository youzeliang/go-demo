package main

import (
	"fmt"
	"os"
)

//var values = [5]int{10, 11, 12, 13, 14}

func main() {


	//// 版本D: 输出值:
	//for ix := range values {
	//	val := values[ix]
	//	go func() {
	//		fmt.Print(val, " ")
	//	}()
	//}
	//time.Sleep(1e9)



	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)

}
