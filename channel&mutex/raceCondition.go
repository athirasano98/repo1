package main
import (
	"fmt"
	"sync"
)
var wg=&sync.WaitGroup{}
var mutex=&sync.Mutex{}
func main() {
	var num = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup,mutex *sync.Mutex){
		mutex.Lock()
		fmt.Println("first")
	   	num=append(num,1)
		mutex.Unlock()
	   	wg.Done()
	}(wg,mutex)
	go func(wg *sync.WaitGroup,mutex *sync.Mutex){
		mutex.Lock()
		fmt.Println("second")
	   	num=append(num,2)
		mutex.Unlock()
	   	wg.Done()
	}(wg,mutex)
	go func(wg *sync.WaitGroup ,mutex *sync.Mutex){
		mutex.Lock()
		fmt.Println("third")
	   	num=append(num,3)
		mutex.Unlock()
	  	wg.Done()
	}(wg,mutex)
	wg.Wait()
	fmt.Println(num)
}

