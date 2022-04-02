package pkg

type Queue[Q any] struct {
	mem []Q
}

func (s *Queue[Q]) Enqueue(v Q) {
	s.mem = append(([]Q)(s.mem), v)
}

func (s *Queue[Q]) EnqueueAll(all ...Q) {
	for _, v := range all {
		s.mem = append(([]Q)(s.mem), v)
	}
}

func (s *Queue[Q]) Dequeue() Q {
	if s.IsEmpty() {
		var zerod Q
		return zerod
	}

	v := s.mem[0]
	s.mem = s.mem[1:]
	return v
}

func (s *Queue[Q]) IsEmpty() bool {
	return s.mem == nil || len(([]Q)(s.mem)) == 0
}
