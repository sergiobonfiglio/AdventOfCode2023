package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}

func BenchmarkPart2_opt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2Opt(input)
	}
}