package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filePath string) ([]int, []int) {
	f, err := os.Open(filePath)
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
	return column1, column2
}
