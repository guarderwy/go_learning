package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 150 || age < 0 {
		fmt.Println("age error")
	} else {
		p.age = age
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	p.sal = sal
}

func (p *person) GetSal() float64 {
	return p.sal
}
