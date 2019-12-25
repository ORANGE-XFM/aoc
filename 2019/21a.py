laby = [l for l in open('21.txt')]
for l  in laby : print(l[:-1])
# find transport pods
W = len(laby[0])
H = len(laby)

portals = {}
portal_pos = {}

# verticals
for x in range(W): 
	for y,dpos in [(0,2),(37,-1),(96,2),(133,-1)] :
	# for y,dpos in [(0,2),(7,-1),(10,2),(17,-1)] :
		print(x,y)
		if laby[y][x].isupper() : 
			k=laby[y][x]+laby[y+1][x]
			p=(x,y+dpos)
			portals.setdefault(k,[]).append(p)
# horizontals
for y in range(H): 
	for x,dpos in [(0,2),(37,-1),(90,2),(127,-1)] :
	# for x,dpos in [(0,2),(7,-1),(12,2),(19,-1)] :
		if laby[y][x].isupper() : 
			k=laby[y][x]+laby[y][x+1]
			p=(x+dpos,y)
			portals.setdefault(k,[]).append(p)


inv_portals = {v[0]:k for k,v in portals.items()} 
inv_portals.update( {v[1]:k for k,v in portals.items() if len(v)==2} )

for k,v in portals.items(): 
	print (k,v)
	if len(v)==2 : 
		a,b=v
		portal_pos[a]=b
		portal_pos[b]=a
for i in portal_pos.items(): print (i)


# part A
# return accessible keys from there + distance
visited = {}  # pos:dist
start = portals['AA'][0]
frontier = [(start[0],start[1],0)]

while frontier : 
	x,y,d  = frontier.pop(0)
	visited[x,y] = d
	for nx,ny in [(x+1,y),(x-1,y),(x,y+1),(x,y-1)] :
		nd=d+1
		if (nx,ny) in portal_pos : 			
			nx,ny=portal_pos[(nx,ny)]
			nd+=1
			print('zap! from',portal_pos[(nx,ny)],'to',inv_portals[(nx,ny)],nd,'dist')

		if laby[ny][nx]=='.' :
			if (nx,ny) in visited : 
				if nd<visited.get((nx,ny),9999999): 
					frontier.append((nx,ny,nd))
					print('taken, better ')
				# else : 
				# 	print(nx,ny,'not taken, already best',nd,'>=',visited[nx,ny])
			else :
				frontier.append((nx,ny,nd))
		if (nx,ny)==portals['ZZ'][0]:
			print("-----",nx,ny,nd)


# for y,_ in enumerate(laby):
#   for x,_ in enumerate(laby[0]):
#   	print(f'{visited.get((x,y),"xx"):2}',end=' ')
#   print()
