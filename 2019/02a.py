p=[int(x) for x in open('02.txt').read().split(',')]
print(p)
p[1]=12
p[2]=2
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

print(p[0])
