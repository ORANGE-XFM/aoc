use intcode::{Program, Word};
use permutohedron::LexicalPermutation;
use std::fs::read_to_string;

fn runmachine_a(pr: &Program, arg: Vec<Word>) -> Option<Word> {
    // runs 5 machines with inputs, same program but clone it 5 times
    let mut p = [pr.clone(), pr.clone(), pr.clone(), pr.clone(), pr.clone()];
    let mut x = 0;

    for i in 0..5 {
        p[i].input(arg[i])
    }

    for i in 0..5 {
        p[i].input(x);
        x = p[i].next()?;
    }
    Some(x)
}

fn get_maxprog(progfile: &str) -> Word {
    let data = read_to_string(progfile).expect("cannot open file ?");
    let prog = Program::new(&data);
    let mut vmax = 0;

    let mut data = [0, 1, 2, 3, 4];
    loop {
        let v = runmachine_a(&prog, data.to_vec()).expect("😢 did not return anything?");
        if v > vmax {
            vmax = v
        }
        if !data.next_permutation() {
            break;
        }
    }
    vmax
}

fn runmachine_loop(pr: &Program, arg: Vec<Word>) -> Word {
    // runs 5 machines with inputs, same program but clone it 5 times
    let mut p = [pr.clone(), pr.clone(), pr.clone(), pr.clone(), pr.clone()];
    let mut x = 0;
    for i in 0..5 {
        p[i].input(arg[i])
    }

    loop {
        let prev_x = x;
        for i in 0..5 {
            p[i].input(x);
            match p[i].next() {
                Some(n) => x = n,
                None => return prev_x,
            }
        }
    }
}

fn looper_prog(progfile: &str) -> Word {
    let data = read_to_string(progfile).expect("cannot open file ?");
    let prog = Program::new(&data);
    let mut vmax = 0;

    let mut data = [5, 6, 7, 8, 9];
    loop {
        let v = runmachine_loop(&prog, data.to_vec());
        if v > vmax {
            vmax = v
        }
        if !data.next_permutation() {
            break;
        }
    }
    vmax
}

fn main() {
    // tests
    let line05test = read_to_string("../05_test.txt").expect("Something went wrong reading the file ");
    let line05test8 = read_to_string("../05_test8.txt").expect("Something went wrong reading the file ");
    let line05 = read_to_string("../05.txt").expect("Something went wrong reading the file ");

    let mut p = Program::new(&line05test);
    p.input(100);
    let res: Vec<Word> = p.collect();
    assert_eq!(res, vec![123]);

    let mut p = Program::new(&line05test8);
    p.input(7);
    let res: Vec<Word> = p.collect();
    assert_eq!(res, vec![999]);

    let mut p = Program::new(&line05test8);
    p.input(8);
    let res: Vec<Word> = p.collect();
    assert_eq!(res, vec![1000]);

    // several outputs
    let mut p = Program::new(&line05);
    p.input(1);
    assert_eq!(p.last(), Some(13210611));

    assert_eq!(get_maxprog("ex_a1.txt"), 43210);
    assert_eq!(get_maxprog("ex_a2.txt"), 54321);
    assert_eq!(get_maxprog("ex_a3.txt"), 65210);
    println!("{:?}", get_maxprog("prog.txt"));

    assert_eq!(looper_prog("ex_b1.txt"), 139629729);
    assert_eq!(looper_prog("ex_b2.txt"), 18216);
    println!("{:?}", looper_prog("prog.txt"));
}
