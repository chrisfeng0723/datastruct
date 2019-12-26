/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 16:12
 */

package palindrome

import (
	"testing"
)

func TestPalindrome(t *testing.T) {

	for _, unit := range []struct {
		x      int
		expect bool
	}{
		{121, true},
		{-121, false},
		{10,false},
	} {
		if actually := isPalindrome(unit.x); actually != unit.expect {
			t.Errorf("Palindrome:[%v]:actually[%v]", unit, actually)
		}
	}

}
