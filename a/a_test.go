package a_test

import (
	"testing"

	"github.com/MohabMohamed/challenge/a"
)

func TestWithExistingNeedle(t *testing.T) {
	if res := a.FindNeedleIndex("consciousness", "ss"); res != 11 {
		t.Errorf("ss should be found in index 11 in the word consciousness but found in index: %d\n", res)
	}
	t.Log("test with existing needle passed.")
}

func TestWithNotExistingNeedle(t *testing.T) {
	if res := a.FindNeedleIndex("hello", "go"); res != -1 {
		t.Errorf("go shouldn't be found in the word hello but found in index: %d\n", res)
	}
	t.Log("test with not existing needle passed.")
}

func TestWithEmptyStrings(t *testing.T) {
	if res := a.FindNeedleIndex("hello", ""); res != -1 {
		t.Errorf("should get -1 from empty needle but found : %d\n", res)
	}
	if res := a.FindNeedleIndex("", "potato"); res != -1 {
		t.Errorf("should get -1 from empty haystack but found : %d\n", res)
	}
	if res := a.FindNeedleIndex("", ""); res != -1 {
		t.Errorf("should get -1 if the  strings are empty by got: %d\n", res)
	}
	t.Log("test with not existing needle passed.")
}
