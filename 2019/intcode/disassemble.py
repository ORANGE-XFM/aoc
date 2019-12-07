import sys

elts = [int(x) for x in open(sys.argv[1]).read().split(',')]
el_it=iter(elts)

ip=0
try : 
	while True :
		op = next(el_it)
		INSTRUCTIONS = {
			1:("add",3),		2:("mul",3),
			3:("input",1),		4:("output",1),
			5:("je",2),			6:("jne",2),
			7:("lt",3),			8:("eq",3),
			99:("end",0),
			}
		instr,n = INSTRUCTIONS.get(op%100,(f'WTF:{op}',0))
		def opmode(i) : return '' if op//(10**(2+i))%10 else '@'
		operands = [opmode(i)+str(next(el_it)) for i in range(n)]
		print(ip,instr,*operands,sep='\t')
		ip += n+1
except StopIteration : 
	pass
