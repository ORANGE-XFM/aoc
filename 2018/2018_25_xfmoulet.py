def read(filename) : 
	stars = set()
	for l in open(filename) : 
		stars.add(tuple(int(s) for s in l.split(',')))
	return stars 

def near (s1,s2) : 
	return sum(abs(x1-x2) for x1,x2 in zip(s1,s2))<=3

def extend_constellation(star,stars, constellation) : 
	constellation.add(star)
	for i,star2 in enumerate(stars-constellation) : 
		if near(star,star2) : 
			extend_constellation(star2, stars-constellation, constellation)
	return constellation

stars = read('2018_25_xfmoulet.txt')

print(len(stars), stars)

print('-'*40)
n=0
remain = stars 
while remain : 
	const = extend_constellation(remain.pop(), stars, set())
	print('const',len(const), const)
	remain -= const
	n+=1
print('nb constellations',n)

