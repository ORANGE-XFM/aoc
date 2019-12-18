use std::fs::read_to_string;

fn gen_line(size:usize, n:usize) -> Vec<i32> {
	let mut v = vec![0;size];
	let pattern = [0,1,0,-1];
	for i in 0..size {
		v[i] = pattern[((i+1)/n)%4];
	}
	v
}


/* multiply by a matrix of 

111111111111
011111111111
   ...
000000000111
000000000011
000000000001
*/
fn mul_tri_matrix(v: &Vec<i32>, to:&mut Vec<i32>) {
	to[v.len()-1] = v[v.len()-1];
	for i in (0..v.len()-1).rev() {
		to[i] = to[i+1]+v[i];
	}
}

fn main() {
	let line = read_to_string("16.txt").expect("Something went wrong reading the file");
	// let line = "12345678";
	// let line = "80871224585914546619083218645595";
	//let line = "03036732577212944063491565474664";

	let start = line[..7].parse().unwrap();

	let singlevec : Vec<i32> = line.trim().chars().map(|x| x.to_digit(10).expect("not a digit") as i32).collect();

	// ------------------------------------------------------------------	
	println!("Part A ------");

	let mut v1 = singlevec.clone();
	println!("v1={:?}", v1);
	// println!("a1 {:?}",gen_line(8,1) );

	let mut nv = Vec::new();

	for nb in 0..100 {
		nv.clear();	
		for i in 1..=v1.len() {
			let d = v1.iter().zip(gen_line(v1.len(),i)).map(|cpl| cpl.0*cpl.1).sum::<i32>();
			nv.push((d%10).abs());
		}
		if nb%10 == 0 { println!("{:?}", nb); }
		v1.clone_from_slice(&nv);
	}
	println!("{:?}",v1);

	// ------------------------------------------------------------------	

	println!("Part B ------");

	let mut v0 = Vec::new();
	for _i in 0..10000 { v0.append(&mut singlevec.clone()); } // duplicate
	// take only end part
	println!("{}",v0.len());
	let mut v1 : Vec<i32> = v0[start..].to_vec();
	println!("{}",v1.len());

	let mut nv = vec![0;v1.len()];

	for _nb in 0..100 {
		mul_tri_matrix(&v1,&mut nv);
		for i in 0..nv.len() {
			nv[i] = (nv[i]%10).abs();
		}
		v1.clone_from_slice(&nv);
	}
	for i in 0..8 {
		print!("{}",v1[i]);
	}
	println!("");
	println!("{:?} --",&v1[..8]);
}
