package main

import (
	"fmt"
    "calculator/calc/add"
	"calculator/calc/divide"
	"calculator/calc/product"
	"calculator/calc/subtract"
)

func main(){
	var(
		num1 ,num2,ch int
	)
	fmt.Println("enter two numbers")
	_,err:=fmt.Scan(&num1,&num2)
	if err !=nil{
		panic(err)
	}
	fmt.Println("1:Addition \n 2:Division \n 3:Multiplication \n 4:Subtraction")
	fmt.Println("enter your choice")
	_,err1:=fmt.Scan(&ch)
	if err1 != nil{
		panic(err1)
	}
	switch ch {
	case 1:
		fmt.Println("addition ",add.Add(num1,num2))
	case 2:
		fmt.Println("Quotient ",divide.Divide(num1,num2))
	case 3:
		fmt.Println("product ",product.Product(num1,num2))
	case 4:
		fmt.Println("difference ",subtract.Sub(num1,num2))
	default:
		fmt.Println("invalid choice")
	}
}