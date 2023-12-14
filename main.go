package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(label string) string {
	var str string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		str, _ = r.ReadString('\n')
		if str != "" {
			break
		}
	}
	return strings.TrimSpace(str)
}

func main() {
	whichDay := Prompt("Which day? (1, 2, or 3): ")
	fmt.Printf("Running day %s\n", whichDay)
	switch whichDay {
	case "1":
		fmt.Println("Day1Part1: ", Day1Part1())
		fmt.Println("Day1Part2: ", Day1Part2())
	case "2":
		fmt.Println("Day2Part1: ", Day2Part1())
		fmt.Println("Day2Part2: ", Day2Part2())
	case "3":
		fmt.Println("Day3Part1: ", Day3Part1())
		fmt.Println("Day3Part2: ", Day3Part2())
	default:
		fmt.Println("Invalid day")
	}
}
