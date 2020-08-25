package prose

import (
	"fmt"
	"testing"
	// "github.com/headfirstgo/prose"
)

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = '%s' \n wants '%s'", list, got, want)
}

func TestOneElement(t *testing.T) {
	list := []string{"apples"}
	fxnOut := JoinWithCommas(list)
	expectedOut := "apples"
	if fxnOut != expectedOut {
		t.Errorf(errorString(list, fxnOut, expectedOut))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apples", "oranges"}
	fxnOut := JoinWithCommas(list)
	expectedOut := "apples and oranges"
	if fxnOut != expectedOut {
		t.Errorf(errorString(list, fxnOut, expectedOut))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	fxnOut := JoinWithCommas(list)
	expectedOut := "apple, orange, and pear"
	if fxnOut != expectedOut {
		t.Errorf("JoinWithCommas(%#v) = '%s' \n wants '%s'", list, fxnOut, expectedOut)
	}
}
