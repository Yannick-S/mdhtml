package mdhtml

import (
	"fmt"
	"strings"
)


// md_bytes should be populated with the .md file. This function is the wrapper
// function, which calls all other functions. It returns a []byte with the .md
// file converted to .html syntax.
func Convert(md_bytes []byte) []byte{
	fmt.Println(string(md_bytes))
	return md_bytes
}


// splits the bytes into an array of lines
func toLines(b []byte) []string{
	stringArr := strings.Split(string(b), "\n")
	fmt.Println(stringArr)

	var nonNil []string
	for _, str := range stringArr {
		if str != "" {
			nonNil = append(nonNil, str)
		}
	}
	return nonNil 
}

