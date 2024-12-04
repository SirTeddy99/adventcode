package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file for reading
	file, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var yeetStr []string
	var yootStr []string
	var yeetInt []int
	var yootInt []int
	// Loop through each line in the file
	for scanner.Scan() {
		// Get the current line from the scanner
		line := scanner.Text()

		parts := strings.Split(line, "   ")

		fmt.Println("splitt line:", parts)

		for i, part := range parts {
			if i == 0 {
				fmt.Println("parts:", part)
				yeetStr = append(yeetStr, part)
			} else {
				fmt.Println("parts:", part)
				yootStr = append(yootStr, part)
			}
		}
	}
	fmt.Println("list yeetStr", yeetStr)
	fmt.Println("list yootStr", yootStr)

	for _, str := range yeetStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting yeet-string to int:", err)
			continue
		}
		yeetInt = append(yeetInt, num)
	}

	for _, str := range yootStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting yoot-string to int:", err)
			continue
		}
		yootInt = append(yootInt, num)
	}

	sort.Ints(yootInt)
	sort.Ints(yeetInt)
	fmt.Println("list yeetInt sorted", yeetInt)
	fmt.Println("list yootInt sorted", yootInt)

	if len(yeetInt) != len(yootInt) {
		log.Fatal("Error: yeetInt and yootInt not same length")
	}

	var sum int
	for i, yoo := range yootInt {
		for j, yee := range yeetInt {
			if i == j {
				if yee > yoo {
					sum += yee - yoo
				} else {
					sum += yoo - yee
				}
			}
		}
	}
	fmt.Println("finnal diff", sum)

	// or use this method:
	//
	// var sumNew int
	// for i := range yeetInt {
	// 	diff := abs(yeetInt[i] - yootInt[i])
	// 	sumNew += diff
	// }
	// fmt.Println("finnal diff", sum)
	// // Check for errors that may have occurred during the scan
	// if err := scanner.Err(); err != nil {
	// 	log.Fatalf("Error reading file: %v", err)
	// }

	// part 2

	var sum2 int
	for _, yee := range yeetInt {
		newYorkTimes := 0
		for _, yoo := range yootInt {
			if yee == yoo {
				newYorkTimes++
			}
		}
		sum2 += yee * newYorkTimes
	}
	fmt.Println("finnal diff", sum2)
}

// Absolute difference function
// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
