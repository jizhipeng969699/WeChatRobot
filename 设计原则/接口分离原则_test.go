package 设计原则

import (
	"fmt"
	"testing"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Eater interface {
	Eat() error
}
type Flyer interface {
	Fly() error
}
type Studyer interface {
	Study() error
}

type Personer interface {
	Eater
	Studyer
}

type Person struct {
	Name string
}

func (p Person) Eat() error {
	fmt.Println(p.Name, `	eat`)
	return nil
}

func (p Person) Study() error {
	fmt.Println(p.Name, `	study`)
	return nil
}

func TestName(t *testing.T) {
	p := new(Person)
	p.Name = `xiaoming`

	err := p.Study()
	if err != nil {
		t.Fatal(err)
	}

	err = p.Eat()
	if err != nil {
		t.Fatal(err)
	}

}
