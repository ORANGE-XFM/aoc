use std::fs::File;
use std::io::{BufRead, BufReader, Result};

fn main() -> Result<()> {
    let file = File::open("01.txt")?;
    let mut total = 0;

    for line in BufReader::new(file).lines() {
    	let mass = line?.parse::<i32>().expect("could not parse int");
    	let fuel = (mass/3)-2;
    	total += fuel;
    }
    println!("total: {}", total);
    Ok(())
}
