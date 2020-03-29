package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"getlines"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	start := mustAtoi(flag.Arg(1))
	end := mustAtoi(flag.Arg(2))

	lines, err := getlines.ReadRange(filename, start, end)
	if err != nil {
		panic(err)
	}
	for i, line := range lines {
		row := fillSpace(strconv.Itoa(i+start), flag.Arg(2))
		fmt.Printf("%s: %s\n", row, line)
	}
}

func fillSpace(target, to string) string {
	if len(to) < len(target) {
		return ""
	}
	count := len(to) - len(target)
	space := strings.Repeat(" ", count)
	return target + space
}

func mustAtoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
