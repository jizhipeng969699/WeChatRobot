package simple_factory

import (
	"fmt"
	"testing"
)

func TestNewFood(t *testing.T) {
	s := NewFood(AppleType)
	t.Log(s.GetSource())

	m := NewFood(MeatType)
	t.Log(m.GetSource())


	fmt.Println(AppleType)
	fmt.Println(MeatType)
}
