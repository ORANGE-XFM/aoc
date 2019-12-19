use std::fs::read_to_string;
use intcode::Program;

fn line(y: i64, prog: &Program) -> (i64,i64) {
	// gets first, last ones for a line, slow method
	let mut x0=0;
	loop {
		let mut p = prog.clone();
		p.input(x0); p.input(y);
		let c = p.next().unwrap();
		if c == 1 { break }
		x0 += 1;
		if y>0 && (x0/y)>100 {
			return (0,-1) // no # on line 
		}
	}
	let mut x1=x0+1; // this is the first zero
	loop {
		let mut p = prog.clone();
		p.input(x1); p.input(y);
		let c = p.next().unwrap();
		if c == 0 { break }
		x1 += 1;
	}
	(x0,x1-1)
}

// print it / count it 
fn _print_50(prog_start: &Program) {
	let mut nb=0;
    for y in 0..50 {
	    print!("{:3} ",y);
	    for x in 0..50 {
	    	let mut p=prog_start.clone();
	    	p.input(x); p.input(y);
	    	match p.next() {
	    		Some(c) => if c==1 {print!("#");nb += 1} else {print!(".")},
	    		None => print!("?"),
	    	}
	    }
	    println!("");
	}
	println!("{:?}",nb);
}

fn star_a(prog_start: &Program) {
	println!("{}", (0..50).map(|y| {let c=line(y,prog_start); c.1-c.0+1}).sum::<i64>());
}

fn star_b(prog_start: &Program) {
	let y = 1328; // found by dichotomy, could be faster with slope first approx, .. 
	//let y = 45; // found by dichotomy, could be faster with slope .. 
	let sz = 100;

	let (a0,a1) = line(y,&prog_start);
	let (b0,b1) = line(y+sz-1,&prog_start);
	println!("{}:{}--{}\n{}:  {}--{}\n{} == {}, pos={}",y,a0,a1,y+sz-1,b0,b1,a1-b0+1,sz,b0*10000+y);

    println!("");
}


fn main() {
    let prog_src = read_to_string("../19.txt").expect("Something went wrong reading the file");
    let prog = Program::new(&prog_src);
    _print_50(&prog);
    star_a(&prog);
    star_b(&prog);
}
