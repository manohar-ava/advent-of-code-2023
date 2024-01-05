package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error while reading file!!!")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	queue := make([]string, len(lines))
	copy(queue, lines)
	for len(queue) > 0 {
		line := strings.Split(queue[0], ":")
		cardNo, err := strconv.Atoi(strings.TrimSpace(getTrimmedSlice(line[0])[1]))
		if err != nil {
			fmt.Println("error while getting cardNo", cardNo)
		}
		sum++
		nums := strings.Split(line[1], "|")
		queue = queue[1:]
		length := computeVal(nums[0], nums[1])
		if length > 0 {
			queue = append(queue, lines[cardNo:cardNo+length]...)
		}
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
			sum++
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
