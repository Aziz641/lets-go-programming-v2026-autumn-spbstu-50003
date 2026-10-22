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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)

	defer func() { _ = out.Flush() }()

	nextToken := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}

		return scanner.Text(), true
	}

	token, hasNext := nextToken()
	if !hasNext {
		return
	}

	numDays, err := strconv.Atoi(token)

	if err != nil {
		return
	}

	for range numDays {
		token, hasNext = nextToken()
		if !hasNext {
			return
		}

		numConstraints, err := strconv.Atoi(token)

		if err != nil {
			return
		}

		low, high := minTemp, maxTemp

		for range numConstraints {
			operation, ok := nextToken()
			if !ok {
				return
			}

			valueStr := strings.TrimLeft(operation, "<>=")
			operation = operation[:len(operation)-len(valueStr)]

			if valueStr == "" {
				valueStr, ok = nextToken()
				if !ok {
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
