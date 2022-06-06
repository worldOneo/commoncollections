package commoncollections

import "testing"

const allocTestN = 10

func TestSlabAllocator(t *testing.T) {
	allocator := NewSlabAllocator[int](allocTestN)
	refs := make([]*AllocatorRef[int], allocTestN+1)
	for i := 0; i < allocTestN; i++ {
		refs[i] = allocator.Allocate()
	}
	// Allocate one more to trigger the reallocation
	refs[allocTestN] = allocator.Allocate()
	// Validity check
	for i := 0; i < allocTestN; i++ {
		*refs[i].Get() = i
	}
	*refs[allocTestN].Get() = allocTestN
	for i := 0; i < allocTestN+1; i++ {
		allocator.Free(refs[i])
	}
}
