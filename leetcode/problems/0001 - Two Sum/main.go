package main

import "fmt"

func main() {

	// Example 1:
	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Println("result1:", result1)

	// Example 2:
	// Input: nums = [3,2,4], target = 6
	// Output: [1,2]
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Println("result2:", result2)

	// Example 3:
	// Input: nums = [3,3], target = 6
	// Output: [0,1]
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Println("result3:", result3)
}

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		comp := target - num
		if iComp, exists := indexMap[comp]; exists {
			return []int{iComp, i}
		}
		indexMap[num] = i
	}
	return nil
}
