package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rowCheckSum(row string) int {
	numbers := strings.Split(row, "\t")
	var min, max int
	for _, num := range numbers {
		cNum, _ := strconv.Atoi(num)
		if min == 0 || cNum < min {
			min = cNum
		}
		if max == 0 || cNum > max {
			max = cNum
		}

	}
	return max - min
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
