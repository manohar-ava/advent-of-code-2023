package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// 	test := `1abcj
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet`
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error while reading file!!!")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)
		wordRunes := []rune(word)
		r := len(word) - 1
		l := 0
		var first rune
		var second rune
		for {
			if unicode.IsDigit(wordRunes[l]) && first == 0 {
				first = wordRunes[l]
			}
			if unicode.IsDigit(wordRunes[r]) && second == 0 {
				second = wordRunes[r]
			}
			if first != 0 && second != 0 {
				break
			}
			r--
			l++
		}
		val, err := strconv.Atoi(string([]rune{first, second}))
		if err != nil {
			fmt.Println("error while converting")
		}
		sum += val
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading lines")
	}
	fmt.Println(sum)
}
