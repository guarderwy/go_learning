package model

type animal struct {
	Name string
	age  int
}

func NewAnimal(name string) *animal {
	return &animal{Name: name}
}

func (a *animal) SetAge(age int) {
	a.age = age
}

func (a *animal) GetAge() int {
	return a.age
}
