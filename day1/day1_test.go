package main

import (
	"testing"
)

func TestSummed(t *testing.T) {
	input := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	want := 142

	actual := ProcessData(input)
	if actual != want {
		t.Errorf("sums don't match: expected %v, got %v", want, input)
	}
}

func TestData(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "1abc2", want: 12},
		{input: "pqr3stu8vwx", want: 38},
		{input: "a1b2c3d4e5f", want: 15},
		{input: "treb7uchet", want: 77},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := ProcessLine(test.input)

			if actual != test.want {
				t.Errorf("sums don't match: expected %v, got %v", test.want, test.input)
			}
		})
	}
}

func TestWordyData(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "two1nine", want: 29},
		{input: "eightwothree", want: 83},
		{input: "abcone2threexyz", want: 13},
		{input: "xtwone3four", want: 24},
		{input: "4nineeightseven2", want: 42},
		{input: "zoneight234", want: 14},
		{input: "7pqrstsixteen", want: 76},
		{input: "2onenine", want: 29},
		{input: "2oneight", want: 28},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := ProcessLine(test.input)

			if actual != test.want {
				t.Errorf("sums don't match: expected %v, got %v", test.want, actual)
			}
		})
	}
}

// 281
