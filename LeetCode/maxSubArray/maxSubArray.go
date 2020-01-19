/**
 * @Author: dudu
 * @Description: 
 * @File:  maxSubArray.go
 * @Version: 1.0.0
 * @Date: 2020/1/10 17:02
 */
package maxSubArray

func maxSubArray(nums []int) int {
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			thisSum := 0
			for k := i; k <= j; k++ {
				thisSum += nums[k];
			}
			if thisSum > maxSum {
				maxSum = thisSum;
			}
		}
	}
	return maxSum
}
