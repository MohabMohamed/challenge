package b_test

import (
	"testing"

	"github.com/MohabMohamed/challenge/b"
)

func TestWithAscii(t *testing.T) {
	if res := b.IsAnagram("restful", "fluster"); res != true {
		t.Errorf("restful should be an anagram of fluster but got: %v\n", res)
	}

	if res := b.IsAnagram("funeral", "real fun"); res != true {
		t.Errorf("funeral should be an anagram of real fun but got: %v\n", res)
	}

	if res := b.IsAnagram("noon", "none"); res != false {
		t.Errorf("noon shouldn't be an anagram of none but got: %v\n", res)
	}

	t.Log("tests with ASCII passed.")
}

func TestWithUnicode(t *testing.T) {
	if res := b.IsAnagram("مهاب ", "هام ب"); res != true {
		t.Errorf("مهاب should be an anagram of هام ب but got: %v\n", res)
	}

	if res := b.IsAnagram("بطاطا", "بطاطس"); res != false {
		t.Errorf("بطاطس shouldn't be an anagram of بطاطا but got: %v\n", res)
	}
	t.Log("test with Unicode passed.")
}
