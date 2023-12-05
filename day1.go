package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ints_from_str(s string) map[int]string {
	numbers := make(map[int]string)
	str_slice := strings.Split(s, "")

	for i, char := range str_slice {
		if num, err := strconv.Atoi(char); err == nil {
			numbers[i] = strconv.Itoa(num)
		}
	}
	return numbers
}

func letters_from_str(s string) map[int]string {
	digit_map := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	result_map := make(map[int]string)

	for digit, letters := range digit_map {
		start := 0
		for {
			index := strings.Index(s[start:], letters)
			if index == -1 {
				break // No more occurrences found
			}
			indices := start + index
			result_map[indices] = fmt.Sprintf("%d", digit)
			start += index + 1
		}
	}
	return result_map
}

func Day1Part1() int {
	/*
		On each line, the calibration value can be found by combining
		the first digit and the last digit (in that order) to form a single two-digit number
		ignore the letters.  ignore the other numbers in the middle.
		then add all the numbers together to get the sum of all calibration values.
	*/

	var total int = 0

	if file, err := os.Open("./input/day1.txt"); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			str := scanner.Text()

			number_map := ints_from_str(str)

			var numbers []string
			for _, value := range number_map {
				numbers = append(numbers, value)
			}

			num_strings := make([]string, len(numbers))
			copy(num_strings, numbers)

			first_last := [2]string{num_strings[0], num_strings[len(numbers)-1]}
			combined := strings.Join(first_last[:], "")
			if line_digit, err := strconv.Atoi(combined); err != nil {
				fmt.Print(err)
			} else {
				total += line_digit
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
	return total
}

func Day1Part2() int {
	/*
		It looks like some of the digits are actually spelled out with letters:
		one, two, three, four, five, six, seven, eight, and nine
		also count as valid "digits"
	*/
	var total int = 0

	if file, err := os.Open("./input/day1.txt"); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			str := scanner.Text()

			number_map := ints_from_str(str)
			letter_map := letters_from_str(str)

			combinedMap := make(map[int]string)

			for key, value := range number_map {
				combinedMap[key] = value
			}

			for key, value := range letter_map {
				if _, exists := combinedMap[key]; !exists {
					combinedMap[key] = value
				}
			}

			// Find the minimum and maximum keys
			var minKey, maxKey int
			minMaxSet := false
			for key := range combinedMap {
				if !minMaxSet {
					minKey, maxKey = key, key
					minMaxSet = true
				} else {
					minKey = min(minKey, key)
					maxKey = max(maxKey, key)
				}
			}

			// Retrieve values for the minimum and maximum keys
			minValue := combinedMap[minKey]
			maxValue := combinedMap[maxKey]

			combined := minValue + maxValue

			if line_digit, err := strconv.Atoi(combined); err != nil {
				fmt.Print(err)
			} else {
				total += line_digit
			}
		}
	}
	return total
}
