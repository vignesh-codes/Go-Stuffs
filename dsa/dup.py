def removeDuplicates(nums):
    
    
      
    nums.sort()
    for i in range(2,len(nums)-1):
        if nums[i] == nums[i-1]:
            nums.pop(i)
    print (nums, len(nums))  
nums=[0,0,1,1,1,2,2,3,3,4,5]  
removeDuplicates(nums)