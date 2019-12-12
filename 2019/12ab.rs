//const PLANETS : [[i64;3];4]  = [ [-1,0,2], [2,-10,-7], [4,-8,8], [3,5,-1] ];
//const PLANETS : [[i64;3];4]  = [ [-8,-10, 0],[5, 5, 10],[2, -7, 3],[9, -8, -3] ];
const PLANETS : [[i64;3];4] = [ [17, 5, 1],[-2, -8, 8],[7, -6, 14],[1, -10, 4] ];

fn simulate_axis_step(pos:&mut [i64;4], vel:&mut[i64;4]) {
	let mut new : [i64;4] = [0;4];
    for i in 0..4 { // for each planet	 
    	let acc : i64 = pos.iter().map(|p2| (p2-pos[i]).signum()).sum();
        vel[i] += acc; // acceleration 
        new[i] = pos[i]+vel[i];
    }
    pos.clone_from_slice(&new); // copy new values to pos
}

fn gcd(mut a: i64, mut b: i64) -> i64 {
    while b != 0 {
        let tmp = a;
        a = b;
        b = tmp % b;
    }
    a
}

fn lcm(a:i64,b:i64) -> i64 {
	a*b/gcd(a,b)
}

fn main() {
	let mut planets    = PLANETS;
	let mut velocities : [[i64;3];4] = [[0;3];4];

	for axis in 0..3 {
		// initialize position
		let mut pos : [i64;4] = [0;4];
		for i in 0..4 { pos[i]=planets[i][axis] }
		let mut vel : [i64;4] = [0;4];

		for _i in 1..=1000 {
			simulate_axis_step(&mut pos,&mut vel);
		}
		for i in 0..4 { 
			planets[i][axis]=pos[i]; 
			velocities[i][axis]=vel[i];
		}
	}
	// println!("pos:{:?} vel:{:?}",planets, velocities);
	let mut energy = 0;
	for (pos,vel) in planets.iter().zip(velocities.iter()) {
		energy += pos.iter().map(|x| x.abs()).sum::<i64>() * vel.iter().map(|x| x.abs()).sum::<i64>()
	}
	println!("A: End energy {:?}",energy);

	// --------- part B

	let mut rev = [0_i64;3];

	for axis in 0..3 {
		// initialize position
		let mut pos : [i64;4] = [0;4];
		for i in 0..4 { pos[i]=PLANETS[i][axis] }
		let mut vel : [i64;4] = [0;4];

		let past = (pos,vel).clone(); 

		for i in 1..=1_000_000 {
			simulate_axis_step(&mut pos,&mut vel);
			if (pos,vel) == past { 
				rev[axis]=i;
				break;
			}
		}
	}
	println!("B: Steps to first state: {:?} (from {:?})",lcm(lcm(rev[0],rev[1]),rev[2]), rev)
}
