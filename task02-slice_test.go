package homework

import (
	"testing"
)
//Input -> (1, 2, 5, 15)
//Output -> (15, 5, 2, 1)
func TestReverse (t *testing.T){
	a:=reverse([]int64{1, 2, 5, 15,33})

	b:=[]int64{33,15, 5, 2, 1}
	t.Log(a)
	t.Log(b)
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			t.Errorf("test finished with error, in %d element: got %d, want %d",i, a[i],b[i])
		}
	}
	t.Log("Test passed OK!")
}
