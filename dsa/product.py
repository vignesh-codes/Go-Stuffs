nums = [1,2,3,4,5]
#[6,3,2]

larr = []
rarr = []
oarr = []
l = len(nums)
left = list[1,1,1,1,1]
right = list[1,1,1,1,1]


for i in range(1,len(nums)):
    left = left[i-1]*nums[i-1]
    larr.append(left)
    
for j in range(0,l-1):
    right = right[j+1]*nums[j+1]
    rarr.append(right)
 
larr.insert(0,0)
rarr.reverse()
rarr.append(0)
oarr = [sum(x) for x in zip(larr,rarr)]
print(larr, rarr, oarr)