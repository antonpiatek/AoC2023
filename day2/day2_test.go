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
