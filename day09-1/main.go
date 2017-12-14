package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		tot := 0
		val := 0
		ignore := false
		for pos := 0; pos < len(value); pos++ {
			chr := value[pos : pos+1]
			if chr == "!" {
				pos++
			} else if chr == "{" && ignore == false {
				val++
			} else if chr == "}" && ignore == false {
				tot += val
				val--
			} else if chr == "<" {
				ignore = true
			} else if chr == ">" {
				ignore = false
			}
		}

		fmt.Println(tot)
	}
}
