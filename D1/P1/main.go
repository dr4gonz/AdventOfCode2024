package main

import (
	"AOC2024/utils"
	"fmt"
	"slices"
)
func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

func main() {
	column1, column2 := utils.ReadFile("../input.txt")

	slices.Sort(column1)
	slices.Sort(column2)

	differences := 0
	for i := range column1 {
		differences += absDiffInt(column1[i], column2[i])
	}

	fmt.Printf("Result: %d\n", differences)
}
