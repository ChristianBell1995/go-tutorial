package main

import (
	"fmt"
	"time"
)

func p(s string) {
	fmt.Println(s)
}

func main() {
	p("These")
	p("are")
	p("executed")
	p("in")
	p("order")
	p("**************")

	go p("These")
	go p("are")
	go p("executed")
	go p("in")
	go p("their")
	go p("own")
	go p("threads")

	time.Sleep(time.Second)
	fmt.Println("done")
}
