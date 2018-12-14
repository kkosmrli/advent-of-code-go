package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//part1()
	part2()
}

func part1() {
	file, err := os.Open("day1.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("%s", err)
	}

	scanner := bufio.NewScanner(file)
	frequency := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		frequency = frequency + num
	}
	fmt.Println(frequency)
}

func part2() {
	file, err := os.Open("day1.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("%s", err)
	}

	scanner := bufio.NewScanner(file)
	m := map[int]bool{0: true}

	var changes []int
	for scanner.Scan() {
		line := scanner.Text()
		change, _ := strconv.Atoi(line)
		changes = append(changes, change)
	}

	frequency := 0
	for {
		for _, n := range changes {
			frequency += n

			if m[frequency] {
				fmt.Println(frequency)
				return
			}
			m[frequency] = true
		}
	}
}
