
def pos(n) : 
	'get x,y position of data where 1 is 0,0'
	n -= 1
	dirs = [(1,0),(0,1),(-1,0),(0,-1)]
	# go 1 step right, 1 step up, 2 steps left, 2 steps down, 3 steps right, ...
	pos = [0,0]
	nsteps = 1
	branch=0

	for i in range(n) :
		nsteps -= 1
		pos[0]+=dirs[branch%4][0]
		pos[1]+=dirs[branch%4][1]
		#print(i,'pos',pos)

		if nsteps == 0  : 
			branch +=1
			nsteps = branch/2+1
			print ('branche',branch,'nsteps', nsteps)
	return pos


assert pos(1) == [0,0]
assert pos(5) == [-1,1]
assert pos(12) == [2,1]
assert pos(23) == [0,-2]
x,y = pos(1024) 
assert abs(x)+abs(y) == 31

x,y = pos(368078)
print (x,y,abs(x)+abs(y))