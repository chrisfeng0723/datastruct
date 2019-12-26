/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2019/12/26 14:22
 */

package reverse

import (
	"testing"
)

func TestReverse(t *testing.T){

	for _,unit := range []struct{
		x int
		expect int
	}{
		{-2147483648,0},
		{-2147483647,0},
		{2147483647,0},
		{-21474,-47412},
		{21474,47412},
		{0,0},
	}{
		if actually := reverse(unit.x); actually !=unit.expect{
			t.Errorf("reverse:[%v]:actually[%v]",unit,actually)
		}
	}

}
