package main  
import (  
    "fmt"
    "sync"
    )
var x  = 0  
func increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock()
    x = x + 1
    fmt.Println("The value of x :", x)
    m.Unlock()
    wg.Done()   
}
func main() {  
    var w sync.WaitGroup
    var m sync.Mutex
    w.Add(2)
        go increment(&w,&m)
        go increment(&w, &m)
    w.Wait()

}