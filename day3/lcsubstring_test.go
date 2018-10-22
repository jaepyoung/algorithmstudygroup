package lcs

import "testing"

func Testlcs(t *testing.T) {
	total := getlongestring("testtest", "testhaha")
	if total != 2 {
		t.Errorf("Return value was incorrect, got: %d, want: %d.", total, 2)
	}
}
