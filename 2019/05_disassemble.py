import sys

elts = [int(x) for x in open(sys.argv[1]).read().split(',')]
el_it=iter(elts)
ip=0

try : 
	while True :
		op = next(el_it)
		INSTRUCTIONS = {
			1:("ADD",3),		2:("MUL",3),
			3:("INPUT",1),		4:("OUTPUT",1),
			5:("JE",2),			6:("JNE",2),
			7:("LT",3),			8:("EQ",3),
			99:("END",0),
			}
		instr,n = INSTRUCTIONS.get(op%100,(f'WTF:{op}',0))
		operands = [next(el_it) for _ in range(n)]
		print(ip,instr,operands,sep='\t')
		ip += n+1
except StopIteration : 
	pass
