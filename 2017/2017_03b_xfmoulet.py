
SZ = 20

matrix  = [[0]*SZ for _ in range(SZ)] # large one

def fill(n) : 
	'get x,y position of data where 1 is 0,0'
	n -= 1
	dirs = [(1,0),(0,1),(-1,0),(0,-1)]
	# go 1 step right, 1 step up, 2 steps left, 2 steps down, 3 steps right, ...
	x=y=0
	nsteps = 1
	branch=0
	matrix[SZ//2][SZ//2] = 1

	for i in range(n) :
		nsteps -= 1
		x+=dirs[branch%4][0]
		y+=dirs[branch%4][1]
		print(i,'pos',x,y)

		if nsteps == 0  : 
			branch +=1
			nsteps = branch/2+1

		number = sum(matrix[SZ/2+y+dy][SZ/2+x+dx] for dx,dy in [(-1,-1),(-1,0),(-1,1),(0,1),(0,-1),(1,-1),(1,0),(1,1)])

		matrix[SZ//2+y][SZ//2+x] = number

fill(100)
for l in matrix:  
	print(l)

