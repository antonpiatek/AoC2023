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
	lines, err := ReadData()
	if err != nil {
		return err
	}

	res := ProcessLines(lines)
	fmt.Println("part a:", res)

	res2 := TotalPower(lines)
	fmt.Println("part b:", res2)
	return nil
}

type GameResult struct {
	name int
	data []GameData
}

type GameData struct {
	red   int
	blue  int
	green int
}

type MinData = GameData

func ProcessLines(lines []string) int {
	total := 0
	for _, line := range lines {
		gameData := ParseLine(line)
		if GameValid(gameData) {
			total += gameData.name
		}
	}
	return total
}

func TotalPower(lines []string) int {
	total := 0
	for _, line := range lines {
		gameData := ParseLine(line)
		pow := GamePower(gameData)
		total += pow
	}
	return total
}

func GamePower(data GameResult) int {
	min := GameMin(data)
	return min.blue * min.green * min.red
}

func GameMin(data GameResult) MinData {
	res := MinData{}
	for _, i := range data.data {
		if i.red > res.red {
			res.red = i.red
		}
		if i.green > res.green {
			res.green = i.green
		}
		if i.blue > res.blue {
			res.blue = i.blue
		}
	}

	return res
}

// The Elf would first like to know which games would have been possible if the bag contained
// only 12 red cubes, 13 green cubes, and 14 blue cubes?
const redLimit = 12
const greenLimit = 13
const blueLimit = 14

func GameValid(gameresult GameResult) bool {
	for _, game := range gameresult.data {
		if game.red > redLimit || game.blue > blueLimit || game.green > greenLimit {
			return false
		}
	}
	return true
}

var r1 = regexp.MustCompile(`Game (\d+):\s*(.+)`)
var r2 = regexp.MustCompile(`(\d+)\s*(\w+)`)

func ParseLine(line string) GameResult {

	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	x1 := r1.FindStringSubmatch(line)

	gameNum, _ := strconv.Atoi(x1[1])
	res := GameResult{name: gameNum}
	for _, game := range strings.Split(x1[2], ";") {
		var gdata = extractGameData(game)
		// fmt.Printf("%+v\n", gdata)
		res.data = append(res.data, gdata)
	}

	return res
}

func extractGameData(game string) GameData {
	result := GameData{}
	for _, colourData := range strings.Split(game, ",") {
		x2 := r2.FindStringSubmatch(colourData)
		count, _ := strconv.Atoi(x2[1])
		colour := x2[2]
		switch colour {
		case "red":
			result.red = count
		case "green":
			result.green = count
		case "blue":
			result.blue = count
		default:
			panic("unimplemented")
		}

	}

	return result
}

func ReadData() ([]string, error) {
	f, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	f.Close()

	return lines, nil
}
