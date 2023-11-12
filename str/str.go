package str

import (
	"fmt"
	"strconv"
	"strings"
)

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

func IsPalindrome(s string) bool {
	var l = 0
	var r = len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

const ALPHABET_SIZE = 26

var MAP_ALPHABET = map[byte]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}

func GroupAnagram(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		letterCounts := make([]int, ALPHABET_SIZE)
		for i := 0; i < len(s); i++ {
			letterCounts[MAP_ALPHABET[s[i]]] += 1
		}
		letters := make([]string, ALPHABET_SIZE)
		for i, count := range letterCounts {
			letters[i] = strconv.Itoa(count)
		}
		hash := strings.Join(letters, "#")

		groups[hash] = append(groups[hash], s)
	}

	fmt.Println(groups)

	var values [][]string
	for _, value := range groups {
		values = append(values, value)
	}
	return values
}
