// https://www.hackerrank.com/challenges/array-left-rotation/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rotate(numbers []int, rotation int) []int {
	// initialise the output array
	var output = make([]int, len(numbers))

	// copy array from index "rotation" until the end to new array
	var idx = 0
	for i := rotation; i < len(numbers); i++ {
		output[idx] = numbers[i]
		idx++
	}

	// copy array from index 0 until "rotation" to new array
	for i := 0; i < rotation; i++ {
		output[idx] = numbers[i]
		idx++
	}

	return output
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var strNumOfNumbers, _ = reader.ReadString(' ')
	var numOfNumbers, _ = strconv.Atoi(strings.TrimSpace(strNumOfNumbers))

	var strRotation, _ = reader.ReadString('\n')
	var rotation, _ = strconv.Atoi(strings.TrimSpace(strRotation))

	var strNumbers, _ = reader.ReadString('\n')
	var strNumbersArr = strings.Split(strings.TrimSpace(strNumbers), " ")
	var numbers = make([]int, numOfNumbers)

	for i := 0; i < numOfNumbers; i++ {
		numbers[i], _ = strconv.Atoi(strings.TrimSpace(strNumbersArr[i]))
	}

	var result = rotate(numbers, rotation)

	for i := 0; i < numOfNumbers; i++ {
		fmt.Printf("%d ", result[i])
	}
}
