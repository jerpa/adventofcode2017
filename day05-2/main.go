package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sum := 0
	var list []int
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		list = append(list, value)
	}
	pos := 0
	for {
		step := list[pos]
		if step >= 3 {
			list[pos] = list[pos] - 1
		} else {
			list[pos] = list[pos] + 1
		}
		pos += step
		sum++
		if pos < 0 || pos >= len(list) {
			break
		}
	}
	fmt.Println(sum)
}
