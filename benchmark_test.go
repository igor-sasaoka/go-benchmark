package main

import (
	"math/rand"
	"testing"
)
const n = 1000000
func BenchmarkIteration(b *testing.B) {
    var arr [n]int 
    for i := 0; i < b.N; i++ {
        for i := 0; i < n; i++ {
            arr[i] = rand.Int() 
        }

        for i := 0; i < n; i++ {
            arr[i] = 0
        }
    }
}

func BenchmarkRange(b *testing.B) {
    var arr [n]int 
    for i := 0; i < b.N; i++ {
        for k := range arr {
            arr[k] = rand.Int() 
        }

        for k := range arr {
            arr[k] = 0
        }
    }
}
