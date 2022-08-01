package main

import (
	"fmt"
	"src/mypackages/geometry"
	"src/mypackages/package1"
	"src/mypackages/package2"
)

func main() {

	fmt.Println(package1.Hello())
	fmt.Println("Add:", package2.Add(1, 2))
	fmt.Println("Sub:", package2.Sub(10, 4))
	fmt.Println("Circle")
	fmt.Println("\tArea:", geometry.Ar(10))
	fmt.Println("\tPerimeter:", geometry.Per(5))
	fmt.Println("Rectangle")
	fmt.Println("\tArea:", geometry.Area(4, 6))
	fmt.Println("\tPerimeter:", geometry.Perimeter(5, 9))

}