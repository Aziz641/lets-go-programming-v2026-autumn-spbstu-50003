package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type intMinHeap []int

func (h *intMinHeap) Len() int           { return len(*h) }
func (h *intMinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *intMinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *intMinHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *intMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func readInt(scanner *bufio.Scanner) (int, bool) {
	if !scanner.Scan() {
		return 0, false
	}

	val, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, false
	}

	return val, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)

	defer func() { _ = out.Flush() }()

	numDishes, ok := readInt(scanner)
	if !ok {
		return
	}

	dishRatings := make([]int, 0, numDishes)

	for range numDishes {
		rating, okRating := readInt(scanner)
		if !okRating {
			return
		}

		dishRatings = append(dishRatings, rating)
	}

	k, ok := readInt(scanner)
	if !ok {
		return
	}

	if k > len(dishRatings) || k <= 0 {
		return
	}

	h := &intMinHeap{}
	heap.Init(h)

	for _, rating := range dishRatings {
		if h.Len() < k {
			heap.Push(h, rating)
		} else if rating > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, rating)
		}
	}

	_, _ = fmt.Fprintln(out, (*h)[0])
}
