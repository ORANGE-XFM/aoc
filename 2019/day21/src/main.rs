use std::fs::read_to_string;
use intcode::{Program,Word};


fn run_prog(line: &str, input : &str) {
	println!("====");
	let mut p = Program::new(&line);
	for c in input.bytes() {
		p.input(c as Word);
	}

    loop {
    	match p.next() {
    		Some(c) => { 
    			if c<256 { print!("{}",c as u8 as char) } 
    			else { print!("{}",c ) }
    	},
    		None => break,
    	}
    }
}


fn main() {
    let line0 = read_to_string("../day21.txt").expect("Something went wrong reading the file");
    let jumpcode_a = "NOT C J
AND D J
NOT A T
OR T J
WALK
";
    run_prog(&line0, &jumpcode_a);

    run_prog(&line0, &"OR D J
AND H J
NOT C T
AND C J
NOT A T
OR T J
RUN
");

    //starB(&line0);
}
