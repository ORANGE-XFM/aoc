import re
from heapq import heappop, heappush

bots = [] # x,y,z,R
for l in open('2018_23.txt') : 
	bots.append(tuple(int(x) for x in re.match(r"pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)",l).groups()))

def dist(a,b) : 
	return sum(abs(a[i]-b[i]) for i in (0,1,2))

def intersect(b1,b2) : return dist(b1,b2)<= b1[3]+b2[3]
def nb_intersect(b) : return -sum(1 for b2 in bots if intersect(b,b2))

biggest = max(bots, key=lambda x:x[3])
nb_within = sum(1 for bot in bots if dist(bot,biggest)<=biggest[3])
print ("biggest ",biggest,"nb within:",nb_within)

init = (0,0,0,100000000)
heap=[(nb_intersect(init),0,init)]

"""take 6 smaller shperes and keep the ones that intersect with the most bots. 
if ties, take the closest from center 
stop when sphere is of radius 1
"""

while True :
	n,d,(x,y,z,r) = heappop(heap) # pop
	print ((x,y,z,r),"intersects with",-n,"dist from center",d)
	if r<=1 : break
	r2 = r//2
	# take 6 subcubes, find best (ie the one intersecting with the most cubes)
	for sub in (x-r2,y,z,r2),(x+r2,y,z,r2),(x,y-r2,z,r2),(x,y+r2,z,r2),(x,y,z-r2,r2),(x,y,z+r2,r2) : 
		nb = nb_intersect(sub)
		if nb : heappush(heap,(nb,dist(init,sub),sub))
