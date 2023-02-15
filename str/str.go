package str

import "fmt"

func BackspaceCompare(s string, t string) bool {
	var p1, p2 = len(s) - 1, len(t) - 1

	for p1 >= 0 && p2 >= 0 {
		fmt.Printf("%c, %c\n", s[p1], t[p2])
		if s[p1] == '#' {
			var backCount = 2

			for backCount > 0 {
				backCount -= 1
				p1 -= 1

				if p1 >= 0 && s[p1] == '#' {
					backCount += 2
				}
			}
		}

		if t[p2] == '#' {
			var backCount = 2

			for backCount > 0 {
				backCount -= 1
				p2 -= 1

				if p2 >= 0 && t[p2] == '#' {
					backCount += 2
				}
			}
		}
		if p1 >= 0 && p2 >= 0 && s[p1] != t[p2] {
			return false
		}

		p1 -= 1
		p2 -= 1
	}

	return true
}

func LengthOfLongestSubString(s string) int {
	var left, right = 0, 0

	var seen = make(map[byte]int)
	var maxLength = 0

	for right < len(s) {
		fmt.Printf("right %d, seen %v\n", right, seen)
		if idx, ok := seen[s[right]]; ok && idx >= left {
			fmt.Printf("char %c,  index %d\n", s[right], idx)
			left = idx + 1
		}
		seen[s[right]] = right
		var length = right - left + 1
		if length > maxLength {
			maxLength = length
		}
		right += 1
	}

	fmt.Println("---------------------------")
	return maxLength
}
