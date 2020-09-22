package main
import "testing"
func TestApp(t *testing.T){
	x:=2
	y:=2*x
	if y%3==0{
		t.Fail()
	}
}