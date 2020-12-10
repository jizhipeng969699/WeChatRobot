package goTestAndBenchmark

import "testing"

func TestPerson_Eat(t *testing.T) {
	p := new(Student)

	p.Name  = `xiaoming`

	p.Eat()
	p.Study()

}
