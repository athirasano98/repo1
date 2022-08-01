//program to convert all numbers to even


package main
import (
	"fmt"
	"sync"
)
  var wg sync.WaitGroup
func sub(s1[] int,ch chan int){
	defer wg.Done()
	for _,val:=range s1{
		if val%2 !=0{
			go inc(val,ch)
		}else{
			ch <- val	
		}		
	}
}
func inc(value int,ch chan int){
	ch <- value+1
}
func main(){
	wg.Add(1)
	var limit int
	fmt.Println("enter the limit")
	fmt.Scan(&limit)
	var slice1=make([]int,limit)
	var ch=make(chan int,limit)
	fmt.Println("enter the elements")
	for i:=0;i<limit;i++{
		fmt.Scan(&slice1[i])
	}
	go sub(slice1,ch)
	for i:=0;i<limit;i++{
		fmt.Println("value is",<-ch)
	} 
	wg.Wait()
}