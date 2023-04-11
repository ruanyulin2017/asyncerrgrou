package main

import (
	"errors"
	"time"

	"github.com/ruanyulin2017/asyncerrgroup"
)

type aegfunc func() error

func getGroupFunc(msg string, err error) aegfunc {
	return func() error {
		// Simulate program blocking
		if err == nil {
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(200 * time.Millisecond)
		}
		println("msg: ", msg)
		return err
	}
}
func main() {
	g := asyncerrgroup.NewAsyncGroup()

	func1 := getGroupFunc("func1", nil)
	func2 := getGroupFunc("func2", nil)
	funce := getGroupFunc("func3", errors.New("func err"))

	funcs := []aegfunc{func1, func2, funce}
	for _, f := range funcs {
		g.Run(f)
	}

	if err := g.Wait(); err != nil {
		println("err: ", err.Error())
	} else {
		println("no err")
	}
}
