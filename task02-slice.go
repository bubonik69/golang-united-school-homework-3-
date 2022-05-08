package homework

func reverse(input []int64) (result []int64) {
	for i:=0;i<len(input);i++{
		result = append(result,input[len(input)-i-1])
	}
	return
}
