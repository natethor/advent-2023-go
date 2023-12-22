package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	CardId          int
	WinningNumbers  []int
	PossibleNumbers []int
}

func convertToIntegers(stringNumbers []string) ([]int, error) {
	var intNumbers []int

	for _, str := range stringNumbers {
		str = strings.TrimSpace(str)
		if str != "" {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			intNumbers = append(intNumbers, num)
		}
	}
	return intNumbers, nil
}

func processInput() (cards []Card) {
	if file, err := os.Open("./input/day4.txt"); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		counter := 1

		for scanner.Scan() {
			str := scanner.Text()

			str = strings.Split(str, ": ")[1]
			str = strings.TrimSpace(str)

			splitStr := strings.Split(str, "|")

			winningNumStrings := strings.Split(splitStr[0], " ")
			possibleNumStrings := strings.Split(splitStr[1], " ")
			winningNumbers, _ := convertToIntegers(winningNumStrings)
			possibleNumbers, _ := convertToIntegers(possibleNumStrings)

			currentCard := Card{
				CardId:          counter,
				WinningNumbers:  winningNumbers,
				PossibleNumbers: possibleNumbers}

			cards = append(cards, currentCard)
			counter++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
	return cards
}

func (c Card) getWinningNumbers() []int {
	var winners []int
	for _, winningNumber := range c.WinningNumbers {
		for _, possibleNumber := range c.PossibleNumbers {
			if winningNumber == possibleNumber {
				winners = append(winners, winningNumber)
			}
		}
	}
	return winners
}

func cardScore(matches []int) int {
	var totalScore int = 0
	if len(matches) > 0 {
		score := 1
		for i := 1; i < len(matches); i++ {
			score *= 2
		}
		totalScore = score
	}
	return totalScore
}

func Day4Part1() int {
	cards := processInput()
	var totalScore int = 0
	for _, card := range cards {
		fmt.Println("Card: ", card.CardId)
		winners := card.getWinningNumbers()
		fmt.Println("Winning Numbers: ", winners)
		cardScore := cardScore(winners)
		fmt.Println("Score: ", cardScore)
		totalScore += cardScore
	}
	return totalScore
}
