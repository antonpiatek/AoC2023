// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

// The Elf would first like to know which games would have been possible if the bag contained
// only 12 red cubes, 13 green cubes, and 14 blue cubes?

// In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration.
// However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once;
// similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once.
// If you add up the IDs of the games that would have been possible, you get 8.

//what is the fewest number of cubes of each color that could have been in the bag to make the game possible?
//
//In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
// Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
// Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
// Game 4 required at least 14 red, 3 green, and 15 blue cubes.
// Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.
// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
// The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

// For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
package main

import (
	"reflect"
	"testing"
)

func TestProcessLines(t *testing.T) {
	gameData := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	want := 8
	res := ProcessLines(gameData)
	if res != want {
		t.Errorf("result doesn't match: expected %v, got %v", want, res)
	}
}

// The power of the minimum set of cubes in game 1 is 48.
// In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

func TestTotalPower(t *testing.T) {
	gameData := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	want := 2286
	res := TotalPower(gameData)
	if res != want {
		t.Errorf("result doesn't match: expected %v, got %v", want, res)
	}
}

func TestGamePower(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  48,
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  12,
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:  1560,
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:  630,
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:  36,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			data := ParseLine(test.input)
			actual := GamePower(data)
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf("GameMin doesn't match: expected %v, got %v", test.want, actual)
			}
		})
	}
}

func TestGameMin(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  MinData
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  MinData{red: 4, green: 2, blue: 6},
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  MinData{red: 1, green: 3, blue: 4},
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:  MinData{red: 20, green: 13, blue: 6},
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:  MinData{red: 14, green: 3, blue: 15},
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:  MinData{red: 6, green: 3, blue: 2},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			data := ParseLine(test.input)
			actual := GameMin(data)
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf("GameMin doesn't match: expected %v, got %v", test.want, actual)
			}
		})
	}
}

func TestGameValid(t *testing.T) {
	tests := []struct {
		name  string
		input GameResult
		want  bool
	}{
		{
			name: "game1",
			input: GameResult{name: 1,
				data: []GameData{
					{red: 4, blue: 3},
					{red: 1, green: 2, blue: 6},
					{green: 2},
				}},
			want: true,
		},
		{
			name: "game3",
			input: GameResult{name: 1,
				data: []GameData{ //8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
					{red: 20, blue: 6, green: 8},
					{red: 4, green: 13, blue: 5},
					{green: 5, red: 1},
				}},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GameValid(test.input)
			if actual != test.want {
				t.Errorf("GameValid doesn't match: expected %v, got %v", test.want, test.input)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		input string
		want  GameResult
	}{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: GameResult{name: 1,
				data: []GameData{
					{red: 4, blue: 3},
					{red: 1, green: 2, blue: 6},
					{green: 2},
				}}},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: GameResult{name: 2,
				data: []GameData{
					{blue: 1, green: 2},
					{red: 1, green: 3, blue: 4},
					{green: 1, blue: 1},
				}}},
		{input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: GameResult{name: 4,
				data: []GameData{
					{red: 3, blue: 6, green: 1},
					{red: 6, green: 3},
					{green: 3, blue: 15, red: 14},
				}}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := ParseLine(test.input)
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf("line doesn't match: expected %v, got %v", test.want, test.input)
			}
		})
	}
}
