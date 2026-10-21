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
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	defer func() { _ = out.Flush() }()

	next := func() (string, bool) {
		if !sc.Scan() {
			return "", false
		}
		return sc.Text(), true
	}

	tok, ok := next()
	if !ok {
		return
	}
	n, err := strconv.Atoi(tok)

	if err != nil {
		return
	}

	for d := 0; d < n; d++ {
		tok, ok = next()
		if !ok {
			return
		}
		k, err := strconv.Atoi(tok)

		if err != nil {
			return
		}

		lo, hi := minTemp, maxTemp
		for i := 0; i < k; i++ {
			op, ok := next()
			if !ok {
				return
			}
			val := strings.TrimLeft(op, "<>=")
			op = op[:len(op)-len(val)]
			if val == "" {
				if val, ok = next(); !ok {
					return
				}
			}
			t, err := strconv.Atoi(val)

			if err != nil {
				return
			}

			switch op {
			case ">=":
				if t > lo {
					lo = t
				}
			case "<=":
				if t < hi {
					hi = t
				}
			}

			if lo <= hi {
				_, _ = fmt.Fprintln(out, lo)
			} else {
				_, _ = fmt.Fprintln(out, -1)
			}
		}
	}
}
