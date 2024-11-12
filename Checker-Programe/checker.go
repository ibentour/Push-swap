package main

import (
	"bufio"
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

	executeInstructions(PS)

	if PS.StackA.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func executeInstructions(PS *al.Stacks) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		instruction := scanner.Text()
		switch instruction {
		case "pa":
			PS.Pa()
		case "pb":
			PS.Pb()
		case "sa":
			PS.Sa()
		case "sb":
			PS.Sb()
		case "ss":
			PS.Ss()
		case "ra":
			PS.Ra()
		case "rb":
			PS.Rb()
		case "rr":
			PS.Rr()
		case "rra":
			PS.Rra()
		case "rrb":
			PS.Rrb()
		case "rrr":
			PS.Rrr()
		case "":
			continue
		default:
			fmt.Println("Error")
			os.Exit(1)
		}
	}
}