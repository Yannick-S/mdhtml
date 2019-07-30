package mdhtml

import (
	"strings"
)

func Convert(b []byte) [][]string {
	lines := toLines(b)
	lines = shortenBlankLines(lines)
	lines = spaceToTab(lines)
	blocks := toBlocks(lines)

	return blocks
}

// splits the bytes into an array of lines
func toLines(b []byte) []string {
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
func shortenBlankLines(strArr []string) []string {
	for index, line := range strArr {
		trim := strings.TrimSpace(line)
		if trim == "" {
			strArr[index] = ""
		}
	}
	return strArr
}

// replaces 4 spaces with tabs, for easier processing later
func spaceToTab(strArr []string) []string {
	for index := range strArr {
		nrSpaces := 0
		nrTabs := 0
		for _, c := range strArr[index] {
			if c == ' ' {
				nrSpaces++
			} else if c == '\t' {
				nrTabs++
			} else {
				break
			}
		}
		strArr[index] = strArr[index][nrTabs+nrSpaces:]
		nrSpaces = nrSpaces / 4

		for i := 0; i < nrSpaces+nrTabs; i++ {
			strArr[index] = "\t" + strArr[index]
		}
	}
	return strArr
}

// remove starting spaces
//func removeFrontSpaces(strArr []string) []string {
//	for index := range strArr {
//		for len(strArr[index]) > 0 && strArr[index][0:1] == " " {
//			strArr[index] = strArr[index][1:]
//		}
//	}
//	return strArr
//}

// blocks are separated by "", and should be handled independ of one another
func toBlocks(strArr []string) [][]string {
	var finalArr [][]string
	var thisBlock []string

	for _, line := range strArr {
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
