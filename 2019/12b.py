
def sign(x) : return 1 if x>0 else 0 if x==0 else -1

class vec:
    def __init__(self,x,y,z) :  self.v = (x,y,z)
    def __add__(self,other) : return vec( self.v[0]+other.v[0], self.v[1]+other.v[1], self.v[2]+other.v[2])
    def __sub__(self,other) : return vec( self.v[0]-other.v[0], self.v[1]-other.v[1], self.v[2]-other.v[2])
    def sign(self): return vec(sign(self.v[0]),sign(self.v[1]),sign(self.v[2]))
    def energy(self) : return abs(self.v[0])+abs(self.v[1])+abs(self.v[2])
    def __repr__(self) : return 'vec'+repr(self.v)

#planets = [(vec(*x),vec(0,0,0)) for x in ((-1,0,2), (2,-10,-7), (4,-8,8), (3,5,-1) )]
#planets = [(vec(*x),vec(0,0,0)) for x in (( -8,-10, 0),(5, 5, 10),(2, -7, 3),(9, -8, -3) ) ]
planets = [(vec(*x),vec(0,0,0)) for x in [ (17, 5, 1),(-2, -8, 8),(7, -6, 14),(1, -10, 4) ]]

allXs = set()
allYs = set()
allZs = set()

for i in range(300001):
    #print(i,'---',sum(p.energy()*v.energy() for p,v in planets))
    newplanets = []
    allX = tuple((p[0].v[0],p[1].v[0]) for p in planets) ; allXs.add(allX)
    allY = tuple((p[0].v[1],p[1].v[1]) for p in planets) ; allYs.add(allY)
    allZ = tuple((p[0].v[2],p[1].v[2]) for p in planets) ; allZs.add(allZ)
    #print(allX,allY,allZ)

    for pos,vel in planets : 
        acc = sum(((p2-pos).sign() for p2,_ in planets),vec(0,0,0))
        vel += acc
        newplanets.append((pos+vel, vel))
    planets = newplanets

print (pos,vel)

def gcd(a, b):
    """Return greatest common divisor using Euclid's Algorithm."""
    while b:      
        a, b = b, a % b
    return a

def lcm(a, b):
    """Return lowest common multiple."""
    return a * b // gcd(a, b)


print(len(allXs),len(allYs),len(allZs))
print(lcm(lcm(len(allXs),len(allYs)),len(allZs)))
