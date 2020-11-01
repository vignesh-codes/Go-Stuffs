arr = [1,2,24,-40,6,6,5,-30,4,4,4,5,10,20]

curSum = 0
maxSum = arr[0]
start = 0
end = 0
temp = 0

for i in range(len(arr)):
    curSum = curSum + arr[i]
    if maxSum<curSum:
        maxSum = curSum
        start = temp
        end = i
        
    if curSum<0:
        curSum = 0
        temp = i+1
print(maxSum, arr[start:end+1])