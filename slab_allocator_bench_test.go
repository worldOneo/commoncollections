package commoncollections

import "testing"

const allocBenchBytes = 2<<11

func BenchmarkSlabAllocator_Alloc(b *testing.B) {
	allocator := NewSlabAllocator[[allocBenchBytes]byte](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		allocator.Allocate()
	}
}

func BenchmarkSlabAllocator_Free(b *testing.B) {
	allocator := NewSlabAllocator[[allocBenchBytes]byte](b.N)
	refs := make([]*AllocatorRef[[allocBenchBytes]byte], b.N)

	for i := 0; i < b.N; i++ {
		refs[i] = allocator.Allocate()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		allocator.Free(refs[i])
	}
}

func BenchmarkStandardMake_Alloc(b *testing.B) {
	ref := make([]byte, allocBenchBytes)
	for i := 0; i < b.N; i++ {
		ref = make([]byte, allocBenchBytes)
	}
	if len(ref) < allocBenchBytes {
		b.Fatal("len(ref) < allocBenchBytes")
	}
}
