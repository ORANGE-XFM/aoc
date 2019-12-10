from cmath import *
from math import pi 

coords = []
for y,l in enumerate(open('10.txt')) : 
	coords += [x+y*1j for x,c in enumerate(l) if c == '#']
print(coords)
def pphase(x) : return ((pi/2-phase(x.conjugate()))%(2*pi))
print([pphase(x)/pi for x in (-1j,1-1j,1,1+1j,1j,-1+1j,-1,-1-1j)])
def unique_phases(c) : return set(pphase(a-c) for a in coords if a !=c )
# part A
print ( max(len(unique_phases(c)) for c in coords))

# part B
center =  max(coords, key=lambda c:len(unique_phases(c)))
print('center:',center)

phases = sorted(unique_phases(center))
# print(phases)
# for c in [11+12j,12+1j,12+2j,12+8j,16,16+9j,10+16j,9+6j,8+2j,10+9j,11+1j] : 
# 	phi = pphase(c-center)
# 	print(c,c-center, phi,phases.index(phi))

print(phases[199], [(c,abs(center-c)) for c in coords if pphase(c-center)==phases[199]])
