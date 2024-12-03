package main

import (
	"AOC2024/utils"
	"fmt"
)

func main() {
	column1, column2 := utils.ReadFile("../input.txt")

	similarityScore := 0
	for _, val1 := range column1 {
		occurances := 0
		for _, val2 := range column2 {
			if val1 == val2 {
				occurances += 1
			}
		}
		similarityScore += val1 * occurances
	}
	fmt.Printf("Similarity Score: %d\n", similarityScore)
}
