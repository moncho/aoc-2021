package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var count int
	if !scanner.Scan() {
		panic("empty input")
	}
	line := scanner.Text()
	last, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	for scanner.Scan() {
		line = scanner.Text()
		cur, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if cur > last {
			count++
		}
		last = cur
	}
	fmt.Printf("How many measurements are larger than the previous measurement? %d\n", count)
}
