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
	for scanner.Scan() {
		word := scanner.Text()
		drawsPerGame := strings.Split(strings.Split(word, ":")[1], ";")
		minCubes := map[string]int{}
		for _, val := range drawsPerGame {
			for _, ball := range strings.Split(val, ",") {
				res := strings.Split(strings.Trim(ball, " "), " ")
				num, err := strconv.Atoi(res[0])
				if err != nil {
					fmt.Println("error while coverting ball num")
				}
				if minCubes[res[1]] < num {
					minCubes[res[1]] = num
				}
			}
		}
		curVal := 1
		for _, val := range minCubes {
			curVal *= val
		}
		if curVal != 1 {
			sum += curVal
		}
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
}
