package main

import (
	"fmt"
	"time"
)

func timeLoop(c chan int64) {
	for i := 0; i <= 20; i++ {
		c <- int64(i)
		time.Sleep(time.Second * 1)
	}
}


func say(text string, c chan<- string) {
	c <- text
}


func main() {
	c := make(chan int64, 1)

	go timeLoop(c)

	for {
		fmt.Println(<-c)
	}

	//Otro ejemplo 

	d := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bay", d)

	fmt.Println(<-d)	
}
