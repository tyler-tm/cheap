package cheap

import (
	"container/heap"
	"errors"
)

func NewCHeap() *CHeap {
	c := &CHeap{
		elements: make(InternalPriorityQueue, 0),
		lock:     make(chan int, 1),
	}
	heap.Init(&c.elements)
	return c
}

func (c *CHeap) Push(data any, priority int) {
	c.lock <- 1
	go c.pushAndUnlock(data, priority)
}

func (c *CHeap) Pop() (any, error) {
	c.lock <- 1
	if len(c.elements) < 1 {
		<-c.lock
		return nil, errors.New("Popped from empty CHeap")
	}
	result := c.elements[0].Data
	go c.popAndUnlock()
	return result, nil
}

func (c *CHeap) Peek() (any, error) {
	c.lock <- 1
	if len(c.elements) < 1 {
		<-c.lock
		return nil, errors.New("Peeked from empty CHeap")
	}
	result := c.elements[0].Data
	<-c.lock
	return result, nil
}

func (c *CHeap) pushAndUnlock(data any, priority int) {
	newElement := &CHeapElement{Data: data, Priority: priority}
	heap.Push(&c.elements, newElement)
	<-c.lock
}

func (c *CHeap) popAndUnlock() {
	_ = heap.Pop(&c.elements).(*CHeapElement).Data
	<-c.lock
}
