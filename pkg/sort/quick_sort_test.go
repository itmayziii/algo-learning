package sort_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{4, 8, 2, 1, 9, 7},
			[]int{1, 2, 4, 7, 8, 9},
		},
		{
			[]int{69, 20_000, 690, 420, 42, 71, 37},
			[]int{37, 42, 69, 71, 420, 690, 20_000},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := sort.QuickSort(tt.arr)
			if len(actual) != len(tt.expected) {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
				return
			}

			for j := range tt.expected {
				if tt.expected[j] != actual[j] {
					t.Errorf("actual %v, expected %v at", actual, tt.expected)
				}
			}
		})
	}
}
