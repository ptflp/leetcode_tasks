package main

import "strings"

func numberOfBeams(bank []string) int {
	var res, prevCount, count, i int
	s := strings.Join(bank, "")
	perRow := len(bank[0])

	for i = range s {
		if s[i] == 49 {
			count++
		}
		if (i+1)%perRow == 0 {
			res += count * prevCount
			if count != 0 {
				prevCount = count
			}
			count = 0
		}
	}

	return res
}
