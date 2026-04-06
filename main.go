package main

import (
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10000000
	CHUNKS = 8
)

func generateRandomElements(n int) []int {
	if n <= 0 {
		return []int{}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = r.Int()
	}
	return result
}

func maximum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	maxVal := s[0]
	for _, v := range s[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func maxChunks(s []int) int {
	if len(s) == 0 {
		return 0
	}
	chunkSize := len(s) / CHUNKS
	if chunkSize == 0 {
		return maximum(s)
	}
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(s)
		}
		go func(idx int, part []int) {
			defer wg.Done()
			maxValues[idx] = maximum(part)
		}(i, s[start:end])
	}
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	data := generateRandomElements(SIZE)

	startSingle := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startSingle).Microseconds()

	startMulti := time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(startMulti).Microseconds()

	println("Однопоточный максимум:", maxSingle, "время:", elapsedSingle, "мкс")
	println("Многопоточный максимум:", maxMulti, "время:", elapsedMulti, "мкс")
}
