package circle

//环形队列
type circle struct {
	maxSize int
	front   int //front 指向队列的第一个元素
	rear    int //rear 指向队列的最后一个元素的后一个位置
	arr     []int
}

func New(maxSizeNum int) *circle {
	return &circle{
		maxSize: maxSizeNum + 1,
		arr:     make([]int, maxSizeNum+1),
	}
}

func (c *circle) isFull() bool {
	return (c.rear+1)%c.maxSize == c.front
}

func (c *circle) isEmpty() bool {
	return c.rear == c.front
}

func (c *circle) addQueue(n int) {
	if c.isFull() {
		println("队列满，不能加入数据~")
		return
	}
	c.arr[c.rear] = n
	c.rear = (c.rear + 1) % c.maxSize
}

func (c *circle) getQueue() int {
	if c.isEmpty() {
		// 通过抛出异常
		panic("队列空，不能取数据")
	}
	value := c.arr[c.front]
	c.front = (c.front + 1) % c.maxSize
	return value
}

func (c *circle) showQueue() {
	if c.isEmpty() {
		println("队列空的，没有数据~~")
		return
	}
	for i := c.front; i < c.front+c.size(); i++ {
		println("arr[%d]=%d\n", i%c.maxSize, c.arr[i%c.maxSize])
	}
}

func (c *circle) size() int {
	return (c.rear + c.maxSize - c.front) % c.maxSize
}

func (c *circle) headQueue() int {
	if c.isEmpty() {
		panic("队列空的，没有数据~~")
	}
	return c.arr[c.front]
}
