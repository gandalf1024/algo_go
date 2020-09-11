package main

import "fmt"

type circle struct {
	arr   []int
	size  int
	first int
	last  int
}

func New(sizeNum int) *circle {
	sizeNum++
	return &circle{
		arr:  make([]int, sizeNum),
		size: sizeNum,
	}
}

func (c *circle) isFull() bool {
	return (c.last+1)%c.size == c.first
}

func (c *circle) isEmpty() bool {
	return c.last == c.first
}

func (c *circle) add(i int) {
	if c.isFull() {
		return
	}
	c.arr[c.last] = i
	c.last = (c.last + 1) % c.size //当 7 % 7 则等于0,则反转
}

func (c *circle) get() int {
	if c.isEmpty() {
		return -1
	}
	ret := c.arr[c.first]
	c.first = (c.first + 1) % c.size
	return ret
}

func (c *circle) arrNum() int {
	return (c.last + c.size - c.first) % c.size
}

func (c *circle) show() {
	for i := c.first; i < c.first+c.arrNum(); i++ {
		fmt.Println(i%c.size, ":", c.arr[i%c.size])
	}
	fmt.Println("---------------------------------------------")
}

func main() {
	c := New(6)
	c.add(23)
	c.add(56)
	c.add(67)
	c.add(43)
	c.get()
	c.add(76)
	c.add(89)
	c.get()
	c.add(289)
	c.add(2839)
	c.get()
	c.add(2849)
	c.add(2859)
	c.show()

	//c.get()
	//c.get()
	//c.get()
	////c.get()
	////c.get()
	////c.get()
	//c.show()
	//
	//c.add(23)
	//c.add(56)
	//c.add(67)
	//c.add(43)
	//c.add(76)
	//c.add(89)
	//c.show()

}
