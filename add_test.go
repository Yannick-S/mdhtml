package mdhtml

import "testing"

type testpair struct {
	values []int
	sum int 
  }
var tests = []testpair{
	{[]int{1,2}, 3},
	{[]int{2,-1}, 1},
}
func TestAdd(t *testing.T){
	for _, pair := range tests {
		v := Add(pair.values[0], pair.values[1])
		if v != pair.sum {
			t.Error(
				"Error adding", pair.values,
				"expected" , pair.sum,
				"got", v, 
			)
		}
	}
}