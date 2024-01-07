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
	prod := 1
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	time, dis := getTimeAndDistance(lines)
	prod *= getNoOfPossibleWins(time, dis)
	fmt.Println(prod)
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
}

func getNoOfPossibleWins(time int, dis int) int {
	sum := 0
	mid := int(math.Floor(float64(time / 2)))
	for {
		if isValidTime(mid, time, dis) {
			break
		}
		innerMid := int(math.Floor(float64(time / 2)))
		mid += innerMid
	}
	l := mid - 1
	r := mid + 1
	lbreak := false
	rbreak := false
	for {
		lbreak = !isValidTime(l, time, dis)
		rbreak = !isValidTime(r, time, dis)
		if lbreak && rbreak {
			sum = r - l - 2
			sum += 1
			break
		}
		if !lbreak {
			l--
		}
		if !rbreak {
			r++
		}
	}
	return sum
}

func isValidTime(i int, time int, dis int) bool {
	return i*(time-i) > dis
}

func getTimeAndDistance(arr []string) (int, int) {
	dis, _ := strconv.Atoi(trimSplitAndAppend(strings.Split(arr[1], ":")[1]))
	time, _ := strconv.Atoi(trimSplitAndAppend(strings.Split(arr[0], ":")[1]))
	return time, dis
}

func trimSplitAndAppend(s string) string {
	sl := []string{}
	for _, val := range strings.Split(strings.TrimSpace(s), " ") {
		if strings.TrimSpace(val) != "" {
			sl = append(sl, val)
		}
	}
	return strings.Join(sl, "")
}
