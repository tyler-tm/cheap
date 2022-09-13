# CHeap

A Priority Queue in Go that uses concurrency to return from Pop() and Push() in linear time, before starting a [goroutine](https://go.dev/tour/concurrency/1) to re-heapify the underlying data structure [O(logn)]. Read/write calls while the heap is still being restructured will block until the queue is ready again.

This was mostly just a quick playground for me to try using a single-item capacity buffered channel as a lock for shared memory (instead of the standard sync/mutex); in practical application, one should probably just leverage the container/heap library (by implementing the heap interface, like in `internalPriorityQueue.go`) and use their own goroutines around Peek() and Pop().

## Usage

- Call the constructor with `c := NewCHeap()`
- `c.Push(object, priority)` your items
- `c.Pop()` them off the queue
- `c.Peek()` to see what's next, without affecting the queue

## Notes

This uses `any`, which was added in go1.18; `any` is just an alias for `interface{}`, so you can safely find and replace for use in projects referencing an older version of Go.
