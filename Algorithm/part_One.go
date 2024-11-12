package algorithm


func Part_One(ps *Stacks) {
	for !ps.StackA.IsEmpty() {
		LoopRange := X >= 0-Range || Y <= Size+Range
		for LoopRange {
			if X < 0 || Y > Size {
				X = 0
				Y = Size - 1
			}

			if ps.EmptyA() {
				return
			}

			if ps.StackB.SizeIs() == Y-X {
				X -= Range
				Y += Range
			}
			LoopRange = X >= 0-Range || Y <= Size+Range
		}
	}
}

func (ps *Stacks) EmptyA() bool {
	for index := Y; index >= X; index-- {
		if ps.StackA.IsEmpty() {
			return true
		}
		if ps.StackA.FirstIndexValue() == RefST[index] {
			ps.Pb()
			if ps.StackB.FirstIndexValue() <= Median {
				ps.Rb()
			}
			return false
		}
	}

	if !ps.StackA.IsEmpty() {
		ps.Ra()
	}

	return false
}
