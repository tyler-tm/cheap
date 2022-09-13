# CHeap

A Priority Queue in Go that uses concurrency to return a value from Pop() in linear time, before starting a [goroutine](https://go.dev/tour/concurrency/1) to re-heapify the underlying data structure [O(logn)].

This was mostly just a quick playground for me to try using a single-item capacity buffered channel as a lock for shared memory (instead of the standard sync/mutex); in practical application, one should probably just leverage the container/heap library (by implementing the heap interface, like in `internalPriorityQueue.go`) and use their own goroutines around Peek() and Pop().

