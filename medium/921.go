package main

import "strings"

func minAddToMakeValid(s string) int {
	text := ""
	for {
		b := replaceString(s)
		if b != text {
			text = b
			s = b
		} else {
			break
		}
	}
	return len(text)
}

func replaceString(input string) string {
	return strings.Replace(input, "()", "", 1)
}
