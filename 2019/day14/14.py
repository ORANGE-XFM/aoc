import re
reg = re.compile(r"(\d+) (\w+)")
recipes = {} # produced : qty, ingredients

for l in open('14.txt') : 
	m = [(int(x),y )for x,y in reg.findall(l)]
	n,elt = m[-1]
	recipes[elt]=n,m[:-1]

to_produce = set(recipes)
production = ['ORE']

while to_produce : 
	# get one that can be done 
	print (to_produce)
	for p in to_produce : 
		needed = set(x[1] for x in recipes[p][1])
		if needed.issubset(set(production)) :
			production.append(p)
			to_produce.remove(p)
			break
print (production)

# get quantities
print('-'*80)

def need_ore(nb_fuel) :
	needed = {'FUEL':nb_fuel}
	for to_produce in reversed(production[1:]) :  # avoid ORE
		nb_produced, needed_elements = recipes[to_produce]
		#print("now needed",needed)
		nb_needed = needed.pop(to_produce)
		nb_recipe = (nb_needed+nb_produced-1) // nb_produced
		#print('producing:',nb_produced,to_produce,'with',needed_elements,': make it',nb_recipe,'time')
		for nb,elt in needed_elements : # do we need it for more than one recipe ?		
			needed[elt] = needed.get(elt,0)+nb*nb_recipe
	return needed['ORE']
# A
print(need_ore(1))
# B
needed = 1000000000000
a = 0
b = 1000000000
while (b != a) :
   c = (b + a)//2 # c is like a midpoint
   if c==a : 
   		print('solution',a)
   		break
   n = need_ore(c)
   print(a,b,c,n)
   if n>needed :
      b = c
   else :
      a = c
