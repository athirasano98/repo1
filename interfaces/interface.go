package main
import "fmt"
type student_details interface{
	basic_details()
	location_details()
	mark_percentage()
}
type student struct{
	name string
	age int
	course string
	total_mark float64
	city string 
	pincode int

}
func (stud student)basic_details(){
	fmt.Println("basic details")
	fmt.Println("*************")
	fmt.Println("Name is",stud.name)
	fmt.Println("Age is",stud.age)
	fmt.Println("course is",stud.course)
}
func (stud student)location_details(){
	fmt.Println("\nlocation details")
	fmt.Println("****************")
	fmt.Println("city is",stud.city)
	fmt.Println("Pincode is",stud.pincode)
}
func (stud student)mark_percentage(){
	fmt.Println("\nmark percentage details")
	fmt.Println("***********************")
	fmt.Println("Percentage is",(stud.total_mark/500)*100,"%")
}
func full_details(s student_details){
	s.basic_details()
	s.location_details()
	s.mark_percentage()
}
func main(){
	var limit int
   fmt.Println("Enter the number of students")
   _,err:=fmt.Scan(&limit)
   if err != nil{
	panic(err)
   }
   var s=make([]student,limit)
   for i:=0;i<limit;i++{
	fmt.Println("Enter the student name")
	fmt.Scan(&s[i].name)
	fmt.Println("Enter the student age")
	_,err1:=fmt.Scan(&s[i].age)
	if err1 != nil{
		panic(err1)
	}
	fmt.Println("Enter the student course")
	fmt.Scan(&s[i].course)
	fmt.Println("Enter the student city")
	fmt.Scan(&s[i].city)
	fmt.Println("Enter the student pincode")
	_,err2:=fmt.Scan(&s[i].pincode)
	if err2 !=nil{
		panic(err2)
	}
	fmt.Println("Enter the total marks out of 500")
	_,err3:=fmt.Scan(&s[i].total_mark)
	if err3 !=nil{
		panic(err3)
	}
   }
   for i:=0;i<limit;i++{
		values:=s[i]
		full_details(values)
   } 
}
