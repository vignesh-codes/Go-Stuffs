"""Find the Mth element of the Array after K left rotations"""

def getFirstElement(a, N, K, M):  
      
    # The array comes to original state  
    # after N rotations  
    K %= N  
  
    # Mth element after k left rotations  
    # is (K+M-1)%N th element of the  
    # original array  
    index = (K + M - 1) % N  
  
    result = a[index]  
  
    # Return the result  
    return result  

a = [ 3, 4, 5,6,11,22,33,55,12,23 ]  
  
    # Size of the array  
N = len(a)  
  
    # Given K rotation and Mth element  
    # to be found after K rotation  
K = 3
M = 1
  
    # Function call  
print(getFirstElement(a, N, K, M))  

# Function to return Mth element of 
# array after k right rotations 
def getFirstElement(a, N, K, M): 
  
    # The array comes to original state 
    # after N rotations 
    K %= N 
  
    # If K is greater or equal to M 
    if (K >= M): 
  
        # Mth element after k right 
        # rotations is (N-K)+(M-1) th 
        # element of the array 
        index = (N - K) + (M - 1) 
  
    # Otherwise 
    else: 
  
        # (M - K - 1) th element 
        # of the array 
        index = (M - K - 1) 
  
    result = a[index] 
  
    # Return the result 
    return result 

      
aa = [ 1, 2, 3, 4, 5 ] 
NN = len(aa) 
  
KK , MM = 3, 2
  
print( getFirstElement(aa, NN, KK, MM)) 
