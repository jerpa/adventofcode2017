package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := os.Args[1]
	//	data = data + data[0:1]
	sum := 0
	for pos := 0; pos < len(data); pos++ {
		cPos := (pos + (len(data) / 2)) % len(data)
		if data[pos:pos+1] == data[cPos:cPos+1] {
			val, _ := strconv.Atoi(data[pos : pos+1])
			sum += val
		}
	}
	fmt.Println(sum)
}
