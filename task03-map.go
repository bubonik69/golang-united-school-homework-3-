package homework

import (
	"sort"
)

//function that returns map values sorted in order of increasing keys.
//Input -> {2: "a", 0: "b", 1: "c"}
//Output -> ["b", "c", "a"]
//Input -> {10: "aa", 0: "bb", 500: "cc"}
//Output -> ["bb", "aa", "cc"]

func sortMapValues(input map[int]string) (result []string) {
	//result=make([]string,len(input),len(input))
	var s []int
	for k,_:=range input{
		s = append(s, k)
	}
	sort.Ints(s)
	// s - сортированный массив keys
	for _,v := range s{
		result=append(result,input[v])
	}



	return result
}
