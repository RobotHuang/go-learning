package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Age int
	Name string
}

type myInt int

func main() {
	//var stu Student
	/* stu = Student{
		10,
		"abc",
	} */
	var a myInt
	a = 2
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println(t.Kind())
	fmt.Println(v)
}