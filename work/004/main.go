package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func Solve(a [][]int) [][]int {
	ret := make([][]int, len(a))
	for i := range ret {
		ret[i] = make([]int, len(a[i]))
	}
	hsum, wsum := make([]int, len(a[0])), make([]int, len(a))
	for i, ai := range a {
		for j, aij := range ai {
			hsum[j] += aij
			wsum[i] += aij
		}
	}
	// log.Println("Solve: hsum[", hsum, "], wsum[", wsum, "]")
	for i, ai := range a {
		for j, aij := range ai {
			ret[i][j] = hsum[j] + wsum[i] - aij
		}
	}
	return ret
}

func ReadInputFromFile(filename string) [][]int {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	firstline := strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(firstline[0])
	w, _ := strconv.Atoi(firstline[1])
	// Initialize
	matrix := make([][]int, h)
	for i := range matrix {
		matrix[i] = make([]int, w)
	}
	count := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for i := range line {
			matrix[count][i], _ = strconv.Atoi(line[i])
		}
		count++
	}
	return matrix
}

func ReadOutputFromFile(filename string) [][]int {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	// Initialize
	var matrix [][]int
	count := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var num int
		var buf []int
		for i := range line {
			num, _ = strconv.Atoi(line[i])
			buf = append(buf, num)
		}
		matrix = append(matrix, buf)
		count++
	}
	return matrix
}
