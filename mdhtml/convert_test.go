package mdhtml

import "testing"
import "fmt"

// func TestConvert(t *testing.T){
// 	Convert()
// 	if false {
// 			t.Error(
// 				"Got error",
// 			)
// 		}
// }
// 

func TestToLines(t *testing.T){
	type tests struct{
		Input []byte
		Output []string
	}
	var testPairs = []tests{
		{nil, nil},
		{[]byte(""), nil},
		{[]byte("\n"), nil},
		{[]byte("\n\n"), nil},
		{[]byte("first\n"),[]string{"first"}},
		{[]byte("\nfirst"),[]string{"first"}},
		{[]byte("\nfirst\n"),[]string{"first"}},
		{[]byte("first\nsecond"),[]string{"first", "second"}},
		{[]byte("first\nsecond\n"),[]string{"first", "second"}},
		{[]byte("first\n\nsecond"),[]string{"first", "second"}},
		
	}

	for _, pair := range testPairs {
		fmt.Println("Start")
		fmt.Println(pair)
		lines := toLines(pair.Input)
		fmt.Println(lines)
		fmt.Println(len(lines))
		fmt.Println(len(pair.Output))
		for index, line:= range lines {
			if line != pair.Output[index]{
			t.Error(
				"Error Spliting", string(pair.Input),
				"expected" , pair.Output,
				"got", lines, 
			)
			}
		}
	}
	

}