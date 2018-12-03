import re
p = re.compile (r'^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$')

lines = []
for l in open('2018_03.data') : 
	match = p.match(l)
	if match is not None : 
		lines.append([int(x) for x in match.groups()])

matrix = [[0]*1000 for _ in range(1000)]
for id,x,y,w,h in lines : 
	for iy in range(y,y+h) : 
		for ix in range(x,x+w) : 
			if matrix[iy][ix]==0 :
				matrix[iy][ix] = id 
			else :
				matrix[iy][ix] = 'x'
print('sq in',sum(sum(1 for x in l if x=='x') for l in matrix))

# check
for id,x,y,w,h in lines : 
	ok = True
	for iy in range(y,y+h) : 
		for ix in range(x,x+w) : 
			if matrix[iy][ix]=='x' :
				ok=False
	if ok : print ('ok piece:',id)


