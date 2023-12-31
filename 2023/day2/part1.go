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
		gameNo, err := strconv.Atoi(strings.Split(strings.Split(word, ":")[0], " ")[1])
		drawsPerGame := strings.Split(strings.Split(word, ":")[1], ";")
		shouldAdd := true
		cubesPerGame := map[string]int{}
		for _, val := range drawsPerGame {
			for _, ball := range strings.Split(val, ",") {
				res := strings.Split(strings.Trim(ball, " "), " ")
				num, err := strconv.Atoi(res[0])
				if err != nil {
					fmt.Println("error while coverting ball num")
				}
				cubesPerGame[res[1]] = num
			}
			if cubesPerGame["green"] > 13 || cubesPerGame["red"] > 12 || cubesPerGame["blue"] > 14 {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {
			sum += gameNo
		}
		if err != nil {
			println("error while coverting")
		}
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
}
