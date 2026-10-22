package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp = 15
	maxTemp = 30
)

type tokenReader struct {
	scanner *bufio.Scanner
}

func (r *tokenReader) next() (string, bool) {
	if !r.scanner.Scan() {
		return "", false
	}

	return r.scanner.Text(), true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)

	defer func() { _ = out.Flush() }()

	reader := &tokenReader{scanner: scanner}

	token, hasNext := reader.next()
	if !hasNext {
		return
	}

	numDays, err := strconv.Atoi(token)
	if err != nil {
		return
	}

	for range numDays {
		token, hasNext = reader.next()
		if !hasNext {
			return
		}

		numConstraints, err := strconv.Atoi(token)
		if err != nil {
			return
		}

		low, high := minTemp, maxTemp

		for range numConstraints {
			operation, hasNext := reader.next()
			if !hasNext {
				return
			}

			valueStr := strings.TrimLeft(operation, "<>=")
			operation = operation[:len(operation)-len(valueStr)]

			if valueStr == "" {
				valueStr, hasNext = reader.next()
				if !hasNext {
					return
				}
			}

			temp, err := strconv.Atoi(valueStr)
			if err != nil {
				return
			}

			switch operation {
			case ">=":
				if temp > low {
					low = temp
				}
			case "<=":
				if temp < high {
					high = temp
				}
			}

			if low <= high {
				_, _ = fmt.Fprintln(out, low)
			} else {
				_, _ = fmt.Fprintln(out, -1)
			}
		}
	}
}
