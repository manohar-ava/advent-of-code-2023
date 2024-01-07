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
	dic := getTimeAndDistance(lines)

	for time, dis := range dic {

		val := getNoOfPossibleWins(time, dis)
		prod *= val
	}
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

func getTimeAndDistance(arr []string) map[int]int {
	res := map[int]int{}
	dis := getTrimmedIntSlice(strings.Split(arr[1], ":")[1])
	time := getTrimmedIntSlice(strings.Split(arr[0], ":")[1])
	for i, val := range time {
		res[val] = dis[i]
	}
	return res
}

func getTrimmedIntSlice(s string) []int {
	sl := []int{}
	for _, val := range strings.Split(strings.TrimSpace(s), " ") {
		if strings.TrimSpace(val) != "" {
			v, _ := strconv.Atoi(val)
			sl = append(sl, v)
		}
	}
	return sl
}
