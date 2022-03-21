package twoSum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numSlices := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}

	targets := []int{9, 6, 6}

	expectedResults := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}

	for r := range numSlices {
		nums := numSlices[r]
		target := targets[r]
		expected := expectedResults[r]

		result := twoSum(nums, target)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %q, but got %q", expected, result)
		}

		fmt.Println("Success! Test case", r+1, "gave result", result)
	}
}
