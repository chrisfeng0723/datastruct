/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 14:54
 */

package twosum

import (
	"testing"
)

func TestTwosum(t *testing.T) {

	for _, unit := range []struct {
		x      []int
		target int
		expect []int
	}{
		{[]int{2, 5, 7}, 9, []int{0, 2}},
		{[]int{1, 1, 1}, 5, []int{0, 0}},
	} {
		if actually := twoSum(unit.x, unit.target); IntSliceEqual(actually, unit.expect) == false {
			t.Errorf("reverse:[%v]:actually[%v]", unit, actually)
		}
	}

}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for key,val := range a{
		if val !=b[key]{
			return false
		}
	}
	return true
}
