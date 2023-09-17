package main

import (
	"fmt"
)

type tuple struct{
	a int
	b int
}

//String() is an interface.
//type Stringer interface{
//	String() string
//}
func (t tuple) String() string{
	return fmt.Sprintf("(%d, %d)",t.a,t.b)
}

type ZeroDivisionException struct{
}
func (e *ZeroDivisionException) String() string{
	if(e == nil){
		return ""
	}
	return "cannot divide by zero!"
}

func (t tuple) division() (string,*ZeroDivisionException){
	if(t.b!=0){ 
		return fmt.Sprintf("%d",t.a/t.b), nil
	}
	return "",&ZeroDivisionException{}	
}

func main(){
	a := [2]tuple{
		tuple{1,2},
		tuple{3,0} }
	fmt.Println("a[0] =",a[0])	
	fmt.Println(a[0].division())
	fmt.Println("a[1] =",a[1])	
	fmt.Println(a[1].division())
}

