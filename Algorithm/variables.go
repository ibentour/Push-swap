package algorithm

import (
	"fmt"
	"os"
	"sort"
)

type Stack []int

type Stacks struct {
	StackA Stack
	StackB Stack
}

var (
	Moves         []string
	RefST         Stack
	Range         int
	Size          int
	MaxIndex      int
	Queue         int
	ReadyToRotate bool
	Median        int
	X             int
	Y             int
)

func SortReff(slc []int) {
	sort.SliceStable(slc, func(a, b int) bool {
		return slc[a] < slc[b]
	})
}

func (slc Stack) IsSorted() bool {
	return sort.SliceIsSorted(slc, func(a, b int) bool {
		return slc[a] < slc[b]
	})
}

func (slc Stack) FindIndex(num int) int {
	for i, v := range slc {
		if v == num {
			return i
		}
	}
	return -1
}

func (slc Stack) FirstIndexValue() int {
	return slc[0]
}

func (slc Stack) LastIndexValue() int {
	return slc[len(slc)-1]
}

func (slc Stack) IsEmpty() bool {
	return len(slc) == 0
}

func (slc Stack) SizeIs() int {
	return len(slc)
}

func Error_P() {
	fmt.Println("Error")
	os.Exit(1)
}
