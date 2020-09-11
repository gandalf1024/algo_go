package simple

//简答队列

type queue struct {
	arr   []interface{}
	size  int
	first int
	last  int
}

func New(sizeNum int) *queue {
	if sizeNum < 1 {
		panic("size small")
	}

	return &queue{
		arr:   make([]interface{}, sizeNum),
		size:  sizeNum,
		first: 0,
		last:  sizeNum - 1,
	}
}

func (q *queue) isFull() bool {

	return false
}

func (q *queue) idEmpty() bool {

	return false
}

func (q *queue) put(i interface{}) bool {

	return false
}

func (q *queue) get() interface{} {

	return false
}
