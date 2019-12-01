use std::fs::File;
use std::io::{BufRead, BufReader, Result};

fn fuel(mass:i32) -> i32 {
	let mut tot=0;
	let mut curr_mass=mass;
	loop {
		let fu = curr_mass/3-2;
		if fu>=0 {
			tot += fu;
			curr_mass = fu;
		} else {
			return tot;
		}
	}
}

fn main() -> Result<()> {
	println!("fuel for {} is {}",1969,fuel(1969));

    let file = File::open("01.txt")?;
    let mut total = 0;

    for line in BufReader::new(file).lines() {
    	let mass = line?.parse::<i32>().expect("could not parse int");
    	total += fuel(mass);
    }
    println!("total: {}", total);
    Ok(())
}
