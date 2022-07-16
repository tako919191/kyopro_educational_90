package main

import (
	"reflect"
	"testing"
)

const (
	InputPath  string = "testcase/in/"
	OutputPath string = "testcase/out/"
)

var filenames = []string{
	"sample_01.txt",
	"sample_02.txt",
	"sample_03.txt",
	"sample_04.txt",
}

func TestReadInputFromFile(t *testing.T) {
	got := ReadInputFromFile(InputPath + filenames[0])
	want := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	if !(reflect.DeepEqual(got, want)) {
		t.Errorf("Case(TestReadInputFromFile) got %v want %v", got, want)
	}
}

func TestReadOutputFromFile(t *testing.T) {
	got := ReadOutputFromFile(OutputPath + filenames[0])
	want := [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}
	if !(reflect.DeepEqual(got, want)) {
		t.Errorf("Case(TestReadOutputFromFile) got %v want %v", got, want)
	}
}

func TestSolve(t *testing.T) {
	// for idx, f := range filenames {
	// 	a := ReadInputFromFile(f[idx])
	// 	got := Solve(a)
	// 	want := ReadOutputFromFile(f[idx])
	// 	if got != want {
	// 		t.Errorf("Case(%d) got %q want %q", idx, got, want)
	// 	}
	// }
}
