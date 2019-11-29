import re
r = re.compile(r'^position=<\s*([-0-9]+),\s*([-0-9]+)> velocity=<\s*([-0-9]+),\s*([-0-9]+)>$')

state = [[int(x) for x in r.match(l).groups()] for l in open('2018_10.data')]

def size(state) : 
	dx = max(s[0] for s in state)-min(s[0] for s in state)
	dy = max(s[1] for s in state)-min(s[1] for s in state)
	return dx*dy

def show(state) : 
	maxx = max(s[0] for s in state)
	maxy = max(s[1] for s in state)

	mat = [[ ' ' for _ in range(maxx+1) ] for _ in range(maxy+1)]

	for s in state : 
		mat[s[1]][s[0]]='+'

	print('showing',maxx,maxy,size(state))
	for l in mat : 
		print (''.join(l))
	print()

def evolve(state) : 
	for s in state : 
		s[0] += s[2]
		s[1] += s[3]


prevm = None

for i in range(400000) : 
	prevm=[l[:] for l in state]
	evolve(state) # test if a certain number are aligned ?
	sz = size(state)
	if prevm : print (sz,size(prevm))
	if prevm and sz>size(prevm) : 
		show(prevm)
		print(i)
		break

print('evolved')
#show(state)
