package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	binNoSize = 12
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines int
	var countOnes [binNoSize]int
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '1' {
				countOnes[binNoSize-1-i]++
			}
		}
		lines++
	}
	var gamma, epsilon int
	half := lines / 2

	for i := binNoSize - 1; i >= 0; i-- {
		if countOnes[i] > half {
			gamma += 1 << i
		} else {
			epsilon += 1 << i
		}
	}
	fmt.Printf("What is the power consumption of the submarine? %d\n", gamma*epsilon)
}
