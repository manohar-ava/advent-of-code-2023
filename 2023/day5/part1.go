package main

import (
	"bufio"
	"fmt"
	"math"
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
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	seeds := lines[0]
	lines = lines[2:]
	maps := getMaps(lines)
	seedArr := getTrimmedSlice(strings.Split(seeds, ":")[1])
	valArr := []int{}
	for _, seed := range seedArr {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("error while casting seed")
		}
		valArr = append(valArr, getSeedValue(seedInt, maps, 0))
	}
	fmt.Println(findSmallesIntArr(valArr))
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
}
func findSmallesIntArr(nums []int) int {
	smallest := nums[0]
	for _, v := range nums {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func getSeedValue(seed int, maps [][]string, i int) int {
	if i < 0 || i >= len(maps) {
		return seed
	}
	for _, val := range maps[i] {
		valArr := strings.Split(val, " ")
		des, _ := strconv.Atoi(valArr[0])
		src, _ := strconv.Atoi(valArr[1])
		r, _ := strconv.Atoi(valArr[2])
		if seed >= src && seed <= (src+r) {
			diff := math.Abs(float64(seed - src))
			v := float64(des) + diff
			seed = int(v)
			break
		}
	}
	return getSeedValue(seed, maps, i+1)
}

func getMaps(lines []string) [][]string {
	maps := [][]string{}
	tempArr := []string{}
	for i, ln := range lines {
		isKey := strings.Contains(ln, ":")
		isLineEmpty := strings.TrimSpace(ln) == ""
		if isLineEmpty {
			continue
		}
		if !isKey {
			tempArr = append(tempArr, ln)
		} else {
			if len(tempArr) > 0 {
				maps = append(maps, tempArr)
			}
			tempArr = []string{}
		}
		if i == len(lines)-1 && len(tempArr) > 0 {
			maps = append(maps, tempArr)
		}
	}
	return maps
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
