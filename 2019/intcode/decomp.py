"""
use : decomp.py program.txt  map.txt 
map.txt is a file with 
    name addr
    name addr
    ... 

where name, addr : name this address
	special name : start_data (interpret code as data from here )
"""
import sys
program = [int(x) for x in open(sys.argv[1]).read().split(',')]
data_from = 1000000 # never
rev_addr = {}
if len(sys.argv)>2 : # provide map.txt ?
	addrs = dict(s.split() for s in open(sys.argv[2]))
	data_from = int(addrs.get('start_data',data_from))
	rev_addr = {'@'+v:k for k,v in addrs.items()}

INSTRUCTIONS = {
	1:"{2}={0}+{1}",    2:"{2}={0}*{1}",
	3:"{0}=input",      4:"output({0})",
	5:"if {0} jmp {1}", 6:"if not {0} jmp {1}",
	7:"{2} = {0}<{1}",  8:"{2} = {0}=={1}",
	9:"base <- {0}",
	99:"end.",
	}

ip=0
el_it=iter(program)
try : 
	while ip<data_from :
		op = next(el_it)
		instr = INSTRUCTIONS.get(op%100,f'WTF:{op}')
		n=sum(1 for c in instr if c=='{')
		def opmode(i) : return ('@','','BASE+')[op//(10**(2+i))%10]
		operands = [opmode(i)+str(next(el_it)) for i in range(n)]
		operands = [rev_addr.get(op,op) for op in operands]
		if op%100 in (5,6) : operands[1] = rev_addr.get('@'+operands[1],operands[1]) # jmp : treat op1 as indirect

		if '@'+str(ip) in rev_addr : print('\n;',rev_addr['@'+str(ip)],'---')
		print(ip,instr.format(*operands),sep='\t')
		ip += n+1
	# output data
	while True :
		if '@'+str(ip) in rev_addr : print('\n;',rev_addr['@'+str(ip)],'---')
		print(ip,next(el_it),sep='\t')
		ip += 1
except StopIteration : 
	pass
