from collections import defaultdict

f=open('2018_12.data')
state = '.'*20+next(f)[15:-1]+'.'*20 # 20 chars before/after
next(f)
rules = defaultdict(lambda :'.')
rules.update((l.strip().split(' => ') for l in f),)

for k,v in rules.items() : 
	print (k,'->',v)
print()

# next states
END=50000000000
prev=None
for i in range(END) : 
	print (i,state)
	l=[]
	for k in range(len(state)) :
		substr = state[k:k+5]
		l.append(rules[substr])
	prev=state
	state = '..'+''.join(l)
	# detect and fast forwards right sliders
	if '.'+prev+'.'==state : 
		print('we have a slider !')
		break

print ('final state')
print (i,state)

# score ( with sliders !)
score=0
for k,c in enumerate(state) : 
	if c=='#' : score += k-20 + END-i-1 # account for sliders
print(score)