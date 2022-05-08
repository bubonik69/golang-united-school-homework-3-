package homework

import "testing"

func TestTask1 (t *testing.T){
	var arr=[6]float32{1,2,3,4,5,6}
	if average(arr)!=3.5{
	t.Errorf("test finished with error, result of test: ")
		t.Errorf("got %f, want %d", average(arr), 3.5)
}
	t.Log("Test passed OK!")
}
