package algorithm


func Part_Two(ps *Stacks) {
	MaxIndex = Size - 1
	Queue = 0

	for !ps.StackB.IsEmpty() {
		if ps.sortBackToA() {
			break
		}
	}

	for Queue != 0 {
		ps.Rra()
		Queue--
	}
}

func (ps *Stacks) sortBackToA() bool {
	ReadyToRotate = true

	if ps.StackB.FirstIndexValue() == RefST[MaxIndex] {
		ps.Pa()
		MaxIndex--
		ReadyToRotate = false
	} else if Queue != 0 {
		if ps.StackA.LastIndexValue() == RefST[MaxIndex] {
			ps.Rra()
			MaxIndex--
			Queue--
		}
	}

	SendToQueue(ps)

	return false
}


func SendToQueue(ps *Stacks) {

	if ReadyToRotate && (Queue == 0 ||
		(Queue != 0 && ps.StackB.FirstIndexValue() > ps.StackA.LastIndexValue())) {
		ps.Pa()
		ps.Ra()
		Queue++
		ReadyToRotate = false
	}

	if ReadyToRotate {
		if ps.StackB.FindIndex(RefST[MaxIndex]) <= ps.StackB.SizeIs()/2 {
			ps.Rb()
		} else {
			ps.Rrb()
		}
	}
}
