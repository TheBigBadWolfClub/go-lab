package pkg

type Stack[S any] struct {
	mem []S
}

func (s *Stack[S]) Push(v S) {
	s.mem = append(([]S)(s.mem), v)
}

func (s *Stack[S]) PushAll(all ...S) {
	for _, v := range all {
		s.Push(v)
	}
}

func (s *Stack[S]) Pop() S {
	if s.IsEmpty() {
		var zerod S
		return zerod
	}
	v := s.mem[len(([]S)(s.mem))-1:]
	s.mem = s.mem[:len(([]S)(s.mem))-1]
	return v[0]
}

func (s *Stack[S]) IsEmpty() bool {
	return s.mem == nil || len(([]S)(s.mem)) == 0
}
