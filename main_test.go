package main

import (
	"fmt"
	"testing"
)

func TestMTest(t *testing.T) {
	got := GitPullBlogiTest()
	expect := "git pull ok!"
	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}

}

func BenchmarkMTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GitPullBlogiTest()
	}
}

//必须填写Example+函数名，否者无法识别
func ExampleGitPullBlogiTest() {
	h1 := GitPullBlogiTest()
	fmt.Println(h1)
}
