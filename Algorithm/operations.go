package algorithm

import (
	"os"
	"slices"
	"strconv"
)

func InputValidation(numbers []string) []int {
	validInput := make(Stack, len(numbers))

	// Check and Convert the input to integers
	for i, nb := range numbers {
		num, err := strconv.Atoi(nb)
		if err != nil {
			Error_P()
		}
		validInput[i] = num
	}

	// Check for duplicates
	for i := 0; i < len(validInput)-1; i++ {
		for j := i + 1; j < len(validInput); j++ {
			if validInput[i] == validInput[j] {
				Error_P()
			}
		}
	}
	// Check if the input Numbers are already sorted
	if validInput.IsSorted() {
		os.Exit(0)
	}

	return validInput
}

func (ps *Stacks) Sort() {
	switch Size {
	case 2:
		ps.sortTwo()
	case 3:
		ps.sortThree()
	case 4:
		ps.sortFour()
	case 5:
		ps.sortFive()
	case 6:
		ps.sortSix()
	default:
		ps.MaxSorting()
	}
}

//------------------------------------------------------------------------------------

func (ps *Stacks) sortTwo() {
	if ps.StackA[0] > ps.StackA[1] {
		ps.Sa()
	}
}

func (ps *Stacks) sortThree() {
	if ps.StackA[0] > ps.StackA[1] && ps.StackA[1] > ps.StackA[2] {
		ps.Ra()
		ps.Sa()
	}
	if ps.StackA[1] > ps.StackA[2] {
		ps.Rra()
	}
	if ps.StackA[0] > ps.StackA[2] {
		ps.Ra()
	}
	if ps.StackA[0] > ps.StackA[1] {
		ps.Sa()
	}
}
func (ps *Stacks) sortFour() {
	for len(ps.StackA) > 3 {
		if slices.Min(ps.StackA) == ps.StackA[0] {
			ps.Pb()
		} else {
			ps.Ra()
		}
	}
	ps.sortThree()
	ps.Pa()
}

func (ps *Stacks) sortFive() {
	for len(ps.StackB) != 2 {
		if slices.Min(ps.StackA) == ps.StackA[0] {
			ps.Pb()
		} else {
			ps.Rra()
		}
	}
	ps.sortThree()
	ps.Pa()
	ps.Pa()
}

func (ps *Stacks) sortSix() {
	for len(ps.StackB) != 2 {
		if ps.StackA[0] > ps.StackA[1] {
			ps.Sa()
		} else if slices.Min(ps.StackA) == ps.StackA[0] {
			ps.Pb()
		} else {
			x := ps.StackA.FindIndex(slices.Min(ps.StackA))
			if x <= len(ps.StackA)/2 {
				ps.Ra()
			} else {
				ps.Rra()
			}
		}
	}

	ps.sortFour()
	ps.Pa()
	ps.Pa()
	ps.Pa()
}

//------------------------------------------------------------------------------------

func (ps *Stacks) MaxSorting() {
	Range = Size / 6

	X = (Size / 2) - Range
	Y = ((Size - 1) / 2) + Range

	Median = RefST[Size/2]

	Part_One(ps)
	Part_Two(ps)
}
