import heapq

class Cave :
    def __init__(self, depth, target_x, target_y ) :
        # A region's erosion level is its geologic index plus the cave system's depth, all modulo 20183
        def erosion(level) :
            return (level+depth)%20183

        cave = [[0 for _ in range(target_x+100)] for _ in range(target_y*2+1)]
        self.cave = cave
        self.depth = depth
        self.target_x = target_x
        self.target_y = target_y

        cave[0][0] = erosion(0)
        
        # If the region's Y coordinate is 0,
        # the geologic index is its X coordinate times 16807
        for x in range(1,self.w) : 
            cave[0][x] = erosion(x*16807)
        
        # If the region's X coordinate is 0, 
        # the geologic index is its Y coordinate times 48271
        for y in range(1,self.h) : 
            cave[y][0] = erosion(y*48271)

        # Otherwise, the region's geologic index is the result of 
        # multiplying the erosion levels of the regions at X-1,Y and X,Y-1.
        for y in range(1,self.h) : 
            for x in range(1,self.w) :
                if x==target_x and y==target_y :                     
                    cave[y][x] = erosion(0)
                else : 
                    cave[y][x] = erosion(cave[y-1][x]*cave[y][x-1])


    @property
    def w(self) : return len(self.cave[0])

    @property
    def h(self) : return len(self.cave)

    def print(self) : 
        for line in self.cave : 
            print(line)

    def print(self) : 
        for line in self.cave : 
            print(''.join(".=|"[c%3] for c in line))
    

    def risk_level(self) : 
        return sum (
            sum( v%3 for v in line[:self.target_x+1] )
            for line in self.cave[:self.target_y+1]
            )


    def path(self) : 
        "shortest path from 0,0 to target. use Dijkstra, from day 15. State has now the tool+pos"
        NONE, TORCH, CLIMB = 0,1,2
        OTHERS = [(1,2),(0,2),(0,1)] # possible tools from terrain type

        frontier=[] # cost, start( position, tool)
        cost_so_far={} # pos,tool : cost so far
        come_from={} # pos, tool : pos, tool

        def add_cost(next_x, next_y, next_tool, new_cost, comefrom) : 
            next = (next_x,next_y,next_tool)
            if next not in cost_so_far or new_cost < cost_so_far[next] :
                heapq.heappush(frontier, (new_cost,next))
                cost_so_far[next] = new_cost
                come_from[next] = comefrom

        add_cost(0,0,TORCH,0,None)

        while frontier : 
            dist,current = heapq.heappop(frontier)
            current_x,current_y,current_tool = current

            # change tool
            for next_tool in OTHERS[ self.cave[current_y][current_x]%3] : 
                add_cost(current_x,current_y,next_tool,cost_so_far[current]+7,current)

            # move
            move_cost = cost_so_far[current] + 1 
            for dx,dy in ((0,1),(0,-1),(1,0),(-1,0)):
                next_x,next_y  = current_x+dx, current_y+dy
                if next_x<0 or next_y<0 or next_x==self.w or next_y==self.h:
                    continue # out of grid, skip

                next_type = self.cave[next_y][next_x]%3

                # need to change tool ? if so change it now, try both possibilities
                if next_type != current_tool : # change it
                    add_cost(next_x,next_y,current_tool,move_cost,current)

        return min(
            cost_so_far.get((self.target_x,self.target_y,tool),9999999999)+cost 
            for tool,cost in ((TORCH,0),(CLIMB,7)) 
            )

# ---------------------------


def main() : 
    cave = Cave(5355, 14, 796)
    print ("risk",cave.risk_level())
    print ("path dist",cave.path())

if __name__ == '__main__':
    main()
