package main

func checkVertical(arr [][]int, row_index int, col_index int, max_rows int, VALIDCOUNT int) bool {
	index := row_index
    num := arr[row_index][col_index]
    count := 0
	start_row_index := 0

    if row_index > 0 {
		start_row_index = row_index-1
	} else if row_index == 0 {
		start_row_index = 0
	}

    if row_index == 0 || (start_row_index >= 0 && num != arr[start_row_index][col_index]) {
		for {
			if index > max_rows-1 {
				break
			}
			if num == arr[index][col_index] {
				count += 1
			} else {
				break
			}
			index = index + 1
		}
	}
        

    if count >= VALIDCOUNT {
		return true
	}
    return false
}

func checkHorizontal(arr [][]int, row_index int, col_index int, max_cols int, VALIDCOUNT int) bool {
	index := col_index
    num := arr[row_index][col_index]
    count := 0
	start_col_index := 0

    if col_index > 0 {
		start_col_index = col_index-1
	} else if col_index == 0 {
		start_col_index = 0
	}

    if col_index == 0 || (start_col_index >= 0 && num != arr[row_index][start_col_index]) {
		for {
			if index > max_cols-1 {
				break
			}
			if num == arr[row_index][index] {
				count += 1
			} else {
				break
			}
			index = index + 1
		}
	}
        

    if count >= VALIDCOUNT {
		return true
	}
    return false
}

// Time Complexity: O(N*M*(N-1+M-1)). Given N is the number of rows of the array and M is the number of elements in the row.
// Space Complexity: O(1)
func count_chains(arr [][]int) int {
	maxRows := len(arr)
	maxCols := len(arr[0])
	VALIDCOUNT := 3
	validHorizontal := false
	validVertical := false

	count := 0
	for rowIndex, _ := range arr {
		for colIndex, _ := range arr[rowIndex] {
			validHorizontal = checkHorizontal(arr, rowIndex, colIndex, maxCols, VALIDCOUNT)
            validVertical = checkVertical(arr, rowIndex, colIndex, maxRows, VALIDCOUNT)
            if validHorizontal == true && validVertical == false{
				count = count + 1
			} else if validVertical == true && validHorizontal == false {
				count = count + 1
			} else if validVertical == true && validHorizontal == true {
				count = count + 2
			}
		}
	}

	return count
}

func main() {
	test01 := [][]int{
		{0,3,3,},
		{0,0,0,},
		{0,1,4,},
	}
	ans01 := count_chains(test01)
	if ans01 != 2 {
		panic("Wrong answer for [[0,3,3],[0,0,0],[0,1,4]]")
	}

	test02 := [][]int{
		{3,3,3,4,1,2,},
		{0,0,3,2,4,4,},
		{4,0,0,2,2,1,},
		{1,3,4,2,4,0,},
		{0,1,0,2,0,1,},
		{2,2,2,4,1,3,},
	}
	ans02 := count_chains(test02)
	if ans02 != 3 {
		panic("Wrong answer for [[3,3,3,4,1,2],[0,0,3,2,4,4],[4,0,0,2,2,1],[1,3,4,2,4,0],[0,1,0,2,0,1],[2,2,2,4,1,3]]")
	}

	test03 := [][]int{
		{3,3,3,3,3,},
		{0,0,1,2,1,},
		{0,1,4,3,2,},
	}
	ans03 := count_chains(test03)
	if ans03 != 1 {
		panic("Wrong answer for [[3,3,3,3,3],[0,0,1,2,1],[0,1,4,3,2]]")
	}
}