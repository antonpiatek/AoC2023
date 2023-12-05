package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func ProcessLine(name string) int {
	split := strings.Split(name, "")

	first := ""
	last := ""
	for _, element := range split {
		if _, err := strconv.Atoi(element); err == nil {
			if first == "" {
				first = element
			}
			last = element
		}
	}

	result, _ := strconv.Atoi(first + last)
	return result
}
