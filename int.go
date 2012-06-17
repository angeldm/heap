package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by changePriority and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A IntQueue implements heap.Interface and holds Items.
type IntQueue []*Item

func NewIntQueue(n int) IntQueue {
	return make(IntQueue, 0, n)
}

func (pq IntQueue) Len() int { return len(pq) }

func (pq IntQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq IntQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *IntQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	// To simplify indexing expressions in these methods, we save a copy of the
	// slice object. We could instead write (*pq)[i].
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(*Item)
	item.index = n
	a[n] = item
	*pq = a
}

func (pq *IntQueue) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	item.index = -1 // for safety
	*pq = a[0 : n-1]
	return item
}

// update is not used by the example but shows how to take the top item from
// the queue, update its priority and value, and put it back.
func (pq *IntQueue) update(value int, priority int) {
	item := heap.Pop(pq).(*Item)
	item.value = value
	item.priority = priority
	heap.Push(pq, item)
}

// changePriority is not used by the example but shows how to change the
// priority of an arbitrary item.
func (pq *IntQueue) changePriority(item *Item, priority int) {
	heap.Remove(pq, item.index)
	item.priority = priority
	heap.Push(pq, item)
}

// This example pushes 10 items into a IntQueue and takes them out in
// order of priority.
func main() {
	const nItem = 10
	// Random priorities for the items (a permutation of 0..9, times 11)).
	priorities := [nItem]int{
		77, 22, 44, 55, 11, 88, 33, 99, 00, 66,
	}
	values := [nItem]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	// Create a priority queue and put some items in it.
	pq := NewIntQueue(nItem)
	for i := 0; i < cap(pq); i++ {
		item := &Item{
			value:    values[i],
			priority: priorities[i],
		}
		heap.Push(&pq, item)
	}
	// Take the items out; should arrive in decreasing priority order.
	// For example, the highest priority (99) is the seventh item, so output starts with 99:"seven".
	for i := 0; i < nItem; i++ {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%d ", item.priority, item.value)
	}
	// Output:
	// 99:seven 88:five 77:zero 66:nine 55:three 44:two 33:six 22:one 11:four 00:eight
}
