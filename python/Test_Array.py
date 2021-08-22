# Time Complexity: O(1)
# Space Complexity: O(1)
def TestArray(arr):
    if len(arr) == 0:
        return False
    
    if arr[-1] > 10:
        return arr[-1]
    else:
        return False

assert TestArray([7,8,9,10,11]) == 11, "Wrong Answer for [7,8,9,10,11]"
assert TestArray([1,2,3]) == False, "Wrong Answer for [1,2,3]"
assert TestArray([]) == False, "Wrong Answer for []"