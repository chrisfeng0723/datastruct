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


