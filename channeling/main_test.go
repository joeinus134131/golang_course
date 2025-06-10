package main

import "testing"

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan string)
		go func() { ch <- "test" }()
		<-ch
	}
}
