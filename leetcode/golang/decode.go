package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	index := 0
	for index < len(s) {
		if s[index] >= '0' && s[index] <= '9' {

			temp := ""
			balance := 1
			indexEnd := index
			number := 0
			for s[indexEnd] >= '0' && s[indexEnd] <= '9' {
				number = number*10 + int(s[indexEnd]-'0')
				indexEnd++
			}
			indexEnd++
			for indexEnd < len(s) {
				switch s[indexEnd] {
				case '[':
					balance++
				case ']':
					balance--
				}

				temp += string(s[indexEnd])

				if balance == 0 {
					break
				}
				indexEnd++
			}
			temp = temp[0 : len(temp)-1]

			s = s[:index] + strings.Repeat(temp, number) + s[indexEnd+1:]
			index = 0
			continue
		}
		index++
	}

	return s
}
func main() {
	fmt.Println(decodeString("2[leetcode]"))
}
