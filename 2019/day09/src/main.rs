use std::fs::read_to_string;
use intcode::{Program, Word};

fn main() {
    let line0 = read_to_string("09.txt").expect("Something went wrong reading the file");
    let mut p = Program::new(&line0);
    p.input(1);
    let out : Vec<Word> = p.collect();
    println!("{:?}", out);

    let mut p = Program::new(&line0);
    p.input(2);
    let out : Vec<Word> = p.collect();
    println!("{:?}", out);
}
