package main

import (
	"reflect"
)	

// Time Complexity: O(N*M). Given N is the number of rows of the array and M is the number of elements in the row.
// Space Complexity: O(N) assuming that we take into account the answer to be returned.
func Summarize(arr [][]int) []int {
	count := 0
	answer := []int{}

	for row_index, _ := range arr {
		count = 0
		
		for col_index, _ := range arr[row_index] {
			if arr[row_index][col_index] <= 0 {
				count = 0
				break
			}
			count = count + arr[row_index][col_index]
		}

		if count > 0 {
			answer = append(answer, count)
		}
	}

	return answer
}

func main() {
	test01 := [][]int{
		{1,2,3,4},
		{2,3,4,5},
		{5,6,-3},
	}
	ans01 := Summarize(test01)
	if reflect.DeepEqual(ans01, []int{10,14}) == false {
		panic("Wrong Answer for [[1,2,3,4],[2,3,4,5],[5,6,-3]]")
	}

	test02 := [][]int{
		{1,2,3,4},
		{2,3,4,5},
		{5,6,3},
	}
	ans02 := Summarize(test02)
	if reflect.DeepEqual(ans02, []int{10,14,14}) == false {
		panic("Wrong Answer for [[1,2,3,4],[2,3,4,5],[5,6,3]]")
	}

	test03 := [][]int{
		{1,2,3,4},
		{2,3,4,5},
		{},
	}
	ans03 := Summarize(test03)
	if reflect.DeepEqual(ans03, []int{10,14}) == false {
		panic("Wrong Answer for [[1,2,3,4],[2,3,4,5],[]]")
	}
}