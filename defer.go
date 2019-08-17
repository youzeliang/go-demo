package main

import (
	"errors"
	"github.com/aofei/air"
)

func main()  {


	air.Default.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("Hello, 世界")
	})
	air.Default.Serve()

}

func out()  {
	innerFnc()
}

func innerFnc()  {
	panic(errors.New("te"))
}