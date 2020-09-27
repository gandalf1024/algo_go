package simple

import (
	"fmt"
	"testing"
)

func Test_Push_Pop(t *testing.T) {
	s := New(8)
	s.Push("AA")
	s.Push(33)
	s.Push("BB")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
