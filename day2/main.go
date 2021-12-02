package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type submarine struct {
	horPos int
	depth  int
	aimAt  int
}

func (s *submarine) forward(i int) {
	s.horPos += i
	s.depth += s.aimAt * i
}

func (s *submarine) aim(i int) {
	s.aimAt += i
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	s := &submarine{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "forward":
			s.forward(toInt(parts[1]))
		case "down":
			s.aim(toInt(parts[1]))
		case "up":
			s.aim(-toInt(parts[1]))
		}
	}
	fmt.Printf("What do you get if you multiply your final horizontal position by your final depth? %d\n", s.depth*s.horPos)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
