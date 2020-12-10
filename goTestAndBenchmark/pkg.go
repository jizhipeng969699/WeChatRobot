package goTestAndBenchmark

import "fmt"

type Eat interface {
	Eat()
}
type Fly interface {
	Fly()
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (p Person) Eat() {
	fmt.Println(p.Name, `eat`)
}

type Student struct {
	Person
	Id int `json:"id"`
}

func (s Student) Study() error {
	fmt.Println(s.Name, "	study")
	return nil
}
