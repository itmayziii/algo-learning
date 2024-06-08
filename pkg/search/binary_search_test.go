package search_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		expected bool
	}{
		{
			[]int{1, 4, 5},
			1,
			true,
		},
		{
			[]int{1, 4, 5},
			4,
			true,
		},
		{
			[]int{1, 4, 5},
			5,
			true,
		},
		{
			[]int{1, 4, 5},
			2,
			false,
		},
		{
			[]int{1, 10, 100, 1000, 10_000, 100_000},
			10,
			true,
		},
		{
			[]int{1, 10, 100, 1000, 10_000, 100_000},
			200,
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := search.BinarySearch(tt.arr, tt.val)
			if actual != tt.expected {
				t.Errorf("actual: %t, expected %t", actual, tt.expected)
			}
		})
	}
}
