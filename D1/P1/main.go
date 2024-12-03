package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)
func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// The text input is 1000 lines long
	column1 := make([]int, 1000)
	column2 := make([]int, 1000)
	count := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		first, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		column1[count] = first

		second, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		column2[count] = second
		count++
	}

	slices.Sort(column1)
	slices.Sort(column2)

	differences := 0
	for i := range column1 {
		differences += absDiffInt(column1[i], column2[i])
	}

	fmt.Printf("Result: %d\n", differences)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
