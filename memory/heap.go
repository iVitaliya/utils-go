package memory

type maxHeap struct {
	array []int
}

// Initiates a new Heap
//
// Example:
//
//	package main
//
//	import "fmt"
//
//	func main() {
//		heap := memory.NewHeap()
//
//		heaps := []int{ 10, 20, 30, 5, 7, 9, 11, 35, 15, 48 }
//
//		for _, val := range heaps {
//			heap.Insert(val)
//			fmt.Println(heap)
//		}
//
//		for i := 0; i < 5; i++ {
//			heap.Extract()
//			fmt.Println(heap)
//		}
//	}
func NewHeap() *maxHeap {
	return &maxHeap{}
}

// Insert adds an element to the heap
func (h *maxHeap) Insert(key int) {
	h.array = append(h.array, key)

	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *maxHeap) Extract() int {
	var (
		extracted = h.array[0]
		last      = len(h.array) - 1
	)

	// Catches if the heap is less than 1 (so basically a 0 check)
	if len(h.array) == 0 {
		// Logger: couldn't extract an item from the heap as it's currently empty
		return -1
	}

	// Take out the last index and put it in root
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from the bottom to the top
func (h *maxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)

		index = parent(index)
	}
}

// maxHeapifyDown will heapify from the top to the bottom
func (h *maxHeap) maxHeapifyDown(index int) {
	var (
		lastIndex      = len(h.array) - 1
		l, r           = left(index), right(index)
		childToCompare = 0
	)

	// Loop over the parent as the index has only one child
	for l <= lastIndex {
		if l == lastIndex { // When the left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // When left child of the parent is larger than the right child
			childToCompare = l
		} else { // When right child of the parent is larger than the left child
			childToCompare = r
		}

		// Compare the array value of the current index to the larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)

			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index of the heap
func parent(ind int) int {
	return (ind - 1) / 2
}

// Get the left child index of the parent
func left(ind int) int {
	return 2*ind + 1
}

// Get the right child index of the parent
func right(ind int) int {
	return 2*ind + 2
}

// Swap the provided keys with each other
func (h *maxHeap) swap(ind1, ind2 int) {
	h.array[ind1], h.array[ind2] = h.array[ind2], h.array[ind1]
}
