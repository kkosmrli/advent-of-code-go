package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("day2.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("%s", err)
	}

	var three, two int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		occ := make(map[rune]int)
		for _, r := range line {
			occ[r]++
		}
		var th, tw bool
		for _, i := range occ {

			if i == 2 && !tw {
				two++
				tw = true
			}
			if i == 3 && !th {
				three++
				th = true
			}
		}
	}
	fmt.Println(two * three)
}

func part2() {
	file, err := os.Open("day2.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("%s", err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	proto := make(map[string]int)

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			word := removeAtIndex(line, i)
			proto[word]++
		}
	}

	for k, h := range proto {
		if h > 1 {
			fmt.Println(k)
		}
	}
}

func removeAtIndex(in string, i int) string {
	out := []rune(in)
	out = append(out[:i], out[i+1:]...)
	return string(out)
}
