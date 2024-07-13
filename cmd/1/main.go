// Your program is to use the brute-force approach in order to find the Answer
// to Life, the Universe, and Everything. More precisely... rewrite small numbers
// from input to output. Stop processing input after reading in the number 42.
// All numbers at input are integers of one or two digits.

// input: 1, 2, 88, 42, 99
// output: 1, 2, 88

package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

var BREAK_NUMBER = 42

func main() {
	var s scanner.Scanner
	s.Init(os.Stdin)
	var numbers []int
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		number, err := strconv.Atoi(s.TokenText())
		if err != nil {
			fmt.Println("Only numbers")
			continue
		}
		numbers = append(numbers, number)
		if len(numbers) > 2 && numbers[len(numbers)-2] == BREAK_NUMBER {
			break
		}

	}
	for _, n := range numbers {
		if n == BREAK_NUMBER {
			break
		}
		fmt.Println(n)
	}
}
