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
	"random01.txt",
	"random02.txt",
	"random03.txt",
	"random04.txt",
	"random05.txt",
	"random06.txt",
	"random07.txt",
	"random08.txt",
	"hand01.txt",
	"hand02.txt",
	"hand03.txt",
	"hand04.txt",
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
	for idx, f := range filenames {
		a := ReadInputFromFile(InputPath + f)
		got := Solve(a)
		want := ReadOutputFromFile(OutputPath + f)
		if !(reflect.DeepEqual(got, want)) {
			t.Errorf("Case(TestSolve:%d:%q) got %v want %v", idx, f, got, want)
		}
	}
}
