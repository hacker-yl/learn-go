package main

import "fmt"
import "time"
//import "unsafe"

func produce(p chan<- int){
	for i := 0;i < 10;i ++{
		p <- i
		fmt.Println("Send:",i)
	}	
}

func consumer(c <-chan int){
	for i := 0;i < 10;i ++{
		v := <-c
		fmt.Println("receive:",v)
	}
}
func main(){
	ch := make(chan int)
	fmt.Printf("ch=%d\n",ch)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}


