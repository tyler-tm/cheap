package cheap

type CHeapElement struct {
	Data     any
	Priority int
}

type InternalPriorityQueue []*CHeapElement

type CHeap struct {
	elements InternalPriorityQueue
	lock     chan int
}
