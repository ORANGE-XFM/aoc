%: %a %b
	./$*a < $*.txt # ------------
	./$*b < $*.txt  # ------------

%: %ab
	./$*ab < $*.txt

%:: %.rs
	rustc $*.rs

clean:
	git clean -xi
	
05_example.txt: 05_example.s
	python3 05_assembler.py $^ > $@
