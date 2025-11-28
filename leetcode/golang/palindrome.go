package main

import "fmt"

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == x || reversed/10 == x
}
func main() {
	fmt.Println(isPalindrome(121))
}
