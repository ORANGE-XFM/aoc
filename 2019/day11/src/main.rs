use std::fs::read_to_string;
use intcode::{Program, Word};

const W : usize = 100;

enum Dir {
	Up,Down,Left,Right
}

fn printboard (board: &[Word;W*W]){
	for y in 0..W {
		for x in 0..W {
			print!("{}", match board[y*W+x] {0=>'.',1=>'#',_=>' '});
		}
		println!("");
	}
}

fn main() {
    let line0 = read_to_string("11.txt").expect("Something went wrong reading the file");
    let mut p = Program::new(&line0);

    let mut board : [Word;W*W] = [-1;W*W ];
    let mut x:usize=W/2;
    let mut y:usize=W/2;
    let mut facing = Dir::Up;

	board[y*W+x] = 1; // 2nd step

    loop {
    	p.input(if board[y*W+x]==1 {1} else {0} ); // case -1 =>0 

    	// paint 
    	let out = p.next();
    	match out {
    		Some(color) => {
    			println!("{} {} {:?}", x,y,color );
    			board[y*W+x] = color;
    		}
    		None => break,
    	}

    	// turn
    	let turn = p.next();
    	match turn {
    		Some(dir) => {
    			facing = 
    				if dir==0 { // left ?
    					match facing { Dir::Up => Dir::Left,  Dir::Left => Dir::Down,  Dir::Down=>Dir::Right, Dir::Right => Dir::Up }
    				} else {
    					match facing { Dir::Up => Dir::Right, Dir::Right => Dir::Down, Dir::Down=>Dir::Left,  Dir::Left => Dir::Up }
    				}
    		}
    		None => break,
    	}

		// move
		match facing { 
			Dir::Left => x-=1, 
			Dir::Right => x+=1, 
			Dir::Up => y-=1,   
			Dir::Down  => y+=1, 
		}
    }
    println!("{:?}", board.iter().filter(|x| **x != -1 ).count());
    printboard(&board);



}
