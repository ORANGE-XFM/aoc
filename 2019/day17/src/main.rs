use std::fs::read_to_string;
use intcode::{Program,Word};


fn starA(line: &str) {
	let mut p = Program::new(&line);
    loop {
    	match p.next() {
    		Some(c) => print!("{}",c as u8 as char),
    		None => break,
    	}
    }
}

fn starB(line: &str) {
	let mut p = Program::new(&line);
	p.poke(0,2);

	// send ascii data
	let s = "A,B,A,B,C,C,B,A,C,A
L,10,R,8,R,6,R,10
L,12,R,8,L,12
L,10,R,8,R,8
n
";
	for c in s.bytes() {
		println!("{:?}",c );
		p.input(c as Word);
	}

	// print display 
    loop {
    	match p.next() {
    		Some(c) => if c<256 {
    				print!("{}",c as u8 as char)
    			} else {
    				print!("{:?}",c );
    			}
    			,
    		None => break,
    	}
    }
}

fn main() {
    let line0 = read_to_string("../17.txt").expect("Something went wrong reading the file");
    //starA(&line0);
    starB(&line0);
}
