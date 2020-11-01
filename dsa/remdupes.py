arr = [1,2,2,3,4,1,2,2,3,4,2,2,6,5,4,7,8,8,7]
a = arr
arr.sort()
i = 1
while i< len(arr):
    print(arr[i-1])
    if arr[i]==arr[i-1]:
        arr.pop(i)
    else:
        i = i+1
print(arr)