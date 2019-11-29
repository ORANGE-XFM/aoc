
coords = []
for l in open('2018_06.data') : 
	x,y = (int(x.strip()) for x in l.split(','))
	coords.append((x,y))
print (coords)


# find size
maxx = max(x for x,y in coords)+1
maxy = max(y for x,y in coords)+1

# create matrix
matrix = [[0]*maxx for _ in range(maxy)]

def manh(a,b) : return abs(a[0]-b[0])+abs(a[1]-b[1])

# part A
for y in range(maxy) :
	for x in range(maxx) :
		ids = [manh(c,(x,y)) for c in coords]
		id = min(ids)
		if sum(1 for x in ids if id==x)!=1 :
			id = -1
		else :
			id = ids.index(id)
		matrix[y][x] = id
for l in matrix: print (l)

borders = set(matrix[0]+matrix[-1]+[l[0] for l in matrix]+[l[-1] for l in matrix])
print (borders)

def size(x) : 
	return sum(sum(1 for n in l if n==x) for l in matrix)

nb = max(size(x) for x in range(len(coords)) if x not in borders)
print(nb)

# part B 
print('-'*80)
SZ = 10000
n = 0 
for y in range(maxy) :
	for x in range(maxx) :
		dist = sum(manh(c,(x,y)) for c in coords)
		if dist < SZ : n += 1 
		print (dist, end=',')
	print()


print (n)
