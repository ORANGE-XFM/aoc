import re
p = re.compile (r'^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$')

matrix = [[0]*1000 for _ in range(1000)]
for l in open('2018_03.data') : 
	match = p.match(l)
	if match is not None : 
		id,x,y,w,h = [int(x) for x in match.groups()]
		print(l,id,x,y,w,h)
		for iy in range(y,y+h) : 
			for ix in range(x,x+w) : 
				if matrix[iy][ix]==0 :
					matrix[iy][ix]=1 
				else :
					matrix[iy][ix] = 2
					
nb=0

print(sum(sum(1 for x in l if x==2) for l in matrix))

