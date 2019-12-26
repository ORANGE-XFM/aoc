import itertools

def reachable(board,x,y) : 
	# return accessible keys from there + distance
	keys = []
	visited = []
	frontier = [(x,y,0,'')] # x,y dist doors

	while frontier : 
		x,y,d,doors  = frontier.pop(0)
		visited.append((x,y))
		for nx,ny in [(x+1,y),(x-1,y),(x,y+1),(x,y-1)] :
			if (nx,ny) not in visited :
				visited.append((nx,ny))
				c = board[ny][nx]
				if c!='#' :
					ndoors = doors+c.lower() if c.isupper() else doors
					frontier.append((nx,ny,d+1,ndoors))
				if c.islower() : 
					keys.append((c,d+1,doors))
	return keys


# returns a list of key positions
def where_keys(board): 
	d = []
	for y,l in enumerate(board) : 
		for x,c in enumerate(l) : 
			if c.islower() or c=='@' : 
				d.append((x,y)) 
	return d

#load board, find keys on board
start_board = [ l.strip() for l in open('18.txt') ]
for l in start_board : print (l)
allkeys = where_keys(start_board)
reach = { start_board[y][x] : reachable(start_board,x,y) for x,y in allkeys } # (lettera,letterB) -> dist, doors

for i in reach.items() : print (i)

nb = len(allkeys)
dmin = 9999999999999
paths=[('@',0)] # keys so far / path, dist
while paths : 	
	path,d = paths.pop(0)
	print(path,d)
	if len(path) == nb : 
		print(path,d)
		if d<dmin : 
			dmin = d
			print(path,d,'*'*80)
	nextkeys = sorted(reach[path[-1]],key=lambda x:x[1])
	for nkey,dist,doors in nextkeys : # nearest first
		if nkey not in path and all(k in path for k in doors) : 
			if d+dist < dmin :
				paths.append((path+nkey,d+dist))
