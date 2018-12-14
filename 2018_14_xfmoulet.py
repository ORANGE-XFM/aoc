N=765071

def reset() : 
	global recipes, idx1,idx2
	recipes = [3,7]
	idx1=0
	idx2=1

def step() : 
	global recipes,idx1,idx2
	n = recipes[idx1]+recipes[idx2]
	recipes += [ int(digit) for digit in str(n) ]
	sz = len(recipes)
	idx1 = (idx1+1+recipes[idx1])%sz
	idx2 = (idx2+1+recipes[idx2])%sz

# puzzle A
reset()
while len(recipes)<N+11:
	step()
print(''.join(str(x) for x in recipes[N:N+10]))
	
# puzzle B
reset()
lM = [int(x) for x in str(N)]

i=0
while recipes[-len(lM):] != lM and recipes[-len(lM)-1:-1] != lM : # only check end of recipe, not all
	i += 1 
	if i%1000000==0 : print(i)
	step()
	#print (idx1,idx2,recipes)
print( len(recipes)-len(lM) if recipes[-len(lM):] == lM else len(recipes)-len(lM)-1)


