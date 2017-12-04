package main

import (
	"fmt"
	"os"
	"strconv"
)

func diff(from, to int) int {
	if from > to {
		return from - to
	}
	return to - from
}
func shortest(from int, m []int) int {
	min := from + m[len(m)-1]
	for _, num := range m {
		if diff(from, num) < min {
			min = diff(from, num)
		}
	}
	return min
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	sum := 0
	var circle int
	for circle = 1; (circle*2-1)*(circle*2-1) < num; circle++ {
		// Hitta vilken cirkel vi är i, dvs steg från mitten rakt ut åt något håll
	}
	sum += circle - 1
	// Hitta största talet i aktuell cirkel
	maxInCircle := (circle*2 - 1) * (circle*2 - 1)
	// Hitta största talet i cirkeln innanför
	maxInInnerCircle := ((circle-1)*2 - 1) * ((circle-1)*2 - 1)
	// Antal steg från högsta numret i cirkeln till närmsta mittenposition i cirkeln (+)
	stepToMiddle := (maxInCircle - maxInInnerCircle) / 8
	// Jämför mot alla mittpositioner för att se vilken som är närmast
	sum += shortest(num, []int{maxInCircle - 7*stepToMiddle, maxInCircle - 5*stepToMiddle, maxInCircle - 3*stepToMiddle, maxInCircle - 1*stepToMiddle})
	fmt.Println(sum)
}
