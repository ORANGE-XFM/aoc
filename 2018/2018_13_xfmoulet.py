from types import SimpleNamespace
from itertools import combinations
class cart (SimpleNamespace):pass # x,y, dir, turn
dirs='^v<>'

def aff():
	mtx=[list(l) for l in tracks]
	for c in carts : 
		if c.dir!='o' : mtx[c.y][c.x]=c.dir
	for l in mtx: print (''.join(l))

def movecart(c) : 
	newdirs = {
		'>\\':'v',	'>/':'^',
		'^\\':'<',	'^/':'>',
		'<\\':'^',	'</':'v',
		'v\\':'>',	'v/':'<',
	}
	# could be simpler with directions as numbers and modulos but this is prettier
	turns = {
		'^':'<^>',
		'<':'v<^',
		'v':'>v<',
		'>':'^>v',
	}
	if   c.dir=='>' : c.x+=1
	elif c.dir=='<' : c.x-=1
	elif c.dir=='^' : c.y-=1
	elif c.dir=='v' : c.y+=1
	pl = tracks[c.y][c.x]
	c.dir = newdirs.get(c.dir+pl, c.dir)
	if pl=='+' : 
		c.dir=turns[c.dir][c.turn]
		c.turn += 1
		c.turn %= 3
	

# load 
tracks = list(x[:-1] for x in open('2018_13.data'))
carts=[]
for y,s in enumerate(tracks) : 
	# find carts	
	for x,c in enumerate(s) : 
		if c in dirs : carts.append(cart(x=x,y=y,dir=c,turn=0))
	# remove them
	tracks[y]=s.replace('<','-').replace('>','-').replace('^','|').replace('v','|')

ncarts = len(carts)
print (ncarts,'carts')

while ncarts>=2 : # accounts 0 or 1 left
	#aff()
	carts.sort(key=lambda c:(c.y,c.x)) 
	for c in carts : 
		if c.dir=='o' : continue  # skip crashed
		movecart(c)

		# check collision
		for c2 in carts :
			if id(c2)==id(c) or c2.dir=='o': continue 
			if c.x==c2.x and c.y==c2.y  : 
				# 'remove' those carts *immediately* : ie mark as crashed
				c.dir='o'
				c2.dir='o'
				ncarts-=2
				print('crash at',c.x,c.y,ncarts,'carts left')
		if len(tracks)<10 : aff()

print ([c for c in carts if c.dir!='o'])
# not 47,135 ...