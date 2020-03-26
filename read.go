package getlines

import (
	"bufio"
	"errors"
	"os"
)

func ReadRange(filename string, start, end int) ([]string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	_, err = skipLine(scanner, start)
	if err != nil {
		return []string{}, err
	}
	return scanRange(scanner, calcInclusive(start, end))
}

func scanRange(scanner *bufio.Scanner, times int) ([]string, error) {
	var result []string
	for i := 1; i <= times; i++ {
		if !scanner.Scan() {
			err := errors.New("out of file")
			return []string{}, err
		}
		t := scanner.Text()
		result = append(result, t)
	}
	return result, nil
}

func skipLine(scanner *bufio.Scanner, count int) (line int, err error) {
	for line = 1; line < count; line++ {
		scanned := scanner.Scan()
		if !scanned {
			err = errors.New("out of file")
			return 0, err
		}
	}

	return line, nil
}

func calcInclusive(start, end int) int {
	return end - start + 1
}
