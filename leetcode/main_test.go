package main

import "testing"

func BenchmarkDecodeAtIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeAtIndex("y959q969u3hb22odq595", 222280369)
	}
}
