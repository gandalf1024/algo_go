package circle

import (
	"testing"
)

func Test_circle(t *testing.T) {
	c := New(6)
	c.addQueue(99)
	c.addQueue(22)
	c.addQueue(33)
	c.addQueue(44)
	c.addQueue(55)
	c.addQueue(66)
	c.showQueue()

	println(c.getQueue())
	println(c.getQueue())
	println(c.getQueue())
	println(c.getQueue())
	println(c.getQueue())
	println(c.getQueue())
	c.showQueue()

	c.addQueue(99)
	c.addQueue(22)
	c.addQueue(33)
	c.addQueue(44)
	c.addQueue(55)
	c.addQueue(66)
	c.showQueue()

}
