package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type measurement struct {
	val1, val2, val3 int
}

func (m measurement) sum() int {
	return m.val1 + m.val2 + m.val3
}

func (m measurement) new(i int) measurement {
	return measurement{m.val2, m.val3, i}
}
func (m measurement) isFull() bool {
	return m.val1 != 0 && m.val2 != 0 && m.val3 != 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var count int
	var largerSumCount int
	if !scanner.Scan() {
		panic("empty input")
	}
	line := scanner.Text()
	last, err := strconv.Atoi(line)
	curM := measurement{val3: last}
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
		newM := curM.new(cur)
		if !curM.isFull() || !newM.isFull() {
			curM = newM
			continue
		}
		if newM.sum() > curM.sum() {
			largerSumCount++
		}
		curM = newM
	}
	fmt.Printf("How many measurements are larger than the previous measurement? %d\n", count)
	fmt.Printf("How many sums are larger than the previous sum? %d\n", largerSumCount)

}
