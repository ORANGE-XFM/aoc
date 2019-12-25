laby = [l for l in open('21.txt')]
for l  in laby : print(l[:-1])
# find transport pods
W = len(laby[0])
H = len(laby)
print(W,H)
portals = {}
portal_pos = {}

# verticals
for x in range(W): 
	for y,dpos in [(0,2),(37,-1),(96,2),(H-2,-1)] :
	# for y,dpos in [(0,2),(9,-1),(26,2),(H-2,-1)] :
		if laby[y][x].isupper() : 
			k=laby[y][x]+laby[y+1][x]
			p=(x,y+dpos)
			portals.setdefault(k,[]).append(p)
# horizontals
for y in range(H): 
	for x,dpos in [(0,2),(37,-1),(90,2),(W-3,-1)] :
	# for x,dpos in [(0,2),(9,-1),(34,2),(W-3,-1)] :
		if laby[y][x].isupper() : 
			k=laby[y][x]+laby[y][x+1]
			p=(x+dpos,y)
			portals.setdefault(k,[]).append(p)


inv_portals = {v[0]:k for k,v in portals.items()} 
inv_portals.update( {v[1]:k for k,v in portals.items() if len(v)==2} )

for k,v in portals.items(): 
	if len(v)==2 : 
		a,b=v
		if a[0] in (2,W-4) or a[1] in (2,H-3) : 
			dz = -1
		else : 
			dz = 1
		portal_pos[a]=b,dz
		portal_pos[b]=a,-dz

for i in portal_pos.items(): print (i)

DEBUG = False
# return accessible keys from there + distance
visited = {}  # pos:dist
start = portals['AA'][0]
frontier = [(start[0],start[1],0,0)] # xyzd

while frontier : 
	x,y,z,d  = frontier.pop(0)
	visited[x,y,z] = d
	for nx,ny in [(x+1,y),(x-1,y),(x,y+1),(x,y-1)] :
		nd=d+1
		nz=z
		if (nx,ny) in portal_pos : 
			portal_end=portal_pos[(nx,ny)]
			if (portal_end[1]>0 and nz==0) or nz != 0:
				(nx,ny),dz = portal_end
				nd += 1
				nz += dz
				if DEBUG: print('zap! from',portal_pos[(nx,ny)],nd,'dist to',nx,ny,nz)

		if laby[ny][nx]=='.' and abs(nz)<100: # arbitrary max level to avoid getting too low
			if (nx,ny,nz) in visited : 
				if nd<visited.get((nx,ny,nz),9999999): 
					frontier.append((nx,ny,nz,nd))
					if DEBUG: print('taken, better ')
				# else : 
				#  	if DEBUG: print(nx,ny,'not taken, already best',nd,'>=',visited[nx,ny,nz])
			else :
				frontier.append((nx,ny,nz,nd))
		if (nx,ny)==portals['ZZ'][0] and nz==0:
			print("-----",nx,ny,nz,nd)
print('done.')

# for y,_ in enumerate(laby):
#   for x,_ in enumerate(laby[0]):
#   	print(f'{visited.get((x,y),"xx"):2}',end=' ')
#   print()
