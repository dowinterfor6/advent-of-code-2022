package set

var exists = struct{}{}

type set struct {
	// struct{} takes 0 bytes
	m map[string]struct{}
}

// * means pointer, dereferences memory
// & means to get the memory reference, not the actual value
func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

// (s *set) is a receiver, kind of like prototypal inheritance/monkey patching
func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}