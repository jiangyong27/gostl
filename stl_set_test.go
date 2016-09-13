package gostl

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add(1)
	s.Add(2)
	s.Clear()
	if s.IsEmpty() {
		fmt.Println("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	s.Add(1)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println("list of all items", s.List())
}
