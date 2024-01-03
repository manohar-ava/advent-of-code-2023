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
		r := 0
		w := ""
		start := 0
		for r < len(mat[i]) {
			isNum := isNumber(mat[i][r])
			if isNum {
				if w == "" {
					start = r
				}
				w += mat[i][r]
			} else if w != "" {
				res := returnValIfAdj(mat, i, start, r-1)
				if res {
					val, err := strconv.Atoi(w)
					if err != nil {
						fmt.Println("error while converting str to num")
					}
					sum += val
				}
				w = ""
			}
			if r == len(mat[i])-1 && isNum && w != "" {
				res := returnValIfAdj(mat, i, start, r-1)
				if res {
					val, err := strconv.Atoi(w)
					if err != nil {
						fmt.Println("error while converting str to num")
					}
					sum += val
				}
				w = ""
			}
			r++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
	fmt.Println(sum)
}
func returnValIfAdj(mat [][]string, i int, start int, end int) bool {
	cur := checkCurArr(mat[i], start, end)
	if cur {
		return true
	}
	up := false
	down := false
	if start > 0 {
		start--
	}
	if end < len(mat[i])-1 {
		end++
	}
	if i > 1 {
		up = isSymbol(mat[i-1][start : end+1])
	}
	if i < len(mat)-1 {
		down = isSymbol(mat[i+1][start : end+1])
	}
	return up || down
}
func checkCurArr(a []string, start int, end int) bool {
	left := false
	right := false
	if start > 0 {
		start--
	}
	if end < len(a)-1 {
		end++
	}
	if !isNumber(a[end]) && !isDot(a[end]) {
		right = true
	}
	if !isNumber(a[start]) && !isDot(a[start]) {
		left = true
	}
	return left || right
}

func isSymbol(a []string) bool {
	res := false
	for _, val := range a {
		if !isNumber(val) && !isDot(val) {
			res = true
		}
	}
	return res
}
