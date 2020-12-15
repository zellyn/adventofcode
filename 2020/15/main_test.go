package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		want1 int
		want2 int
	}{
		{
			name:  "example1",
			input: "0,3,6",
			want1: 436,
			want2: 175594,
		},
		{
			name:  "example2",
			input: "1,3,2",
			want1: 1,
			want2: 2578,
		},
		{
			name:  "example3",
			input: "2,1,3",
			want1: 10,
			want2: 3544142,
		},
		{
			name:  "example4",
			input: "1,2,3",
			want1: 27,
			want2: 261214,
		},
		{
			name:  "example5",
			input: "2,3,1",
			want1: 78,
			want2: 6895259,
		},
		{
			name:  "example6",
			input: "3,2,1",
			want1: 438,
			want2: 18,
		},
		{
			name:  "example7",
			input: "3,1,2",
			want1: 1836,
			want2: 362,
		},
		{
			name:  "input",
			input: "14,8,16,0,1,17",
			want1: 240,
			want2: 505,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := which(tt.input, 2020)
			if err != nil {
				t.Error(err)
			}

			if got1 != tt.want1 {
				t.Errorf("Want which(%q, 2020)=%d; got %d", tt.input, tt.want1, got1)
			}

			got2, err := which(tt.input, 30000000)
			if err != nil {
				t.Error(err)
			}

			if got2 != tt.want2 {
				t.Errorf("Want which(%q, 30000000)=%d; got %d", tt.input, tt.want2, got2)
			}

		})
	}
}
