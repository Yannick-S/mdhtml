package mdhtml

import (
	"strings"
)

// splits the bytes into an array of lines
func toLines(b []byte) []string{
	if b == nil {
		return nil
	}
	if len(b) == 0 {
		return []string{""}
	}
	if len(b) == 1 {
		if string(b) == "\n" {
			return []string{""}
		}
		return []string{string(b)}
	}

	if string(b)[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	stringArr := strings.Split(string(b), "\n")
	return stringArr
}

// lines which are complelety blank should be turned to just ""
func shortenBlankLines(strArr []string) []string{
	for index, line := range(strArr){
		trim := strings.TrimSpace(line)
		if trim == ""{
			strArr[index] = ""
		}
	}
	return strArr
}

// blocks are separated by "", and should be handled independ of one another
func toBlocks(strArr []string) [][]string {
	var finalArr  [][]string
	var thisBlock []string

	for _, line := range(strArr){
		//An empty line suggests a new block should be started.
		//Add what we have, and reset thisBlock
		if line == "" {
			finalArr = append(finalArr, thisBlock)
			thisBlock = nil
		} else {
			thisBlock = append(thisBlock, line)
		}
	}

	//Add final block
	finalArr = append(finalArr, thisBlock)

	return finalArr
}