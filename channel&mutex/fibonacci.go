package main
import (
    "fmt"
    "sync"
)
var wg=&sync.WaitGroup{}
func fibonacci( n int ,ch chan int,wg *sync.WaitGroup){
    defer wg.Done()
    var x,y=0,1
    for i:=0;i<n;i++{
        ch <- x
        x,y = y,x+y
    }
    close(ch)
}
func main(){
    wg.Add(1)
    var limit int
    fmt.Println("enter the limit")
    fmt.Scan(&limit)
    var ch =make(chan int,limit)
    go fibonacci(cap(ch),ch,wg)
    fmt.Println("The series is:")
    for i:=range ch{
        fmt.Println(i)
    }
    wg.Wait()
}