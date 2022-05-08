package homework

func average(input [15]float32) (result float32) {
	for i:=0;i<(len(input));i++{
		result= result+input[i]
	}
	return result/15
}
