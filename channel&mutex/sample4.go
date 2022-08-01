package main
import (
	"fmt"
	"time"
)

func portal1(ch chan string){
	time.Sleep(5*time.Second)
	ch <- "welcome to portal1"
}
func portal2(ch chan string){
	time.Sleep(3*time.Second)
	ch <- "welcome to portal2"
}

func main(){
	var channel1=make(chan string)
	var channel2=make(chan string)
	go portal1(channel1)
	go portal2(channel2)
    select{
	case op1:=<-channel1:
		fmt.Println(op1)
	case op2:=<-channel2:
		fmt.Println(op2)
	}
}