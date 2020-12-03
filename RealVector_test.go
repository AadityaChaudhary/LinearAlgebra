package LinearAlgebra

import (
	"fmt"
	"testing"
)


func TestPrint(t *testing.T) {
	var tests = []struct {
		a *RealVector
	}{
		{&RealVector{[]float64{0,0,0}}},
		{&RealVector{[]float64{1,-1,0}}},
		{&RealVector{[]float64{0.5, 4.2, 5.5}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.a)
		t.Run(testname, func(t *testing.T) {
			t.Logf("%s", tt.a)
		})
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		a *RealVector
		b *RealVector
		c bool
	}{
		{&RealVector{[]float64{0,0,0}},&RealVector{[]float64{0,0,0}}, true},
		{&RealVector{[]float64{1,-1,0}},&RealVector{[]float64{1,1,0}}, false},
		{&RealVector{[]float64{0.5, 4.2, 5.5}},&RealVector{[]float64{0.5,4.2,5.5}}, true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%s, %s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			eql := tt.a.Equal(tt.b)

			if eql != tt.c {
				t.Errorf("got %t, want %t", eql, tt.c)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a *RealVector
		b *RealVector
		c *RealVector
	}{
		{&RealVector{[]float64{0,0,0}},&RealVector{[]float64{0,0,0}},&RealVector{[]float64{0,0,0}}},
		{&RealVector{[]float64{1,1,0}},&RealVector{[]float64{0,0,1}},&RealVector{[]float64{1,1,1}}},
		{&RealVector{[]float64{0.5,20,500}},&RealVector{[]float64{0.7,0.6,70}},&RealVector{[]float64{1.2,20.6,570}}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%s,%s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans,err := tt.a.Add(tt.b)
			if err != nil {
				t.Errorf("ans is nil")
			}
			if !ans.Equal(tt.c) {
				t.Errorf("got %s, want %s", ans, tt.c)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	var tests = []struct {
		a *RealVector
		b *RealVector
		c *RealVector
	}{
		{&RealVector{[]float64{0,0,0}},&RealVector{[]float64{0,0,0}},&RealVector{[]float64{0,0,0}}},
		{&RealVector{[]float64{1,1,0}},&RealVector{[]float64{1,0,1}},&RealVector{[]float64{0,1,-1}}},
		{&RealVector{[]float64{0.5,20,500}},&RealVector{[]float64{0.7,0.6,70}},&RealVector{[]float64{-0.2,19.4,430}}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%s,%s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans,err := tt.a.Minus(tt.b)
			if err != nil {
				t.Errorf("ans is nil")
			}
			if !ans.Equal(tt.c) {
				t.Errorf("got %s, want %s", ans, tt.c)
			}
		})
	}
}