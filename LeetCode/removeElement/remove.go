/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/27 10:54
 */

package removeElement

func removeElement(nums []int, val int) int {
	length := len(nums)
	temp := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[temp] = nums[i]
			temp = temp + 1
		}
	}
	return temp
}

func removeElementT(nums []int, val int) int {
	low, high := 0, len(nums)
	for low < high {
		if nums[low] == val {
			nums[low], nums[high-1] = nums[high-1], nums[low]
			high = high - 1
		} else {
			low = low + 1

		}
	}

	return high
}
