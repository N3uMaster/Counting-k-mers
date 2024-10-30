package main

import "fmt"

func countkmers(sequence string, k int) map[string]int {
	var kmers = make(map[string]int)

	if k > len(sequence) || k <= 0 {
		return kmers
	}

	var tempSequence = sequence[:k]
	kmers[tempSequence]++

	for i := k; i < len(sequence); i++ {
		tempSequence = tempSequence[1:] + sequence[i:i+1]
		kmers[tempSequence]++
	}

	return kmers
}

func main() {
	fmt.Println(countkmers("ACGAGGTACGA", 3))
}