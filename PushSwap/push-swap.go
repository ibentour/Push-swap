package main

import (
	"fmt"
	"os"
	al "pushswap/Algorithm"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	Numbers := strings.Fields(os.Args[1])
	PS := new(al.Stacks)
	PS.StackA = append([]int{}, al.InputValidation(Numbers)...)
	PS.StackB = make(al.Stack, 0)
	al.Size = len(PS.StackA)
	al.RefST = make([]int, al.Size)
	copy(al.RefST, PS.StackA)
	al.SortReff(al.RefST)

	PS.Sort()

	for _, move := range al.Moves {
		fmt.Println(move)
	}
	// fmt.Println(PS.StackA.IsSorted()) // Print StackA
	// fmt.Println(PS.StackA) // Print StackA
}
