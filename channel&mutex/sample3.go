//displaying elements in the array using channel

package main
import "fmt"

func display(arr[] int,ch chan int){
	fmt.Println("enter the elememts")
	for i:=0;i<len(arr);i++{
		fmt.Scan(&arr[i])
	}
	for i:=0;i<len(arr);i++{
		ch<-arr[i]
	}
} 
func main(){
	var limit int
	fmt.Println("enter the limit")
	fmt.Scan(&limit)
	var arr =make([]int,limit)
	var ch =make(chan int,len(arr))
	go display(arr,ch)
	for i:=0;i<len(arr);i++{
		fmt.Println("the",i+1,"elements is",<-ch)
	}
}