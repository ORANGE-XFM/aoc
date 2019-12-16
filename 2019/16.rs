use std::fs::read_to_string;

fn gen_line(size:usize, n:usize) -> Vec<i32> {
	let mut v = vec![0;size];
	let pattern = [0,1,0,-1];
	for i in 0..size {
		v[i] = pattern[((i+1)/n)%4];
	}
	v
}

fn main() {
	let line = read_to_string("16.txt").expect("Something went wrong reading the file");
	// let line = "12345678";
	// let line = "80871224585914546619083218645595";

	let mut v1 : Vec<i32> = line.trim().chars().map(|x| x.to_digit(10).expect("not a digit") as i32).collect();
	println!("v1 {:?}", v1);
	println!("a1 {:?}",gen_line(8,1) );

	let mut nv = Vec::new();

	for _nb in 0..100 {	
		for i in 1..=v1.len() {
			let d = v1.iter().zip(gen_line(v1.len(),i)).map(|cpl| cpl.0*cpl.1).sum::<i32>();
			nv.push((d%10).abs());
		}
		v1.clone_from_slice(&nv);
		println!("{:?}",v1);
	}
}
