use std::fs::File;
use std::io::{BufRead, BufReader, Result};
use std::collections::HashMap;


fn get_orbits(planets:&HashMap<String,Vec<String>>, planet:&String, base:u32) -> u32{
	match planets.get(planet) {
		Some(v) => {
			base + v.iter().map(|sub| get_orbits(planets,sub,base+1)).sum::<u32>()
		},
		None => base,
	}
}

fn main() -> Result<()> {
    let file = File::open("06.txt")?;
	let mut planets = HashMap::new();
    for l in BufReader::new(file).lines() {
    	let line= l?;
    	let v : Vec<&str>= line.split(")").collect();
    	let (center,around) = (v[0].to_string(),v[1].to_string());
	    planets.entry(center).or_insert_with(|| Vec::<String>::new() ).push(around);
    }
    println!("{:?} planets, {} orbits",planets.keys().count(),get_orbits(&planets,&"COM".to_string(),0) );
    Ok(())
}
