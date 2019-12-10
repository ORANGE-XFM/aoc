use std::fs::File;
use std::io::{BufRead, BufReader};
use std::collections::{HashSet, BTreeSet};

fn parse_input(filename: &str) -> Vec<(i32,i32)>{
	let mut coords = Vec::new();
    let file = File::open(filename).expect("could not open file");
    for (y,line) in BufReader::new(file).lines().enumerate() {
    	for (x,c) in line.expect("could not read line").chars().enumerate() {
    		if c=='#' {
    			coords.push((x as i32,y as i32))
    		}
    	}
    }
    coords
}

// angle -> 0->2*pi ! XXX
fn angle(v:&(i32,i32)) -> i32 {
	let (dx,dy) = *v;
 	let angle = (-dx as f32).atan2(dy as f32); 
 	return if angle <= 3.141592 { ((angle+3.1415926535)*100000_f32).floor() as i32 } else { 0 } ;
}

fn angles(coords:&Vec<(i32,i32)>, i : usize) -> HashSet<i32> {
	let center = coords[i];
	let mut ang = HashSet::new(); // unique angles
	for (x,y) in coords {
		if (*x,*y) == center { continue }
		ang.insert(has_angle((*x,*y),center));
	}
	ang
}

// 200th angle in order, get closest 
fn laser_angle(coords:&Vec<(i32,i32)>, i : usize) -> i32 {
	let (xc,yc) = coords[i];
	let mut angles = BTreeSet::new(); // unique angles, ordered
	for (x,y) in coords {
		if (*x,*y) == (xc,yc) { continue }
		let dx = *x-xc;
		let dy = *y-yc;

		angles.insert(angle(&(dx,dy)));
	}

	println!("{} {:?}",1,angles.iter().nth(0));
	println!("{} {:?}",2,angles.iter().nth(1));
	for i in 195..205 {
		println!("{} {:?}",i+1,angles.iter().nth(i));
	}

	println!("{:?}",angles);
	let angle200 = *angles.iter().nth(199).unwrap(); // 1 less because 
	angle200
}

fn has_angle(v:(i32,i32), c:(i32,i32)) -> i32 {
	angle(&(v.0-c.0,v.1-c.1))
}

fn main() {
	let a = parse_input("10.txt");
	let center_id=(0..a.len()).max_by_key(|i| angles(&a,*i).len()).unwrap();
	let all_angles = angles(&a,center_id);

	println!("Result for a: {} at {:?}", all_angles.len(), a[center_id]);
	println!("angle 1 {}",has_angle((11,12),(11,13)));
	println!("angle 2 {}",has_angle((12,1),(11,13)));
	println!("angle 3 {}",has_angle((12,2),(11,13)));

	println!("angle 200 {}",has_angle((8,2),(11,13)));
	println!("angle 201 {}",has_angle((10,9),(11,13)));
	// let sorted_angles = all_angles.iter().sorted();
	// println!("{:?}",sorted_angles );


	let angle_200 = laser_angle(&a,center_id);
	let asteroids : Vec<&(i32,i32)> = a.iter().filter(|v| has_angle(**v,a[center_id])==angle_200).collect();
	
	println!("Result for b: {} {:?}", angle_200, asteroids);
}
