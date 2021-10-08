package main

import (
	"testing"

)

func IsEqual(r1 float64, r2 float64, t *testing.T){
	if r1 == -1.881966011250105 && r2 == -4.118033988749895{
		t.Log("Correct")
	}else{
		t.Error("Incorrect")
	}
}
func TestDiscriminant(t *testing.T){
	r1, r2 := Discriminant(1, 3, 1)
	IsEqual(r1, r2, t)
}