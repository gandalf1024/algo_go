package simple

type stack struct {
	arr []interface{}
	cru int
}

func New(size int) *stack {
	return &stack{arr: make([]interface{}, size)}
}

func (s *stack) Push(i interface{}) {
	if s.cru < len(s.arr) {
		s.arr[s.cru] = i
		s.cru++
	}
}

func (s *stack) Pop() interface{} {
	if s.cru == 0 {
		return nil
	}
	s.cru--
	i := s.arr[s.cru]
	return i
}
