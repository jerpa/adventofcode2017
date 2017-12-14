package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	tree := make(map[string]string)
	scanner := bufio.NewScanner(file)
	var start string
	for scanner.Scan() {
		value := scanner.Text()
		if strings.Index(value, "->") >= 0 {
			key := value[0:strings.Index(value, " (")]
			start = key
			children := value[strings.Index(value, "->")+3:]
			for _, val := range strings.Split(children, ", ") {
				tree[val] = key
			}
		}
		//		list = append(list, value)
	}
	for {
		if _, ok := tree[start]; !ok {
			fmt.Println(start)
			break
		} else {
			start = tree[start]
		}
	}
}
