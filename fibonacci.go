package main

import (
	"fmt"
)

func fibonacci(ch,quit chan int){
	x,y := 0,1
	for{
		select{
			case ch<-x:
				x,y = y,x+y
			case <-quit:
				fmt.Println("quit")
				return
		}
	}	
}

func whatever(ch,quit chan int){
	for i:=0;i!=10;i++{
		fmt.Println(<-ch)
	}
	quit<-0
}

func main(){
	ch, quit := make(chan int), make(chan int)
	
	go whatever(ch,quit)
	fibonacci(ch,quit)
	
}

