package commoncollections

import (
	"sync"
	"testing"
	"time"
)

const PRODUCERS = 100

func BenchmarkStdChannel(b *testing.B) {
	req := make(chan *sync.Mutex, PRODUCERS)
	for i := 0; i < PRODUCERS; i++ {
		go func() {
			mut := &sync.Mutex{}
			for {
				mut.Lock()
				req <- mut
				mut.Lock()
				mut.Unlock()
			}
		}()
	}

	time.Sleep(time.Second / 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(<-req).Unlock()
	}
	b.SetBytes(1)
}

func BenchmarkChannel(b *testing.B) {
	req := NewChannel[*sync.Mutex]()
	for i := 0; i < PRODUCERS; i++ {
		go func() {
			mut := &sync.Mutex{}
			for {
				mut.Lock()
				req.Send(mut)
				mut.Lock()
				mut.Unlock()
			}
		}()
	}

	time.Sleep(time.Second / 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req.Recv().Unlock()
	}
	b.SetBytes(1)
}

func TestChannel(t *testing.T) {
	channel := NewChannel[int]()
	for i := 0; i < 100; i++ {
		channel.Send(i)
	}
	for i := 0; i < 100; i++ {
		if v := channel.Recv(); v != i {
			t.Fatalf("channel.Recv() = %d, want %d", v, i)
		}
	}
	if v, ok := channel.TryRecv(); ok {
		t.Fatalf("channel.TryRecv() = %v, %v, want %v", v, true, false)
	}
}
