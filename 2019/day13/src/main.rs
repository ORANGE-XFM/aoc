use std::fs::read_to_string;
use intcode::{Program, Word};

const W : usize = 50;

fn printboard (board: &[Word;W*W]){
	for y in 0..W {
		for x in 0..W {
			print!("{}", match board[y*W+x] {
				0=>' ', // empty
				1=>'#', // wall
				2=>'*', // block
				3=>'=', // paddle
				4=>'o', // ball
				_=>'?'
			});
		}
		println!("");
	}
}


fn play_board(line0 : &str, coins: i64) {
    let mut p = Program::new(&line0);
    p.write_mem(0,coins);
    let mut board : [Word;W*W] = [0;W*W ];

	let mut paddle_x :i64=0;
	let mut ball_x :i64;
	let mut score =0;

    loop {
    	let x : i64;

    	match p.next() {
    		Some(pos) => { x = pos;}
    		None => break,
    	}
    	let y = p.next().unwrap();
    	let out = p.next().unwrap();

    	if x==-1 {
    		score = out;
    	} else {
			board[(y*(W as i64)+x) as usize] = out;
			match out {
				3 => { // paddle
					paddle_x = x;
    				//println!("paddle {:?}",(ball_x,paddle_x,(paddle_x-ball_x).signum()) );
    				//printboard(&board);
				},
				4 => {
					ball_x = x;
    				//println!("ball {:?}",(ball_x,paddle_x,(paddle_x-ball_x).signum()) );
    				//printboard(&board);
    				p.input((ball_x-paddle_x).signum()); // paddle comes after the ball on a frame
				},
				_=>{}
			}			
    	}
    }
    printboard(&board);
    println!("{:?} blocks, score= {}", board.iter().filter(|x| **x == 2 ).count(), score);

}

fn main() {
    let line0 = read_to_string("13.txt").expect("Something went wrong reading the file");
    play_board(&line0,1);
    play_board(&line0,2);
}
