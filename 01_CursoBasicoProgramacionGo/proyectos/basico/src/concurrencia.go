package main


import (
	"fmt"
	"time"
)

func say(text string){
	fmt.Println(text)
}

func main (){
	say("Hello")
	go say("world")
	time.Sleep(time.Second *1)
}