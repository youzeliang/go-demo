package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	//addOne := addInt(1)
	//fmt.Println(addOne())
	//fmt.Println(addOne())
	//fmt.Println(addOne())
	//
	//addTwo := addInt(2)
	//fmt.Println(addTwo())
	//fmt.Println(addTwo())
	//fmt.Println(addTwo())

	//str := "禾木课堂"
	//fmt.Println("before calling sendValue, str : ", str)
	//sendValue(str)
	//fmt.Println("after calling sendValue, str : ", str)
	//
	//fmt.Println("before calling sendAddress, str : ", str)
	//sendAddress(&str)
	//fmt.Println("after calling sendAddress, str: ", str)

	var wg sync.WaitGroup
	wg.Add(3)

	go doSomething(1, &wg)
	go doSomething(2, &wg)
	go doSomething(3, &wg)

	wg.Wait()
	log.Printf("finish all jobs\n")

}

func doSomething(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}


func addInt(n int) func() int {

	i := 0
	return func() int {
		i += n
		return i
	}
}


func sendValue(name string) {
	name = "hemuketang"
}

func sendAddress(name *string) {
	*name = "hemuketang"
}