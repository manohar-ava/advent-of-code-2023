package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isDot(s string) bool {
	return s == "."
}
func isNumber(s string) bool {
	runeSilice := []rune(s)
	return unicode.IsDigit(runeSilice[0])
}
func isStar(s string) bool {
	return s == "*"
}
func getProduct(x int, y int, mat [][]string) int {
	dirs := [][]int{
		{0, -1},
		{-1, -1}, {-1, 0}, {-1, 1},
		{1, 1}, {1, 0}, {1, -1},
		{0, 1},
	}
	nums := map[int]int{}
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		cur := mat[dx+x][dy+y]
		if isNumber(cur) {
			rem := getNumAtDirection(dx+x, dy+y, mat)
			nums[rem] = true
		}
	}
	fmt.Println(nums)
	if len(nums) == 2 {
		prod := 1
		for v, _ := range nums {
			prod *= v
		}
		return prod
	}
	return 0
}
func getNumAtDirection(x int, y int, mat [][]string) int {
	w := mat[x][y]
	rr := ""
	ll := ""
	l := y - 1
	r := y + 1
	for r < len(mat[x]) && isNumber(mat[x][r]) {
		rr += mat[x][r]
		r++
	}
	for l >= 0 && isNumber(mat[x][l]) {
		ll = mat[x][l] + ll
		l--
	}
	res, err := strconv.Atoi(ll + w + rr)
	if err != nil {
		println("error while converting single num")
	}
	return res
}
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error while reading file!!!")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	mat := [][]string{}
	for scanner.Scan() {
		word := scanner.Text()
		mat = append(mat, strings.Split(word, ""))
	}
	for i := 0; i < len(mat); i++ {
		row := mat[i]
		for j := 0; j < len(row); j++ {
			if isStar(row[j]) {
				nums := getProduct(i, j, mat)

				sum += nums
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
	fmt.Println(sum)
}
