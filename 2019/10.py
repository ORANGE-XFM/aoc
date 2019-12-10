from cmath import *
from math import pi 

coords = []
for y,l in enumerate(open('10.txt')) : 
	coords += [x+y*1j for x,c in enumerate(l) if c == '#']

def pphase(x) : return ((pi/2-phase(x.conjugate()))%(2*pi))
# print([pphase(x)/pi for x in (-1j,1-1j,1,1+1j,1j,-1+1j,-1,-1-1j)]) # test
def unique_phases(c) : return set(pphase(a-c) for a in coords if a !=c )

# part A
print ( max(len(unique_phases(c)) for c in coords))

# part B
center =  max(coords, key=lambda c:len(unique_phases(c)))
print('center:',center)

phases = sorted(unique_phases(center))
print( '200th shooted:', min((c for c in coords if pphase(c-center)==phases[199]), key = lambda x: abs(center-x)))
