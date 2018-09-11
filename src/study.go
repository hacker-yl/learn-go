/*package main

import "fmt"

var quit chan int //只开一个信道

func foo(id int){
	fmt.Println(id)
	quit <- 0 
}

func main(){
	count := 10
	quit = make(chan int,count)

	for i := 0;i < count;i ++{
		go foo(i)
	}
	for i := 0;i < count;i ++{
		<- quit
	}
}*/

package main

import "fmt"

var ch chan int = make(chan int)
func foo(id int){
	ch <- id
}

func main(){
	//有序输出
	go func(){
		for i := 0;i < 5;i ++{
			foo(i)
		}
	}()
	//无序输出
	/*for i := 0;i < 5;i ++{
		go foo(i)
	}*/
	
	for i := 0;i < 5;i ++{
		fmt.Println(<- ch)
	}
}
