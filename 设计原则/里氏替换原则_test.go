package 设计原则

import "testing"

type Adder interface {
	Add(a, b int) (int, error)
}

func Add(a, b int, adder Adder) (int, error) {
	return adder.Add(a, b)
}

type Ints struct {
}

func (i Ints) Add(a, b int) (int, error) {
	return a + b, nil
}

func TestAdder(t *testing.T) {
	i := new(Ints)

	r, err := Add(1, 2, i)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

}
