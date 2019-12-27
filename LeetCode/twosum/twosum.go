/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 14:45
 */

package twosum


func twoSum(nums []int, target int) []int {
	match := make(map[int]int)
	for key, val := range nums {
		member := target - val
		if _, ok := match[member]; ok  {
			return []int{match[member],key}
		}
		match[val] = key
	}
	return []int{0, 0}
}

//有序数组求和，利用有序性
func twoSumOrder(nums []int, target int) []int {
	length := len(nums) - 1
	for low, high := 0, length; low < high; {

		temp := nums[low] + nums[high]
		if temp == target {
			return []int{low, high}
		}
		if temp > target {
			high = high - 1
		}
		if temp < target {
			low = low + 1
		}

	}
	return []int{0, 0}
}


