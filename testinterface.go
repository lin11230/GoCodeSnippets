package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)

}

func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (e *Employee) SayHi() {
	fmt.Println("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func (e *Student) SayHi() {
	fmt.Println("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.school, e.loan)
}

func (e *Student) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Humando interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

func main() {
	mike := Student{Human{"Mike", 25, "2312-5678"}, "MIT", 0.00}
	sam := Employee{Human{"Sam", 36, "4123-5678"}, "Golang Inc.", 1000}

	tom := Employee{Human{"Tom", 36, "41234-3456"}, "Things Ltd.", 5000}

	var i = &Humando
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

}
