package mdhtml

import (
	"errors"
	"strings"
)
//This function returns a string, which indicates what type of a block this is
// This works well, since there are only a few types of blocks.

func determineBlock(block []string) (string, error) {
	if block == nil {
		return "", errors.New("No block type")
	}
	if len(block) <= 1 {
		return "", errors.New("No block type")
	}
	//White Space stuff first
	trim := strings.TrimSpace(block[0])
	if trim == "" {
		return "" , errors.New("First line empty")
	}
	if block[0][0] == '>' {
		return "Blockquotes", nil
	}

	return "", errors.New("No block type identified")
}