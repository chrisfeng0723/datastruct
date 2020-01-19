/**
 * @Author: dudu
 * @Description: 
 * @File:  array.go
 * @Version: 1.0.0
 * @Date: 2020/1/8 10:49
 */
package experiment

import "fmt"

func ArrayRange(){
	for i,c :=range []rune("世界abc"){
		fmt.Println(i,c)
	}

}