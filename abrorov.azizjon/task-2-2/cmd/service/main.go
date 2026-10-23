package main

import (
    "bufio"
    "container/heap"
    "fmt"
    "os"
    "strconv"
)

type intMinHeap []int

func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intMinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *intMinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]

    return x
}

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

    numDishes, err := strconv.Atoi(token)
    if err != nil {
        return
    }

    dishRatings := make([]int, 0, numDishes)

    for range numDishes {
        token, hasNext = nextToken()
        if !hasNext {
            return
        }

        rating, err := strconv.Atoi(token)
        if err != nil {
            return
        }

        dishRatings = append(dishRatings, rating)
    }

    token, hasNext = nextToken()
    if !hasNext {
        return
    }

    k, err := strconv.Atoi(token)
    if err != nil {
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
