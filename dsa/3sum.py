nums = [1,2,23,3,4,5,-4,-1,4,2,5,-5,-1,0,1]
res = []
def threeSum(nums):
    
    if len(nums) < 3:
        return []
    for i in range(len(nums)-2):
        if i>0 and nums[i] == nums[i-1]:
            continue
        L = i+1
        R = len(nums)-1
        while L<R:
            sums = nums[i]+nums[L]+nums[R]
            if sums>0:
                R = R-1
                # print(nums[i],nums[L],nums[R])
            elif sums<0:
                L = L+1
                #res.append([nums[i],nums[L],nums[R]])
                # print(res)
                # print(nums[i],nums[L],nums[R])
            else:
                res.append([nums[i],nums[L],nums[R]])
                
                while L<R and nums[i]==nums[i+1]:
                    L = L+1
                while L<R and nums[i]==nums[i-1]:
                    R = R-1
                
                L = L+1
                R = R-1
    print(res)
    return res
                    

threeSum(nums)