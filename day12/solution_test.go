package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{

		{name: "example-row1", input: "???.### 1,1,3", want: 1},
		{name: "example-row2", input: ".??..??...?##. 1,1,3", want: 4},
		{name: "example-row3", input: "?#?#?#?#?#?#?#? 1,3,1,6", want: 1},
		{name: "example-row4", input: "????.#...#... 4,1,1", want: 1},
		{name: "example-row5", input: "????.######..#####. 1,6,5", want: 4},
		{name: "example-row6", input: "?###???????? 3,2,1", want: 10},
		{name: "worst", input: "??????.????????????? 3,4,7", want: 12},
		{
			name: "example",
			input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`,
			want: 21,
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
		name  string
		input string
		want  any
	}{
		{name: "example-row1", input: "???.### 1,1,3", want: 1},
		{name: "example-row2", input: ".??..??...?##. 1,1,3", want: 16384},
		{name: "example-row3", input: "?#?#?#?#?#?#?#? 1,3,1,6", want: 1},
		{name: "example-row4", input: "????.#...#... 4,1,1", want: 16},
		{name: "example-row5", input: "????.######..#####. 1,6,5", want: 2500},
		{name: "example-row6", input: "?###???????? 3,2,1", want: 506250},
		{
			name: "example",
			input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`,
			want: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
