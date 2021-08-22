# Time Complexity: O(N*M). Given N is the number of rows of the array and M is the number of elements in the row.
# Space Complexity: O(N) assuming that we take into account the answer to be returned.
def Summarize(arr):
    answer = []
    
    for row_index in range(len(arr)):
        count = 0

        for col_index in range(len(arr[row_index])):
            if arr[row_index][col_index] <= 0:
                count = 0
                break
            count += arr[row_index][col_index]
        
        if count > 0:
            answer.append(count)

    return answer

assert Summarize([[1,2,3,4],[2,3,4,5],[5,6,-3]]) == [10,14], "Wrong Answer for [[1,2,3,4],[2,3,4,5],[5,6,-3]]"
assert Summarize([[1,2,3,4],[2,3,4,5],[5,6,3]]) == [10,14,14], "Wrong Answer for [[1,2,3,4],[2,3,4,5],[5,6,3]]"
assert Summarize([[1,2,3,4],[2,3,4,5],[]]) == [10,14], "Wrong Answer for [[1,2,3,4],[2,3,4,5],[]]"