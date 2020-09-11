package singly_linked

import (
	"fmt"
	"testing"
)

//单向链表测试
func Test_Linked(t *testing.T) {

	l := New()
	l.Add(99, "cv")
	l.Add(34, "yu")
	l.Add(34, "yu-new") //相同编号做覆盖处理
	l.Add(23, "er")
	l.Add(56, "kl")
	l.Add(78, "sd")
	l.Add(89, "qw")

	l.RmByName("cv")
	l.RmByName("yu-new")
	l.RmByName("er")
	l.RmByName("kl")
	//l.RmByName("sd")
	l.RmByName("qw")

	l.List()
	println(l.Size())

}

func Test_Reverse(t *testing.T) {
	l := New()
	l.Add(99, "cv")
	l.Add(34, "yu")
	l.Add(34, "yu-new") //相同编号做覆盖处理
	l.Add(23, "er")
	l.Add(56, "kl")
	l.Add(78, "sd")
	l.Add(89, "qw")
	l.List()
	fmt.Println("------------------")
	l.Reverse()
	l.List()

}

func Test_ReversePrint(t *testing.T) {
	l := New()
	l.Add(99, "cv")
	l.Add(34, "yu")
	l.Add(34, "yu-new") //相同编号做覆盖处理
	l.Add(23, "er")
	l.Add(56, "kl")
	l.Add(78, "sd")
	l.Add(89, "qw")
	l.List()
	fmt.Println("=============================")
	l.ReversePrint()
}
