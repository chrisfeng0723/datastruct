/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 17:13
 */

package array

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i-1]
		if temp == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
