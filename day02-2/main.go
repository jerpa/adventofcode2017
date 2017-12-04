package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func rowCheckSum(row string) int {
	stringNumbers := strings.Split(row, "\t")
	var numbers []int
	for _, num := range stringNumbers {
		iNum, _ := strconv.Atoi(num)
		numbers = append(numbers, iNum)
	}
	sort.Ints(numbers)
	for pos := 0; pos < len(numbers); pos++ {
		for subPos := len(numbers) - 1; subPos > pos; subPos-- {
			if numbers[subPos]%numbers[pos] == 0 {
				return numbers[subPos] / numbers[pos]
			}
		}
	}
	return 0
}

func main() {
	sum := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		if len(input) == 0 {
			break
		}
		sum += rowCheckSum(input)
	}
	fmt.Println(sum)
}
