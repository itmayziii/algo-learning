package search_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/search"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	tests := []struct {
		arr      []bool
		expected int
	}{
		{
			[]bool{true},
			0,
		},
		{
			[]bool{false, true},
			1,
		},
		{
			[]bool{false, true, true},
			1,
		},
		{
			[]bool{false, false, true, true, true},
			2,
		},
		{
			[]bool{false, false, false, false, false, true},
			5,
		},
		{
			[]bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true},
			10,
		},
		{
			[]bool{},
			-1,
		},
		{
			[]bool{false},
			-1,
		},
		{
			[]bool{false, false, false},
			-1,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			actual := search.TwoCrystalBalls(tt.arr)
			if actual != tt.expected {
				t.Errorf("actual: %d, expected %d", actual, tt.expected)
			}
		})
	}
}
