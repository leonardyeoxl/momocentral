def checkVertical(arr, row_index, col_index, max_rows, VALID_COUNT):
    index = row_index
    num = arr[row_index][col_index]
    count = 0

    if row_index > 0:
        start_row_index = row_index-1
    elif row_index == 0:
        start_row_index = 0
    
    if row_index == 0:
        while index <= max_rows-1:
            if num == arr[index][col_index]:
                count += 1
            else:
                break
            index += 1
    elif start_row_index >= 0 and num != arr[start_row_index][col_index]:
        while index <= max_rows-1:
            if num == arr[index][col_index]:
                count += 1
            else:
                break
            index += 1

    if count >= VALID_COUNT:
        return True
    return False

def checkHorizontal(arr, row_index, col_index, max_cols, VALID_COUNT):
    index = col_index
    num = arr[row_index][col_index]
    count = 0

    if col_index > 0:
        start_col_index = col_index-1
    elif col_index == 0:
        start_col_index = 0

    if col_index == 0:
        while index <= max_cols-1:
            if num == arr[row_index][index]:
                count += 1
            else:
                break
            index += 1
    elif start_col_index >= 0 and num != arr[row_index][start_col_index]:
        while index <= max_cols-1:
            if num == arr[row_index][index]:
                count += 1
            else:
                break
            index += 1

    if count >= VALID_COUNT:
        return True
    return False

# Time Complexity: O(N*M*(N-1+M-1)). Given N is the number of elements in a row and M is the number of elements in a column.
# Space Complexity: O(1). No auxiliary memory used.
def count_chains(arr):
    max_rows = len(arr)
    max_cols = len(arr[0])
    VALID_COUNT = 3

    count = 0
    for row_index in range(len(arr)):
        for col_index in range(len(arr[row_index])):
            valid_horizontal = checkHorizontal(arr, row_index, col_index, max_cols, VALID_COUNT)
            valid_vertical = checkVertical(arr, row_index, col_index, max_rows, VALID_COUNT)
            if valid_horizontal and not valid_vertical:
                count += 1
            elif valid_vertical and not valid_horizontal:
                count += 1
            elif valid_vertical and valid_horizontal:
                count += 2

    return count

assert count_chains([[0,3,3],[0,0,0],[0,1,4]]) == 2, "Wrong answer for [[0,3,3],[0,0,0],[0,1,4]]"
assert count_chains([[3,3,3,4,1,2],[0,0,3,2,4,4],[4,0,0,2,2,1],[1,3,4,2,4,0],[0,1,0,2,0,1],[2,2,2,4,1,3]]) == 3, "Wrong answer for [[3,3,3,4,1,2],[0,0,3,2,4,4],[4,0,0,2,2,1],[1,3,4,2,4,0],[0,1,0,2,0,1],[2,2,2,4,1,3]]"
assert count_chains([[3,3,3,3,3],[0,0,1,2,1],[0,1,4,3,2]]) == 1, "Wrong answer for [[3,3,3,3,3],[0,0,1,2,1],[0,1,4,3,2]]"
