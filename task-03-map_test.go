package homework

import "testing"

//function that returns map values sorted in order of increasing keys.
//Input -> {2: "a", 0: "b", 1: "c"}
//Output -> ["b", "c", "a"]
//Input -> {10: "aa", 0: "bb", 500: "cc"}
//Output -> ["bb", "aa", "cc"]



func TestSortMapValues (t *testing.T) {
	var tests = []struct {
		input  map[int]string
		output []string
	}{
		{map[int]string{2: "a", 0: "b", 1: "c"},[]string{"b", "c", "a"}},
		{map[int]string{10: "aa", 0: "bb", 500: "cc"},[]string{"bb", "aa", "cc"}},

	}

	for i:=0;i<len(tests);i++{
		k := tests[i]
		a:=make(map[int]string)
		a=k.input
		b:=k.output
		t.Log("map is - ", a)
		t.Log("want string slice: ", b)
		returnSlice:=make([]string,len(a),len(a))
		returnSlice=sortMapValues(a)
		t.Logf("slise %s,%d ", returnSlice, len(returnSlice))
		for j :=0;j<len(returnSlice);j++ {
			if b[j]!= returnSlice[j]{
				t.Errorf("test finished with error,map# %d in %d element: got %s, want %s",i, j, returnSlice[i],b[i])
			}
		}
		t.Log("Ok")
	}

	}


