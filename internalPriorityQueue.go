package cheap

func (h InternalPriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h InternalPriorityQueue) Len() int           { return len(h) }
func (h InternalPriorityQueue) Less(i, j int) bool { return h[i].Priority > h[j].Priority }

func (h *InternalPriorityQueue) Push(newElement any) {
	*h = append(*h, newElement.(*CHeapElement))
}

func (h *InternalPriorityQueue) Pop() any {
	prev := *h
	n := len(prev)
	result := prev[n-1]
	prev[n-1] = nil
	*h = prev[0 : n-1]
	return result
}
