use std::fs::read_to_string;
use intcode::{Program, Word};

const W : usize = 41;

const UNKNOWN : Word = -1;
const WALL : Word = 0;
const FREE : Word = 1;
const GOAL : Word = 2;
const VISITED : Word = 3;
const START : Word = 4;

const NEXT_X : usize = 1039;
const NEXT_Y : usize = 1040;

const STARTX : usize = 21;
const STARTY : usize = 21;

fn printboard (board: &[Word;W*W]){
	for y in 0..W {
		for x in 0..W {
			print!("{}", match board[y*W+x] {WALL=>'#',FREE=>' ',GOAL=>'O',START=>'X',VISITED=>'.',_=>'.'});
		}
		println!("");
	}
}

#[derive(Clone, Copy)]
enum Dir {
    Up=1,Down,Left,Right
}

fn left(facing:&Dir)->Dir {
    match facing { Dir::Up => Dir::Left,  Dir::Left => Dir::Down,  Dir::Down=>Dir::Right, Dir::Right => Dir::Up }
}
fn right(facing:&Dir) -> Dir {
    match facing { Dir::Up => Dir::Right, Dir::Right => Dir::Down, Dir::Down=>Dir::Left,  Dir::Left => Dir::Up }
}

fn fill(board:&[Word;W*W],x:usize,y:usize,toboard:&mut [Word;W*W]) {
    match board[x+y*W] {
        FREE | VISITED => toboard[x+y*W]=GOAL,
        _ => {},
    }
}

fn main() {
    let line0 = read_to_string("../15.txt").expect("Something went wrong reading the file");
    let mut board : [Word;W*W] = [0;W*W ]; // just print it ?

    // CHEAT : decompile the program to read the maze ...
    for y in 0..W {
        // rewind state
        let mut p = Program::new(&line0);
        p.poke(224,1106); // go through walls
        p.poke(225,1);
        // go to top left
        for _ in 0..STARTX { p.input(1); p.next(); } // go north 21 times
        for _ in 0..STARTY { p.input(3); p.next(); } // go west 21 times
        for _ in 0..y { p.input(2); p.next(); } // go south N times
        
        for x in 0..W-1 {
            p.input(4); // go east now
            board[y*W+x+1] = p.next().unwrap(); 
        }
    }
    board[STARTY*W+STARTX] = START;
    printboard(&board);

    // solve it really : turn right algorithm
    let mut p = Program::new(&line0);

    let mut x:usize=STARTY;
    let mut y:usize=STARTX;
    let mut facing = Dir::Up;

    loop {
        board[y*W+x] = 4-board[y*W+x];  // VISITED <--> FREE

        facing = right(&facing);
        p.input(facing as Word);

        let mut res = p.next().unwrap();
        while res==WALL {
            facing = left(&facing);
            p.input(facing as Word);
            res = p.next().unwrap();
        }

        // move
        match facing { 
            Dir::Left => x-=1, 
            Dir::Right => x+=1, 
            Dir::Up => y-=1,   
            Dir::Down  => y+=1, 
        }
        if res==GOAL {
            break
        }
    }
    board[STARTY*W+STARTX] = START;
    printboard(&board);
    println!("{:?}", board.iter().filter(|x| **x == VISITED ).count());

    println!("Part B -------------------------------");
    for i in 0..1000 {
        let mut new_board = board.clone(); 
        for y in 1..W-1 {
            for x in 1..W-1 {
                if board[y*W+x] == GOAL {
                    fill(&board, x-1,y  ,&mut new_board);
                    fill(&board, x+1,y  ,&mut new_board);
                    fill(&board, x  ,y-1,&mut new_board);
                    fill(&board, x  ,y+1,&mut new_board);            
                }
            }
        }
        let nb=board.iter().filter(|x| **x == VISITED || **x == FREE ).count();
        println!("{} {:?}",i, nb);
        if nb==0 { break }
        board.clone_from_slice(&new_board);
    }
    printboard(&board);
}
