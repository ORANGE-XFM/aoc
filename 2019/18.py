import functools
start_board = [ l.strip() for l in open('18.txt') ]

# remove char in a board
def pprint(board,x,y) : 
	for l in replace(board,x,y,'@') : print(l)
def remove(c,board) : 
	return tuple( l.replace(c,'.') for l in board)
def use_key(c,board) : 
	return remove(c.upper(),remove(c,board))

def replace(board,x,y,c) : 
	return tuple(l if i != y else l[:x]+c+l[x+1:] for i,l in enumerate(board))

@functools.lru_cache(maxsize=None)
def find_keys(board,x,y) : 
	# return accessible keys from there + distance
	keys = []
	visited = []
	dist = []	# correspond to preceding
	frontier = [(x,y,0)]

	def check(x,y) : return board[y][x]=='.' and (x,y) not in visited

	while frontier : 
		x,y,d  = frontier.pop(0)
		visited.append((x,y))
		for nx,ny in [(x+1,y),(x-1,y),(x,y+1),(x,y-1)] :
			if check(nx,ny) : 
				frontier.append((nx,ny,d+1))
			elif board[ny][nx].islower() : 
				keys.append((nx,ny,d+1))
	return keys

b=start_board
# get starting position
for y,l in enumerate(b) : 
	if '@' in l : 
		x = l.index('@')
		break
b = remove('@',b)

frontier = {(b,x,y)}
costs = {(b,x,y):0}

best=99999999999999999999
i=0
while frontier : 
	i += 1
	if i%100==0 : print('>',i,len(frontier))
	b,x,y = frontier.pop()
	c = costs[(b,x,y)]

	keys = find_keys(b,x,y)
	if not keys :
		remain = len(set(''.join(b)))
		if remain != 2 : print('wat')

		if c<best : 
			best=c
			print('best',c)
	else : 
		for kx,ky,cost in keys : 
			# remove key and door
			nb=use_key(b[ky][kx],b)
			state = (nb,kx,ky) # c+cost
			# check if already seen better, skip
			old_cost = costs.get(state,9999999999)
			if c+cost < old_cost : 
				costs[state] = c+cost
				frontier.add(state)
