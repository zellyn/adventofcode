package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	var testdata = []struct {
		input string
		steps int
	}{
		{
			input: "example1",
			steps: 23,
		},

		{
			input: "example2",
			steps: 58,
		},
		{
			input: "input",
			steps: 644,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.input, func(t *testing.T) {
			s, err := newState(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			// s.draw()
			// fmt.Println(s.label(vec2{1, 8}))
			// fmt.Println(s.label(vec2{0, 8}))
			// fmt.Println(s.label(vec2{9, 0}))
			// fmt.Println(s.label(vec2{9, 1}))
			// fmt.Println(s.label(vec2{0, 0}))
			// fmt.Println(s.m[vec2{2, 8}], s.m[vec2{2, 8}].neighborString())
			// fmt.Println(s.m[vec2{9, 6}], s.m[vec2{9, 6}].neighborString())
			// fmt.Println(s.m[vec2{2, 13}], s.m[vec2{2, 13}].neighborString())
			// fmt.Println(s.m[vec2{6, 10}], s.m[vec2{6, 10}].neighborString())
			// fmt.Println(s.m[vec2{2, 15}], s.m[vec2{2, 15}].neighborString())
			// fmt.Println(s.m[vec2{11, 12}], s.m[vec2{11, 12}].neighborString())
			// fmt.Println("start", s.m[vec2{9, 2}], s.m[vec2{9, 2}].neighborString())
			// fmt.Println("end", s.m[vec2{13, 16}], s.m[vec2{13, 16}].neighborString())

			steps, err := s.minSteps()
			if err != nil {
				t.Fatal(err)
			}
			if steps != tt.steps {
				t.Errorf("want minSteps()=%d; got %d", tt.steps, steps)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	var testdata = []struct {
		input string
		steps int
	}{
		{
			input: "example1",
			steps: 26,
		},
		{
			input: "example3",
			steps: 396,
		},
		{
			input: "input",
			steps: 7798,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.input, func(t *testing.T) {
			s, err := newState(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			steps, err := s.minRecursiveSteps()
			if err != nil {
				t.Fatal(err)
			}
			if steps != tt.steps {
				t.Errorf("want minRecursiveSteps()=%d; got %d", tt.steps, steps)
			}
		})
	}
}
