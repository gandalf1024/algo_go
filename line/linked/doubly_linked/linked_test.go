package doubly_linked

import (
	"fmt"
	"testing"
)

func Test_AddBySort(t *testing.T) {
	l := NewLinked()
	l.AddBySort(23, "zj")
	l.AddBySort(34, "kl")
	l.AddBySort(8, "ik")
	l.List()
}

func Test_Del(t *testing.T) {
	l := NewLinked()
	l.AddBySort(23, "zj")
	l.AddBySort(34, "kl")
	l.AddBySort(8, "ik")
	l.List()
	l.Del(23)
	fmt.Println("-------------------")
	l.List()
	l.Del(34)
	fmt.Println("-------------------")
	l.List()
}
