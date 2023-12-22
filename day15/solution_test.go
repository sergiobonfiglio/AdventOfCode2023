package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int32
	}{
		{
			name:  "example",
			input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
			want:  1320,
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
		{
			name:  "example",
			input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
			want:  145,
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

func Test_hash(t *testing.T) {

	tests := []struct {
		args string
		want int32
	}{
		{"HASH", 52},
		{"rn=1", 30},
		{"cm-", 253},
		{"qp=3", 97},
		{"cm=2", 47},
		{"qp-", 14},
		{"pc=4", 180},
		{"ot=9", 9},
		{"ab=5", 197},
		{"pc-", 48},
		{"pc=6", 214},
		{"ot=7", 231},
		{"rn", 0},
		{"cm", 0},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := hash(tt.args); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
