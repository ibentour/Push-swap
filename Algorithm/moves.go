package algorithm

func (s *Stack) Push(number int) {
	*s = append([]int{number}, *s...)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	number := (*s)[0]
	*s = (*s)[1:]
	return number, true
}

func (s *Stack) Swap() {
	if len(*s) >= 2 {
		(*s)[0], (*s)[1] = (*s)[1], (*s)[0]
	}
}

func (s *Stack) Rotate() {
	if len(*s) > 1 {
		*s = append((*s)[1:], (*s)[0])
	}
}

func (s *Stack) ReverseRotate() {
	if len(*s) > 1 {
		*s = append([]int{(*s)[len(*s)-1]}, (*s)[:len(*s)-1]...)
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------

func (ps *Stacks) Pa() {
	if number, ok := ps.StackB.Pop(); ok {
		ps.StackA.Push(number)
		Moves = append(Moves, "pa")
	}
}

func (ps *Stacks) Pb() {
	if number, ok := ps.StackA.Pop(); ok {
		ps.StackB.Push(number)
		Moves = append(Moves, "pb")
	}
}

func (ps *Stacks) Sa() {
	ps.StackA.Swap()
	Moves = append(Moves, "sa")
}

func (ps *Stacks) Sb() {
	ps.StackB.Swap()
	Moves = append(Moves, "sb")
}

func (ps *Stacks) Ss() {
	ps.StackA.Swap()
	ps.StackB.Swap()
	Moves = append(Moves, "ss")
}

func (ps *Stacks) Ra() {
	ps.StackA.Rotate()
	Moves = append(Moves, "ra")
}

func (ps *Stacks) Rb() {
	ps.StackB.Rotate()
	Moves = append(Moves, "rb")
}

func (ps *Stacks) Rr() {
	ps.StackA.Rotate()
	ps.StackB.Rotate()
	Moves = append(Moves, "rr")
}

func (ps *Stacks) Rra() {
	ps.StackA.ReverseRotate()
	Moves = append(Moves, "rra")
}

func (ps *Stacks) Rrb() {
	ps.StackB.ReverseRotate()
	Moves = append(Moves, "rrb")
}

func (ps *Stacks) Rrr() {
	ps.StackA.ReverseRotate()
	ps.StackB.ReverseRotate()
	Moves = append(Moves, "rrr")
}
