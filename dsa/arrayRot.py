"""1. Write a function rotate(ar[], d, n) that rotates arr[] of size n by d elements. O(n) and O(n)"""
arr = [1,2,3,4,5,6,7]
d = 3                                                       
n = len(arr)
def rot(arr, d, n):
    temp = arr[:d]
    arr = arr[d:]
    arr.append(temp)
    print ("Simple Rot: ",arr)
    return arr

rot(arr,d,n)

"""2. Rotote one by one O(nd) and O(1)"""

def rotl(arr,d,n):
    for i in range(d):
        leftRot(arr,n)
    return arr

def leftRot(arr,n):
    temp = arr[0]
    for i in range(n-1):
        arr[i]= arr[i+1]
    arr[n-1]= temp
    # print (arr)


def printArray(arr, size): 
    print("Rotate One by One")
    for i in range(size): 
        print ("% d"% arr[i], end =",")
        # print ("% d"% arr[i], end =" ")


rotl(arr,d,n)
printArray(arr, n)

"""3. Reversal Algorithm for Rotation"""
def rverseArray(arr, start, end): 
    while (start < end): 
        temp = arr[start] 
        arr[start] = arr[end] 
        arr[end] = temp 
        start += 1
        end = end-1
  
# Function to left rotate arr[] of size n by d 
def leftRotate(arr, d): 
  
    if d == 0: 
        return
    n = len(arr) 
    rverseArray(arr, 0, d-1) 
    rverseArray(arr, d, n-1) 
    rverseArray(arr, 0, n-1) 
  
# Function to print an array 
def printArray(arr): 
    print ("\nWith Reversal Algorithm:")
    for i in range(0, len(arr)):
        print(arr[i], end=" ")
  
# Driver function to test above functions 
arr = [1, 2, 3, 4, 5, 6, 7] 
n = len(arr) 
d = 2
  
# in case the rotating factor is  
# greater than array length 
d = d % n       
leftRotate(arr, d) # Rotate array by 2 
printArray(arr) 
    

"""Given an array, cyclically rotate the array clockwise by one
1) Store last element in a variable say x.
2) Shift all elements one position ahead.
3) Replace first element of array with x."""

def crs(arr,d,n):
    for i in range(4):
        cr(arr,n)
def cr(arr,n):
    x = arr[n-1]
    for i in range(n-1,0,-1):
        arr[i]=arr[i-1]
    arr[0]=x
    k = arr
    
    print("\n With Cyclic CW Rotation \n",k)
    return k




crs(arr,d,n)