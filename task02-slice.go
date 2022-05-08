package homework

func reverse(input []int64) (result []int64) {
	result=make([]int64,len(input),len(input))
	copy(result,input)
	for i:=len(input)-1;i>=0;i--{
		result[len(input)-i-1] = input[i]
	}
	return result
}
