package main

import "fmt"

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	var leftChar, rightChar byte
	for left < right {
		leftChar, rightChar = s[left], s[right-1]
		if leftChar == rightChar {
			for left < right && s[left] == leftChar {
				left++
			}

			for left < right && s[right] == rightChar {
				right--
			}
		} else {
			break
		}
	}

	return right - left + 1
}

func main() {
	n := minimumLength("aabccabba")
	fmt.Println(n)
}
