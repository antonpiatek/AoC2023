package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = strings.TrimSuffix(line, "\n")

		// fmt.Println(line)
		total += ProcessLine(line)

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	f.Close()

	fmt.Println(total)

	return nil
}

func ProcessData(data []string) int {
	total := 0
	for _, row := range data {
		rowVal := ProcessLine(row)
		total += rowVal
	}
	return total
}

// Can't do lookaheads in go apparently
// Could almost do with one regex, but doesn't work for only one value in a string
// - `.*?(\d|one|two|three|four|five|six|seven|eight|nine).*(\d|one|two|three|four|five|six|seven|eight|nine)`
var r1 = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
var r2 = regexp.MustCompile(`.*(\d|one|two|three|four|five|six|seven|eight|nine)`)

func ProcessLine(line string) int {
	x1 := r1.FindAllStringSubmatch(line, -1)
	x2 := r2.FindAllStringSubmatch(line, -1)
	first := toNum(x1[0][1])
	last := toNum(x2[len(x2)-1][1])
	// fmt.Print(x)

	result, _ := strconv.Atoi(first + last)
	return result
}

func toNum(s string) string {
	_, err := strconv.Atoi(s)
	if err == nil {
		return s
	}

	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		panic("unimplemented")
	}
}
