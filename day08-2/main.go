package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var values map[string]int
var top int

func math(register, action string, num int) {
	if action == "inc" {
		values[register] += num
	} else {
		values[register] -= num
	}
	if values[register] > top {
		top = values[register]
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	values = make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()

		parts := strings.Split(value, " ")
		cmpnum, _ := strconv.Atoi(parts[6])
		num, _ := strconv.Atoi(parts[2])
		if parts[5] == ">" && values[parts[4]] > cmpnum {
			math(parts[0], parts[1], num)
		} else if parts[5] == "<" && values[parts[4]] < cmpnum {
			math(parts[0], parts[1], num)
		} else if parts[5] == ">=" && values[parts[4]] >= cmpnum {
			math(parts[0], parts[1], num)
		} else if parts[5] == "<=" && values[parts[4]] <= cmpnum {
			math(parts[0], parts[1], num)
		} else if parts[5] == "==" && values[parts[4]] == cmpnum {
			math(parts[0], parts[1], num)
		} else if parts[5] == "!=" && values[parts[4]] != cmpnum {
			math(parts[0], parts[1], num)
		}
	}
	/*
		max := 0
		for _, val := range values {
			if val > max {
				max = val
			}
		}*/
	fmt.Println(top)
}
