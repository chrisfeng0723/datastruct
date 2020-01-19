/**
 * @Author: dudu
 * @Description: 
 * @File:  authcode
 * @Version: 1.0.0
 * @Date: 2020/1/15 18:15
 */
package experiment

import "math/rand"

func GetCode(length int) string {

	charArr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "X", "Y", "Z", "W", "S", "R", "T"}
	charLen := len(charArr)
	var code string
	for {
		temp := rand.Intn(charLen - 1)
		code = code + charArr[temp]
		length = length - 1
		if length < 1 {
			break
		}
	}

	return ""
}
