package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap(15)
	heap.Insert(5)
	heap.Insert(7)
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(6)
	heap.Insert(4)

	heap.RemoveTop()
}
