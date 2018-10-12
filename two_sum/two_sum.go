package leetcode

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice
func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, v := range nums {
		t := target - v
		index, ok := temp[t]
		if ok {
			return []int{index, i}
		} else {
			temp[v] = i
		}
	}
	return []int{-1, -1}
}
