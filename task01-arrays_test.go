package homework

import "testing"

func TestAverage(t *testing.T){
	var arr=[6]float32{1,2,3,4,5,6}
	if average(arr)!=3.5{
		t.Errorf("test finished with error, result of test: ")
		t.Errorf("got %f, want %f", average(arr), 3.5)
	}
	t.Log("Test passed OK!")
}

func TestAverageTable(t *testing.T){
	var tests=[]struct{
		arr [6]float32
		value float32
	}{
		{[6]float32{1,2,3,4,5,6},3.5},
		{[6]float32{1,1,1,1,1,1},1},
		{[6]float32{2,2,2,2,2,2},2},
		{[6]float32{2,4,2,4,2,4},3},
		{[6]float32{5,5,5,5,5,5},5},
	}

for i,arg := range tests{
	if average(arg.arr)!=arg.value{
		t.Errorf("test finished with error, result of test %d: got %f, want %f", i, average(arg.arr),arg.value)
	}
	}
	t.Log("Test passed OK!")
}





