package stack

type Node struct {
	Next  *Node
	Value string
}

type Stack struct {
	Top, Bottom *Node
	Size        int
}

func NewStack() *Stack {
	s := &Stack{}
	return s
}

func (s *Stack) Length() int {
	return s.Size
}

func (s *Stack) Push(value string) {
	n := &Node{
		Value: value,
	}

	if s.Length() == 0 {
		s.Bottom = n
	} else {
		n.Next = s.Top
	}

	s.Top = n
	s.Size++
}

func (s *Stack) Pop() {
	if s.Length() == 0 {
		return
	}

	if s.Length() == 1 {
		s.Top = nil
		s.Bottom = nil
	} else {
		top := s.Top
		s.Top = top.Next
		top.Next = nil
	}

	s.Size--
}

func (s *Stack) Peek() Node {
	return *s.Top
}

func (s *Stack) ToString() string {
	outputString := ""
	curr := s.Top
	for curr != nil {
		outputString += curr.Value + " -> "
		curr = curr.Next
	}

	return outputString
}
