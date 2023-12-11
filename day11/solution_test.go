package main

import "testing"

const exampleInput1 string = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput1,
			want:  374,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expansion int
		want      any
	}{
		{
			name:      "example1",
			input:     exampleInput1,
			expansion: 1,
			want:      374,
		},

		{
			name:      "example10",
			input:     exampleInput1,
			expansion: 10,
			want:      1030,
		},
		{
			name:      "example100",
			input:     exampleInput1,
			expansion: 100,
			want:      8410,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _part2(tt.input, tt.expansion); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2Opt(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expansion int
		want      any
	}{
		{
			name:      "example1",
			input:     exampleInput1,
			expansion: 1,
			want:      374,
		},

		{
			name:      "example10",
			input:     exampleInput1,
			expansion: 10,
			want:      1030,
		},
		{
			name:      "example100",
			input:     exampleInput1,
			expansion: 100,
			want:      8410,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _part2Opt(tt.input, tt.expansion); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
