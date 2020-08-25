package prose

import (
	"fmt"
	"testing"
	// "github.com/headfirstgo/prose"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{""}, want: ""},
		testData{list: []string{"apples"}, want: "apples"},
		testData{list: []string{"apples", "oranges"}, want: "apples and oranges"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		fxnOut := JoinWithCommas(test.list)
		if fxnOut != test.want {
			t.Errorf(errorString2(test.list, fxnOut, test.want))
		}
	}
}

func errorString2(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = '%s' \n wants '%s'", list, got, want)
}

func TestOneElement2(t *testing.T) {
	list := []string{"apples"}
	expectedOut := "apples"
	// to reduce this repeated code we created TestJoinWithCommas
	fxnOut := JoinWithCommas(list)
	if fxnOut != expectedOut {
		t.Errorf(errorString2(list, fxnOut, expectedOut))
	}
}

func TestTwoElements2(t *testing.T) {
	list := []string{"apples", "oranges"}
	expectedOut := "apples and oranges"
	// to reduce this repeated code we created TestJoinWithCommas
	fxnOut := JoinWithCommas(list)
	if fxnOut != expectedOut {
		t.Errorf(errorString2(list, fxnOut, expectedOut))
	}
}

func TestThreeElements2(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	expectedOut := "apple, orange, and pear"
	// to reduce this repeated code we created TestJoinWithCommas
	fxnOut := JoinWithCommas(list)
	if fxnOut != expectedOut {
		t.Errorf(errorString2(list, fxnOut, expectedOut))
	}
}
