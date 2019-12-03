%: %a %b
	./$*a # ------------
	./$*b # ------------

%: %ab
	./$*ab

%:: %.rs
	rustc $*.rs

clean:
	git clean -xi
	
