

class Grid : 
    def __init__(self,filename,elves_power=3,stop_onedeath=False) : 
        self.items=[] # x,y,live,type
        self.grid=[]
        for y,s in enumerate(open(filename)) : 
            l = []
            for x,c in enumerate(s.strip()) : 
                if c in 'GE' : 
                    self.items.append([x,y,200,c])
                    c = '.'
                l.append(c)
            self.grid.append(l)        
        self.turns = 1
        # part 2
        self.elves_power   = elves_power
        self.stop_onedeath = stop_onedeath # stop if one elf dies


    def print(self) : 
        print(self.items)
        grid2=[l[:] for l in self.grid] # copy
        for y,l in enumerate(self.grid) : 
            for x,c in enumerate(l) :
                for ix,iy,l,ic in self.items: 
                    if ix==x and iy==y : 
                        c = ic
                print(c,end='')
            print()
        print('items',self.items)

    def empty(self, x,y) : 
        if self.grid[y][x]=='#' : return False
        if any(it[0]==x and it[1]==y for it in self.items) : return False
        return True

    def shortestpaths(self,start) : 
        frontier=[(0,start)]

        cost_so_far={start:(None,0)} # came_from, cost so far

        reached = []
        while frontier : 
            frontier.sort()
            dist,current = frontier.pop(0)

            for dx,dy in reversed(((0,-1),(-1,0),(1,0),(0,1))):
                next = (current[0]+dx,current[1]+dy)
                if not self.empty(*next) : continue # collision
                new_cost = cost_so_far[current][1] + 1
                if next not in cost_so_far or new_cost < cost_so_far[next][1] :
                    frontier.append((new_cost,next))
                    cost_so_far[next] = (current,new_cost)
        return cost_so_far

    def all_died(self) : 
        wat = set(x[3] for x in self.items)
        return not('E' in wat and 'G' in wat) # not elves and gobelins in items

    def turn(self) : 
        for item in sorted(self.items,key=lambda i : (i[1],i[0])) : # reading order
            if item[3]=='X' : continue # he so ded

            target_type = 'G' if item[3]=='E' else 'E'
            #print ("item:",item, end=' ')
            # move from item position to targets - what if same dist ?
            cost_so_far = self.shortestpaths((item[0],item[1]))

            target_pos=[]
            move=True
            for i in self.items :
                if i[3]==target_type : 
                    for dx,dy in ((0,-1),(-1,0),(1,0),(0,1)) : 
                        nx,ny = i[0]+dx,i[1]+dy
                        if nx==item[0] and ny==item[1] : 
                            #print('attacking directly item:',i,end=' ')
                            move=False
                            break
                        else :
                            if self.empty(nx,ny) : 
                                target_pos.append((nx,ny))
            if not any (pos in cost_so_far for pos in target_pos) : 
                move = False # cannot move

            if move :
                #print ('targets',target_pos)
                # ok we got our positions
                min_cost = min(cost_so_far[pos][1] for pos in target_pos if pos in cost_so_far) # only take reachable ones
                min_pos = [pos for pos in target_pos if pos in cost_so_far and cost_so_far[pos][1]==min_cost] # take all target positions for this cost 

                target = sorted(min_pos, key=lambda i: (i[1],i[0]))[0] # take first one in reading order
                #print('selected target',target)

                # now find path : take reverse path problem from end to starts
                reverse_costs = self.shortestpaths(target)
                #print('reverse costs',reverse_costs)
                min_step=None
                min_cost=9999
                for dx,dy in ((0,-1),(-1,0),(1,0),(0,1)) : # read order : up, left,right,down
                    next = item[0]+dx,item[1]+dy
                    if next not in reverse_costs : continue # not reachable
                    cost=reverse_costs[next][1]
                    if cost<min_cost : 
                        min_step = next
                        min_cost = cost

                # move to next one !
                #print('moving to ',min_step)
                item[0],item[1]=min_step
                #self.print()

            # Attacking - if possible
            attacked=None
            for dx,dy in ((0,-1),(-1,0),(1,0),(0,1)) : # read order : up, left,right,down
                nx,ny = item[0]+dx,item[1]+dy
                for n,it in enumerate(self.items) : 
                    if it[0]==nx and it[1]==ny and (attacked==None or it[2]<attacked[2]) and it[3] == target_type: # found an enemy on this place which has less life points ?
                        attacked = it
            if attacked : 
                #print('attacking',attacked)
                attacked[2] -= 3 if item[3]=='G' else self.elves_power
                if attacked[2]<=0 : # die & move out
                    print('Killed',attacked)
                    
                    # onedeath and one elf dead ?
                    if self.stop_onedeath and attacked[3]=='E' : 
                        return False

                    attacked[0]=99999
                    attacked[1]=99999
                    attacked[2]=0
                    attacked[3]='X'


            if self.all_died() : 
                life = sum(x[2] for x in self.items)
                print ('All ded ! ',self.turns-1,'turns',life,'score',(self.turns-1)*life,'elves power',self.elves_power)
                self.print()
                return True # abort turn, stop all

        print('END TURN ',self.turns,'='*80)
        self.turns += 1
        return None # continue



def main() : 

    # part one
    grid = Grid('2018_15.data',3,False)
    while True : 
        grid.print()
        if grid.turn() : break

    # part two
    for elves_power in range(3,9999) : 
        print('POWER TO THE ',elves_power)
        grid = Grid('2018_15.data',elves_power,True)
        while True : 
            #grid.print()
            res = grid.turn()
            if res == False : break # one ded, one more powa
            if res == True : return

main()