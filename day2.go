package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
which games would have been possible if the bag contained only
12 red cubes, 13 green cubes, and 14 blue cubes?
*/
type CubeCount struct {
	Color string
	Cubes int
}

type Set struct {
	SetNumber int
	Cubes     []CubeCount
}

type Game struct {
	GameNumber int
	Sets       []Set
}

const (
	Green      = "green"
	Red        = "red"
	Blue       = "blue"
	RedCount   = 12
	GreenCount = 13
	BlueCount  = 14
)

func extractAfterColon(input string) string {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func extractBeforeColon(input string) string {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}

func parseGameSets(input string) []Set {
	sets := strings.Split(input, ";")
	var all_game_sets []Set

	for setNumber, set := range sets {
		var current_set Set
		current_set.SetNumber = setNumber + 1

		cubeCounts := strings.Split(set, ",")
		for _, cubeCount := range cubeCounts {
			var current_cube CubeCount

			trimmedString := strings.TrimSpace(cubeCount)
			number_color := strings.Split(trimmedString, " ")

			current_cube.Color = number_color[1]
			current_cube.Cubes, _ = strconv.Atoi(number_color[0])
			current_set.Cubes = append(current_set.Cubes, current_cube)
		}
		all_game_sets = append(all_game_sets, current_set)
	}
	return all_game_sets
}

func validateSet(set Set) bool {
	for _, cubeCount := range set.Cubes {
		switch cubeCount.Color {
		case Green:
			if cubeCount.Cubes > GreenCount {
				return false
			}
		case Red:
			if cubeCount.Cubes > RedCount {
				return false
			}
		case Blue:
			if cubeCount.Cubes > BlueCount {
				return false
			}
		}
	}
	return true
}

func validateGame(game Game) bool {
	for _, set := range game.Sets {
		if !validateSet(set) {
			return false
		}
	}
	return true
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func processFile() []Game {
	all_games := make([]Game, 0)

	if file, err := os.Open("./input/day2.txt"); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			str := scanner.Text()
			game_string := extractBeforeColon(str)
			current_sets := extractAfterColon(str)

			game_string_split := strings.Split(game_string, " ")[1]
			game_number, _ := strconv.Atoi(game_string_split)

			current_game_sets := parseGameSets(current_sets)
			all_games = append(all_games, Game{GameNumber: game_number, Sets: current_game_sets})
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
	return all_games
}

func Day2Part1() int {
	all_games := processFile()

	valid_games := make([]int, 0)
	for _, game := range all_games {
		current_game_valid := validateGame(game)

		if current_game_valid {
			valid_games = append(valid_games, game.GameNumber)
		}
	}

	total := sum(valid_games)
	return total
}

func Day2Part2() int {
	/*
		in each game you played, what is the fewest number of cubes of each color
		that could have been in the bag to make the game possible?
		The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
		For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
	*/
	all_games := processFile()

	type Minimums struct {
		Red   int
		Green int
		Blue  int
	}

	total := 0

	for _, game := range all_games {
		var game_minimums Minimums

		for _, set := range game.Sets {
			for _, cubeCount := range set.Cubes {
				if cubeCount.Color == Green {
					if game_minimums.Green < cubeCount.Cubes {
						game_minimums.Green = cubeCount.Cubes
					}
				} else if cubeCount.Color == Red {
					if game_minimums.Red < cubeCount.Cubes {
						game_minimums.Red = cubeCount.Cubes
					}
				} else if cubeCount.Color == Blue {
					if game_minimums.Blue < cubeCount.Cubes {
						game_minimums.Blue = cubeCount.Cubes
					}
				}
			}
		}
		game_min_power := game_minimums.Red * game_minimums.Green * game_minimums.Blue
		total += game_min_power
	}

	return total
}
