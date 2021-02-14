package stack_and_queue

type CQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() CQueue {
	return CQueue{s1: &Stack{}, s2: &Stack{}}
}

type Stack struct {
	d []int
}

func (s *Stack) Add(v int) {
	s.d = append(s.d, v)
}

func (s *Stack) Pop() int {
	if s.len() <= 0 {
		return -1
	}
	res := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return res
}

func (s *Stack) len() int {
	return len(s.d)
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Add(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s1.len() == 0 && this.s2.len() == 0 {
		return -1
	}

	if this.s2.len() > 0 {
		return this.s2.Pop()
	}
	for this.s1.len() > 0 {
		this.s2.Add(this.s1.Pop())
	}
	return this.s2.Pop()
}
