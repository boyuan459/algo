package str

import (
	"reflect"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	var sr, ta = "ab##c", "ad##c"
	var equal = BackspaceCompare(sr, ta)

	if !equal {
		t.Errorf("expect true, but got %v", equal)
	}
}

func TestLengthOfLongestSubString(t *testing.T) {
	var s = "abcabcbb"
	length := LengthOfLongestSubString(s)

	if length != 3 {
		t.Errorf("expected 3, but got %v", length)
	}

	t.Log("----------------------------")

	var s2 = "bbbbb"
	length2 := LengthOfLongestSubString(s2)

	if length2 != 1 {
		t.Errorf("expected 1, but got %v", length2)
	}
}

func TestIsPalidrome(t *testing.T) {
	var s = "abcba"
	var isPal = IsPalindrome(s)

	if !isPal {
		t.Errorf("expected true, but got %v", isPal)
	}
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	values := GroupAnagram(strs)

	expected := [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}

	if reflect.DeepEqual(values, expected) {
		t.Logf("expected as %v", values)
	} else {
		t.Errorf("expected %v, but got %v", expected, values)
	}
}

func TestRomanToInteger(t *testing.T) {
	s := "MCMXCIV"

	sum := RomanToInteger(s)

	if sum != 1994 {
		t.Errorf("expected 1995, but got %v", sum)
	}
}
