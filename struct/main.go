package main

import "fmt"

type person struct {
	name string
	age int
}

type abc interface {
	getName() string
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age: age,
	}
}

func (p *person) getName() string{
	return p.name
}

func retValue(a int) *int {
	return &a
}

type People interface {
	Speak(string) string
}

type Student struct {}

func (s *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "sb"
	} else {
		talk = "hello"
	}
	return
}

func main() {

	a := newPerson("abc", 12)
	b := newPerson("cde", 13)
	fmt.Printf("%T\n", a)
	fmt.Println((*b).name)
	c := person{"abc", 25}
	var d abc
	d = &c
	d.getName()
	var p People = &Student{}
	fmt.Println(p.Speak("hello"))
}