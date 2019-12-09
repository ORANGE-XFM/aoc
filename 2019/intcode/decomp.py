import sys

elts = [int(x) for x in open(sys.argv[1]).read().split(',')]
el_it=iter(elts)

ip=0
try : 
	while True :
		op = next(el_it)
		INSTRUCTIONS = {
			1:"{2}={0}+{1}",
			2:"{2}={0}*{1}",
			3:"{0}=input",
			4:"output({0})",
			5:"if {0} jmp {1}",
			6:"if not {0} jmp {1}",
			7:"{2} = {0}<{1}",
			8:"{2} = {0}=={1}",
			9:"base <- {0}",
			99:"end.",
			}
		instr = INSTRUCTIONS.get(op%100,f'WTF:{op}')
		n=sum(1 for c in instr if c=='{')
		def opmode(i) : return '' if op//(10**(2+i))%10 else '@'
		operands = [opmode(i)+str(next(el_it)) for i in range(n)]
		print(ip,instr.format(*operands),sep='\t')
		ip += n+1
except StopIteration : 
	pass
