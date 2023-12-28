package main

import "testing"

var example1 = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		steps int
		want  any
	}{
		{
			name:  "2",
			input: example1,
			steps: 2,
			want:  4,
		},
		{
			name:  "6",
			input: example1,
			steps: 6,
			want:  16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _part1(tt.input, tt.steps); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		steps int
		want  any
	}{
		{
			name:  "6",
			steps: 6,
			input: example1,
			want:  16,
		},
		{
			name:  "10",
			steps: 10,
			input: example1,
			want:  50,
		},
		{
			name:  "50",
			steps: 50,
			input: example1,
			want:  1594,
		},
		{
			name:  "100",
			steps: 100,
			input: example1,
			want:  6536,
		},
		{
			name:  "500",
			steps: 500,
			input: example1,
			want:  167004,
		},
		{
			name:  "1000",
			steps: 1000,
			input: example1,
			want:  668697,
		},
		{
			name:  "5000",
			steps: 5000,
			input: example1,
			want:  16733044,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _part2(tt.input, tt.steps); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
