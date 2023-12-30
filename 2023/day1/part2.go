package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		wordRunes := []rune(word)
		r := len(word) - 1
		l := 0
		var first rune
		var second rune
		var fword string = ""
		var lword string = ""
		for {
			if first != 0 && second != 0 {
				break
			}
			fword = fword + string(wordRunes[l])
			lword = string(wordRunes[r]) + lword
			if first == 0 {
				isPresentf, valf := checkDict(fword)
				if isPresentf {
					first = valf
				} else if unicode.IsDigit(wordRunes[l]) {
					first = wordRunes[l]
				}
			}
			if second == 0 {
				isPresentl, vall := checkDict(lword)
				if isPresentl {
					second = vall
				} else if unicode.IsDigit(wordRunes[r]) {
					second = wordRunes[r]
				}
			}
			r--
			l++
		}
		fmt.Println(string([]rune{first, second}))
		val, err := strconv.Atoi(string([]rune{first, second}))
		if err != nil {
			fmt.Println("error while converting", val)
		}
		sum += val
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
	fmt.Println(sum)
}
func checkDict(str string) (bool, rune) {
	isPresent := false
	var val rune
	dict := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for key := range dict {
		if strings.Contains(str, key) {
			isPresent = true
			val = []rune(dict[key])[0]
		}
	}
	return isPresent, val
}
