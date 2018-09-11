package main
 
import "fmt"
//import "time"

func main(){
	ch := make(chan int,3)
	ch <- 1
	ch <- 2
	ch <- 3
	/*for v:= range ch{
		fmt.Println(v)
		if len(ch) <= 0{
			break
		}
	}*/
	//显示关闭信道
	close(ch)
	for v := range ch{
		fmt.Println(v)
	}
}

