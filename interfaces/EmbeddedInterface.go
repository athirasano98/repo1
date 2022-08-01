package main
import (
	"fmt"
	"math"
)

type shape interface{
	area()float64
}
type object interface{
	shape
	volume()float64
}
type square struct{
	side float64
}
type cube struct{
	square
}
func(s square) area()float64{
	return math.Pow(s.side,2)
}
func(c cube) volume()float64{
	return math.Pow(c.side,3)
}
func totalArea(sh ...shape)float64{
	var total float64
	for _,val:=range sh{
		total+=val.area()
	}
	return total
}
func totalVolume(sh ...object)(float64,float64){
	var Volumetotal,Areatotal float64
	for _,val:=range sh{
		Volumetotal+=val.volume()
	}
	for _,val:=range sh{
		Areatotal+=val.area()
	}
	return Volumetotal,Areatotal
}
func main(){
	c:=square{2}
	v:=cube{square{3.0}}//v:=cube{square:square{side:3.0}}
	fmt.Println(totalVolume(v))
	fmt.Println(totalArea(c))
}