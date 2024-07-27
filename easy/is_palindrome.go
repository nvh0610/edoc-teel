package main

import (
	"strconv"
)

//
//func main() {
//	fmt.Println(isPalindrome(123))
//}

func IsPalindrome(num int) bool {
	if num < 10 {
		return false
	}

	str := strconv.Itoa(num)

	i, j := 0, len(str)-1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
