package mdhtml

import (
	"testing"
)

func TestToLines(t *testing.T) {
	type tests struct {
		Input  []byte
		Output []string
	}
	var testPairs = []tests{
		{nil, nil},
		{[]byte(""), []string{""}},
		{[]byte("a"), []string{"a"}},
		{[]byte("\n"), []string{""}},
		{[]byte("\n\n"), []string{"", ""}},
		{[]byte("first\n"), []string{"first"}},
		{[]byte("\nfirst"), []string{"", "first"}},
		{[]byte("\nfirst\n"), []string{"", "first"}},
		{[]byte("first\nsecond"), []string{"first", "second"}},
		{[]byte("first\nsecond\n"), []string{"first", "second"}},
		{[]byte("first\n\nsecond"), []string{"first", "", "second"}},
	}

	for _, pair := range testPairs {
		lines := toLines(pair.Input)
		for index, line := range lines {
			if line != pair.Output[index] {
				t.Error(
					"Error Spliting", string(pair.Input),
					"expected", pair.Output,
					"got", lines,
				)
			}
		}
	}
}

func TestShortenBlankLines(t *testing.T) {
	type tests struct {
		Input  []string
		Output []string
	}
	var testPairs = []tests{
		{nil, nil},
		{[]string{""}, []string{""}},
		{[]string{"\n"}, []string{""}},
		{[]string{" "}, []string{""}},
		{[]string{"   "}, []string{""}},
		{[]string{" \n  \t\n"}, []string{""}},
		{[]string{"\t"}, []string{""}},
		{[]string{"  \t"}, []string{""}},
		{[]string{"  \ta"}, []string{"  \ta"}},
		{[]string{" a \t"}, []string{" a \t"}},
	}
	for _, pair := range testPairs {
		lines := shortenBlankLines(pair.Input)
		for index, line := range lines {
			if line != pair.Output[index] {
				t.Error(
					"Error Shortening:", pair.Input,
					"expected", pair.Output,
					"got", lines,
				)
			}
		}
	}
}

func TestToBlocks(t *testing.T) {
	type tests struct {
		Input  []string
		Output [][]string
	}
	var testPairs = []tests{
		{nil, nil},
		{[]string{""}, [][]string{[]string{""}}},
		{[]string{""}, [][]string{[]string{"1"}}},
		{[]string{"Hi"}, [][]string{[]string{"Hi"}}},
		{[]string{"Hi", "how are", "you"}, [][]string{[]string{"Hi", "how are", "you"}}},
		{[]string{"Hi", "how are", "", "you"}, [][]string{[]string{"Hi", "how are"}, []string{"you"}}},
	}

	for _, pair := range testPairs {
		blocks := toBlocks(pair.Input)
		for index, block := range blocks {
			for block_index, block_line := range block {
				if block_line != pair.Output[index][block_index] {
					t.Error(
						"Error with blocks:", pair.Input,
						"expected", pair.Output,
						"got", block,
					)
				}
			}
		}
	}

}
