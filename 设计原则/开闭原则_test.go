package 设计原则

import (
	"fmt"
	"testing"
)

type person struct {
}

func (p person) Sort() {
	fmt.Println(`person sort`)
}

type student struct {
	person
}

func (s student) Sort() {
	fmt.Println(`student sort`)
}

func TestSort(t *testing.T) {
	s := &student{person{}}

	s.Sort()
	s.person.Sort()
}
