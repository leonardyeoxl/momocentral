package main

// Time Complexity: O(1)
// Space Complexity: O(1)
func TestArray(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	
	if arr[len(arr)-1] > 10 {
		return arr[len(arr)-1]
	}
	return -1
}

func main() {
	test01 := []int{7,8,9,10,11}
	ans01 := TestArray(test01)
	if ans01 != 11 {
		panic("Wrong Answer for [7,8,9,10,11]")
	}

	test02 := []int{1,2,3}
	ans02 := TestArray(test02)
	if ans02 != -1 {
		panic("Wrong Answer for [1,2,3]")
	}

	test03 := []int{}
	ans03 := TestArray(test03)
	if ans03 != -1 {
		panic("Wrong Answer for []")
	}
}