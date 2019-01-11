from math import sqrt
x=10551275
div=[]
for i in range(1,int(sqrt(x))) : 
   if x//i*i==x : div += [i,x//i]
print(div)
print(sum(div))

# or smaller print(sum( i+x//i for i in range(1,int(sqrt(x))) if x//i*i==x ))