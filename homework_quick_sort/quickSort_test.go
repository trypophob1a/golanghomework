package homeworkquicksort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQs(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{2, 3, 1, 4}, expected: []int{1, 2, 3, 4}},
		{input: []int{7, 2, 3, -1, 0, 4}, expected: []int{-1, 0, 2, 3, 4, 7}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			result := QuickSort(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
