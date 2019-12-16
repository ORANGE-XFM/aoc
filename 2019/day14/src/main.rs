extern crate regex;
use std::io::{BufRead, BufReader,Result};
use std::fs::File;
use regex::Regex;
use std::collections::HashMap;



fn main() -> Result<()> {
	let file = File::open("../14_ex1.txt")?;
    let file = BufReader::new(file);

    let item_regex = Regex::new(r"(\d+) (\w+)").unwrap();
    let mut recipes = Vec::new();
    for line in file.lines() {
    	let mut recipe = Vec::new();

    	for mat in item_regex.captures_iter(&line?) {
    		let nb = mat[1].parse::<i32>().expect("cannot parse int");
    		let wat = mat[2].to_string();
    		recipe.push((nb,wat));
		}
		recipes.push(recipe);
    }
    print!("{:?}",recipes );

    // now we have the parsed recipes. 

    let mut needed = HashMap::new();
    let mut _needed_ore = 0;

    needed.insert(1,"FUEL");
    loop {
	    println!("{:?}", needed);

	    // get an arbitrary element
	    let (k,v) = needed.iter().next().unwrap().clone(); // panics if not exist ...
	    needed.remove(&k);
	    /*

	    match ans {
	    	Some((k,v)) => {
	    		println!("{:?}", (&k,&v));
	    		needed.remove(&k);
	    	},
	    	None => break,
	    }
	    */
    }
    println!("done.");

    Ok(())
}
