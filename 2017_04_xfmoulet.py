n1 = n2 = 0
for line in open("2017_04.data") :
	words = line.split()
	sorted_words = [''.join(sorted(w)) for w in words]

	if len(words) == len(set(words)) : n1 += 1
	if len(sorted_words) == len(set(sorted_words)) : n2 += 1

print(n1,n2)