p=[int(x) for x in open('02.txt').read().split(',')]
print(p)


def prog(p0,a,b) :
	p = p0[:]
	p[1]=a
	p[2]=b
	i=0
	while True: 
		if p[i]==99 : break
		a=p[i+1]
		b=p[i+2]
		c=p[i+3]
		if p[i]==1 : 
			p[c]=p[a]+p[b]
		elif p[i]==2 : 
			p[c]=p[a]*p[b]
		i+=4
	return p[0]

for i in range(99) : 
	for j  in range(99) : 
		if prog(p,i,j)==19690720 : 
			print('****')
			print(i,j,prog(p,i,j),19690720)
