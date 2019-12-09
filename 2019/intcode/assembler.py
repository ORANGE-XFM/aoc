# assembler for lang05.py
import sys
instrs = { # opcode:opnum,nb_args
	  "add" : (1,3),   "mul" : (2,3),
  	"input" : (3,1),"output" : (4,1),
	   "je" : (5,2),   "jne" : (6,2),
	   "lt" : (7,3),	"eq" : (8,3),
	 "base" : (9,2),
	  "end" : (99,0),	
}
labels={} 
instr=[]
for l in open(sys.argv[1]) : 
	l=l.lower().split('#')[0].strip().split()
	if not l : continue
	if l[0].endswith(':'):
		labels[l[0][:-1]]=len(instr)
	elif l[0]=='data' : 
		instr+=[int(x) for x in l[1:]]
	else : 
		op,n = instrs[l[0]]
		op += sum(10**(i+2) for i,x in enumerate(l[1:]) if x.isdigit()) # encodes immediate mode args
		instr += [op]+[int(x) if x.isdigit() else x for x in l[1:]]
		if len(l)!=n+1 : raise SyntaxError(f'instruction {l} should have {n} elements')

instr = [labels.get(x,x) for x in instr] # second pass : replace labels
print(*instr,sep=',')
