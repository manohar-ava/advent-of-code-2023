package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error while reading file!!!")
	}
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		cardNums := strings.Split(parts[1], "|")
		value := computeVal(cardNums[0], cardNums[1])
		sum += value
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
	fmt.Println(sum)
}

func computeVal(winningNums string, nums string) int {
	vals := map[string]int{}
	sum := 0
	for _, val := range getTrimmedSlice(winningNums) {
		vals[val] = 1
	}

	for _, val := range getTrimmedSlice(nums) {
		if vals[val] == 1 {
			if sum == 0 {
				sum++
			} else {
				sum += sum
			}
		}
	}
	return sum
}

func getTrimmedSlice(s string) []string {
	sl := []string{}
	for _, val := range strings.Split(strings.TrimSpace(s), " ") {
		if val != "" {
			sl = append(sl, val)
		}
	}
	return sl
}
