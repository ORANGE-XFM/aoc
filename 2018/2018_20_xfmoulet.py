def disp_grid(allpos) : 
	# display grid
	SZ = 20
	grid = [['.' for _ in range(SZ)] for _ in range(SZ)]
	grid[SZ//2][SZ//2]='O'
	for p in allpos : 
		grid[int(p.imag+SZ/2)][int(p.real+SZ/2)]='x'

	for y in range(SZ) : 
		print (' '.join(grid[y]))

def parse(regex) :
	stack = [] # stacks : position, distance

	pos =0+0j # complex number for posisiton
	dist=0
	allpos={}

	maxdists = 0

	for c in regex[1:-1] :
		if c=='(' : 
			stack.append((pos,dist))
		elif c=='|' : 
			pos,dist = stack[-1]
		elif c==')' : 
			stack.pop()
		else : 
			pos += {'N':-1j, 'S':1j, 'E':1, 'W':-1}[c]
			dist +=1
			if pos in allpos : 
				#print ("-------  ",pos,allpos[pos],dist)
				dist = allpos[pos] # better known distance before
			allpos[pos] = dist
			#disp_grid(allpos.keys())
		#print(c,stack,pos,dist)
		maxdists = max(dist,maxdists)

	return maxdists, sum(1 for _,d in allpos.items() if d>=1000 )


for s in [
	#'^N(E|W|S|E)S$',
	#'^WNE$',
	#'^ENWWW(NEEE|SSE(EE|N))$',
	'^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$',
	'^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$',
	'^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$',

	open('2018_20.data').read()
	] : 

	print(s,parse(s))
