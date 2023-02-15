package str

import "testing"

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
