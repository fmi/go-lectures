package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue []int

func (q *Queue) Push(item int) {
	*q = append(*q, item)
}

func (q *Queue) Pop() int {
	if !q.IsEmpty() {
		item := (*q)[0]
		*q = (*q)[1:len(*q)]
		return item
	}
	return 0
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

var (
	cond      *sync.Cond = sync.NewCond(new(sync.Mutex))
	workQueue Queue
	wg        sync.WaitGroup
)

func produce() {
	i := 0
	for {
		time.Sleep(1 * time.Millisecond)
		cond.L.Lock()
		workQueue.Push(i)
		i++
		cond.Signal()
		cond.L.Unlock()
	}
}

func consume(name string) {
	for {
		cond.L.Lock()
		for workQueue.IsEmpty() {
			cond.Wait()
		}
		item := workQueue.Pop()
		cond.L.Unlock()
		fmt.Println(name, item)
	}
}

func main() {
	wg.Add(1)
	go consume("First")
	go consume("Second")
	go consume("Third")
	go produce()
	wg.Wait()
}
