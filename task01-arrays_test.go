package homework

import "testing"

func TestAverage(t *testing.T){
	var arr=[15]float32{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	if average(arr)!=8{
		t.Errorf("test finished with error, result of test: ")
		t.Errorf("got %f, want %f", average(arr), 3.5)
	}
	t.Log("Task 1. Test 1 passed OK!")
}

func TestAverageTable(t *testing.T){
	var tests=[]struct{
		arr [15]float32
		value float32
	}{
		{[15]float32{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15},3.5},
		{[15]float32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},1},
		{[15]float32{2,2,2,2,2,2,2,2,2,2,2,2,2,2,2},2},
		{[15]float32{2,4,2,4,2,4,2,4,2,4,2,4,2,4,2},3},
		{[15]float32{5,5,5,5,5,5,5,5,5,5,5,5,5,5,5},5},
	}

for i,arg := range tests{
	if average(arg.arr)!=arg.value{
		t.Errorf("test finished with error, result of test %d: got %f, want %f", i, average(arg.arr),arg.value)
	}
	}
	t.Log("Task 1. Test 2. Test passed OK!")
}





