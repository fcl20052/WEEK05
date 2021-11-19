package main

import (
	"fmt"
	"time"
	"toytrix/toytrix"
)

var result chan string



func main() {
	result = make(chan string, 100)
	for {
		fmt.Println("Enter to add work,try fast or slow")
		fmt.Scanln()
		go pingByToytrix()
	}

}

func pingByToytrix() error {
	errors := toytrix.Do("get_result", func() error {

		time.Sleep(time.Second * 2)
		result <- "work done"
		fmt.Println("work complete")
		return nil
	}, func(e error) error {
		fmt.Printf("err:%v,do default \n", e)
		return nil
	})
	return errors
}
