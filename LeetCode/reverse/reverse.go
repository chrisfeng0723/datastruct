/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 14:13
 */

package reverse

import "math"

const MaxInt = math.MaxInt32
const MinInt = math.MinInt32

func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}

	if num > MaxInt || num < MinInt {
		return 0
	}
	return
}
