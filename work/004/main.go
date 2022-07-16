package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func Solve() {

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
