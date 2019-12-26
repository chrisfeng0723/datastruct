/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 15:50
 */

package palindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var a, i, c int
	c = x
	for {
		a = c % 10
		c = c / 10
		i = i*10 + a
		if c == 0 {
			break
		}
	}
	if i == x {
		return true
	} else {
		return false
	}
}
