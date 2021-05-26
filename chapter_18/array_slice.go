package main

import "fmt"

func main() {
	// （1）如何截断数组或者切片的最后一个元素：
	line := [6]int{0,1,2,3,4,5}
	s1 := line[:len(line)-1]
	fmt.Printf("slice line: %v\n",s1)

	// （2）如何使用for或者for-range遍历一个数组（或者切片）：
	for i := 0; i < len(line); i++ {
		fmt.Printf("array for : %d -- %d\n", i, line[i])
	}

	for i, v := range s1 {
		fmt.Printf("slice for-range: %d -- %d\n", i, v)
	}

	arr2Dim := [][]int{{1},{2},{3},{4},{5}}
	found := false
	Found: for row := range arr2Dim {
		for column := range arr2Dim[row] {
			if arr2Dim[row][column] == 6 {
				found = true
				break Found
			}
		}
	}
	fmt.Printf("found: %v\n", found)
}