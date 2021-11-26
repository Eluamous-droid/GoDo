package tools

import (
	"testing"
)

func TestArrayToString(t *testing.T){

	expected := "abc\ndef\nghi\n"
	input := []string {"abc","def","ghi"}
	output := BuildStringFromArray(input)
	if output != expected{
		t.Errorf("Output was not as expected, output was "+output+", expected was "+expected)
	}

}