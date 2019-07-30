package mdhtml

import (
	"errors"
	"testing"
)

func TestDetermineBlock(t *testing.T) {
	type tests struct {
		Input  []string
		Output string
		Error  error
	}
	var testPairs = []tests{
		{nil, "", errors.New("No block type")},
		{[]string{}, "", errors.New("No block type")},
		{[]string{"", "abc", "\tabcd"}, "", errors.New("First line empty")},
		{[]string{" ", "abc", "\tabcd"}, "", errors.New("First line empty")},
		{[]string{"\t", "abc", "\tabcd"}, "", errors.New("First line empty")},
		{[]string{"\n", "abc", "\tabcd"}, "", errors.New("First line empty")},
		{[]string{" \t  \n \t", "abc", "\tabcd"}, "", errors.New("First line empty")},
		{[]string{" \tbla  \n \t", "abc", "\tabcd"}, "", errors.New("First is space")},
		{[]string{">bla  \n \t", "  abc", "  \tabcd"}, "Blockquotes", nil},
		{[]string{">bla  \n \t", "  abc", "  \tabcd"}, "Blockquotes", nil},
		//{[]string{"abc","def","ghi"},	"paragraph", nil},
	}
	for _, pair := range testPairs {
		result, err := determineBlock(pair.Input)
		if err != nil {
			if err.Error() != pair.Error.Error() {
				t.Error("Error Finding Block Type:", pair.Input,
					"expected error", pair.Error,
					"got", err,
				)
			}
		}
		if result != pair.Output {
			t.Error("Error Finding Block Type:", pair.Input,
				"expected type", pair.Output,
				"got", result,
			)
		}
	}
}
