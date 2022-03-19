package c_test

import (
	"testing"

	"github.com/MohabMohamed/challenge/c"
)

func TestWithOdd(t *testing.T) {
	if res := c.GetPalindromeNumber(9); res != "123454321" && res != "543212345" {
		t.Errorf("the palindrome number with length 9 should be 123454321 or 543212345 but got: %s\n", res)
	} else {
		t.Log("tests with odd number passed.")
	}
}

func TestWithEven(t *testing.T) {
	if res := c.GetPalindromeNumber(6); res != "123321" && res != "321123" {
		t.Errorf("the palindrome number with length 6 should be 123321 or 321123 but got: %s\n", res)
	} else {
		t.Log("tests with even number passed.")
	}
}
