# Python3program for rearrange an  
# array such that arr[i] = i. 
  
  
# Function to rearrange an array 
# such that arr[i] = i. 

arr = [-1, -1, 6, 1, 9, 
        3, 2, -1, 4,-1] 

def fix(arr):
    s = set()
    for i in range(len(arr)):
        s.add(arr[i])
    
    
    
    for i in range(len(arr)):
        
        if i in s:
            arr[i] = i
        else:
            arr[i] = -1
    return arr    


print(fix(arr))




# Iterative python program to reverse an array
 
# Function to reverse A[] from start to end
def reverseList(A, start, end):
    while start < end:
        A[start], A[end] = A[end], A[start]
        start += 1
        end -= 1
 
# Driver function to test above function
A = [1, 2, 3, 4, 5, 6]
print(A)
reverseList(A, 0, 5)
print("Reversed list is")
print(A)



# Recursive python program to reverse an array
 
# Function to reverse A[] from start to end
def reverseList(A, start, end):
    if start >= end:
        return
    A[start], A[end] = A[end], A[start]
    reverseList(A, start+1, end-1)
 
# Driver function to test above function
A = [1, 2, 3, 4, 5, 6]
print(A)
reverseList(A, 0, 5)
print("Recursively Reversed list is")
print(A)


# Python3 code to move all zeroes 
# at the end of array 
  
# Function which pushes all 
# zeros to end of an array. 
def pushZerosToEnd(arr, n): 
    count = 0 # Count of non-zero elements 
      
    
    for i in range(n): 
        if arr[i] != 0: 
            arr[count] = arr[i] 
            count+=1
        print (arr)
    while count < n: 
        print (arr[count])
        arr[count] = 0
        print ("n",arr[count])
        count += 1
          
# Driver code 
arr = [1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9] 
n = len(arr) 
pushZerosToEnd(arr, n) 
print("Array after pushing all zeros to end of array:") 
print(arr) 
from array import array
def nn(arr,n):
    n=[]
    b=[]
    for i in (arr):
        if i ==0:
            n.append(0)
        else:
            b.append(i)
    print(b+n)
arr = [1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9,4,5,6,0]

n = len(arr) 
nn(arr, n)   
