package main

import (
	"fmt"
	"testing"
)

//func Testsdd(t *testing.T){
//	sum := sdd()
//	if sum != 3{
//		t.Errorf("add error")
//	}
//}

func TestMytest(t *testing.T) {
	got := Mytest("test")
	expect := "handle request ok"
	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}

func BenchmarkMytest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mytest("all")
	}
}

func ExampleMytest() {
	h1 := Mytest("all")
	fmt.Println(h1)
}
