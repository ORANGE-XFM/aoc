from collections import defaultdict

debug = False

steps = defaultdict(list) # to : [from]
todo = set()
for l in open('2018_07_ex.data' if debug else '2018_07.data') : 
	fr,to = l[5],l[36]
	steps[to].append(fr)
	todo.add(fr); todo.add(to)

for k,v in sorted(steps.items()) : 
	print (k,v)
print('-'*80)

workers=tuple([None,0] for _ in range(2 if debug else 4)) # task or '.' if idle

done = []

def doable_task() : 
	doable = sorted(x for x in todo if all(pre in done for pre in steps[x]))
	return doable[0] if doable else None

ntodo = len(todo)
time=0
while len(done)!=ntodo : # wait finished, not nothing to do 
	# check done
	for w in workers : 
		if w[1]==0 and w[0]: 
			done.append(w[0])
			w[0]=None

	# start new tasks
	for w in workers : 
		if w[0]==None : 
			task = doable_task()			
			if task : 
				w[0]=task
				w[1]=ord(task)-ord('A')+(0 if debug else 60)
				todo.remove(task)

		else : 
			w[1]-=1

	print(time, workers,'done',''.join(done),'todo',todo)
	time +=1

print(done, time-1)


