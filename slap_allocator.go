package commoncollections

// AllocatorRef is a reference to an allocated object
// allocated by a SlapAllocator.
// The object is guaranteed to be valid until the
// AllocatorRef is freed.
// The value of the object is undefined after its
// Freed.
// The value is not zeroed.
type AllocatorRef[T any] struct {
	pos  int
	next int
	ref  T
}

// SlapAllocator is a memory allocator that
// allocates a fixed size block of memory
// and returns references to the allocated
// values.
//
// The allocator might grow in size if more
// memory is needed.
// It is not threadsafe.
// Locks and bound checks are not implemented
// and must be done by the caller if needed.
type SlapAllocator[T any] struct {
	next int
	data []AllocatorRef[T]
}

// NewSlapAllocator creates a new SlapAllocator which holds
// size number of elements.
func NewSlapAllocator[T any](size int) *SlapAllocator[T] {
	allocator := &SlapAllocator[T]{
		data: make([]AllocatorRef[T], size),
	}

	for i := 0; i < size; i++ {
		allocator.data[i].pos = i
		allocator.data[i].next = i + 1
	}

	allocator.data[size-1].next = -1
	return allocator
}

// Allocate allocates a new object from the allocator.
// Might trigger a reallocation if the allocator is exhausted.
func (allocator *SlapAllocator[T]) Allocate() *AllocatorRef[T] {
	if allocator.next == -1 {
		newSize := len(allocator.data) * 2
		newData := make([]AllocatorRef[T], newSize)
		for i := 0; i < newSize; i++ {
			newData[i].pos = i
			newData[i].next = i + 1
		}
		newData[newSize-1].next = -1
		copy(newData, allocator.data)
		allocator.data = newData
		allocator.next = newSize - 1
	}
	ref := &allocator.data[allocator.next]
	allocator.next = ref.next
	return ref
}

// Free returns the object to the allocator.
func (allocator *SlapAllocator[T]) Free(ref *AllocatorRef[T]) {
	ref.next = allocator.next
	allocator.next = ref.pos
}

// Get returns the object referenced by the AllocatorRef.
func (ref *AllocatorRef[T]) Get() *T {
	return &ref.ref
}
