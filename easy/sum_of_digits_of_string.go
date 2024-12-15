package main

import (
	"strconv"
)

func getLucky(s string, k int) int {
	strExam := "abcdefghijklmnopqrstuvwxyz"
	m := map[string]string{}
	for index, value := range strExam {
		m[string(value)] = strconv.Itoa(index + 1)
	}

	str := ""
	for i := 0; i < len(s); i++ {
		str += m[string(s[i])]
	}

	for k > 0 {
		total := 0
		for i := 0; i < len(str); i++ {
			num, _ := strconv.Atoi(string(str[i]))
			total += num
		}
		str = strconv.Itoa(total)
		k--
	}

	result, _ := strconv.Atoi(str)

	return result
}
