use std::fs::File;
use std::io::{BufRead, BufReader, Result};
use std::collections::HashMap;


fn ancestors(planets:&HashMap<String,String>, planet:String) -> Vec<String>{
	if planet=="COM" {
		return Vec::new();
	} else {
		let mut v = ancestors( planets, planets.get(&planet).unwrap().to_string() );
		v.push(planet);
		v
	}
}

fn common_ancestors(a : &Vec<String>,b:&Vec<String>) -> u32 {
	let mut n=0;
	for (s1,s2) in a.iter().zip(b.iter()) {
		n+=1;
		println!("{:?} {:?}",s1,s2 );
		if s1 != s2 {
			return n;
		}
	}
	n
}

fn main() -> Result<()> {
    let file = File::open("06.txt")?;
	let mut planets = HashMap::new();
    for l in BufReader::new(file).lines() {
    	let line= l?;
    	let v : Vec<&str>= line.split(")").collect();
    	let (center,around) = (v[0].to_string(),v[1].to_string());
	    planets.insert(around,center);
    }

    let ancestors_you = ancestors(&planets,"YOU".to_string());
    let ancestors_san = ancestors(&planets,"SAN".to_string());
    println!("{:?} planets",planets.iter().count());
    println!("YOU orbits {:?}",ancestors_you);    
    println!("SAN orbits {:?}",ancestors_san);
    let common = common_ancestors(&ancestors_san,&ancestors_you) as usize;
    println!("common elements {:?}, nb : {}", common, ancestors_you.iter().count()+ancestors_san.iter().count()-2*common);

    Ok(())
}
