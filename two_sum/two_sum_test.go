package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	solution := []int{0, 1}
	result := twoSum(input, target)
	if result[0] != solution[0] || result[1] != solution[1] {
		t.Error("failed")
	}
	t.Log("result", result, " -- solution", solution)
}
