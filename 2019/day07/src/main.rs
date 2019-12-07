use intcode::Program;
use permutohedron::LexicalPermutation;

fn runmachine_a(pr: &Program, arg: Vec<i32>) -> Option<i32> {
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

fn get_maxprog(progfile: &str) -> i32 {
    let prog = Program::new(progfile);
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

fn runmachine_loop(pr: &Program, arg: Vec<i32>) -> i32 {
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

fn looper_prog(progfile: &str) -> i32 {
    let prog = Program::new(progfile);
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
    let mut p = Program::new("../05_test.txt");
    p.input(100);
    let res: Vec<i32> = p.collect();
    assert_eq!(res, vec![123]);

    let mut p = Program::new("../05_test8.txt");
    p.input(7);
    let res: Vec<i32> = p.collect();
    assert_eq!(res, vec![999]);

    let mut p = Program::new("../05_test8.txt");
    p.input(8);
    let res: Vec<i32> = p.collect();
    assert_eq!(res, vec![1000]);

    // several outputs
    let mut p = Program::new("../05.txt");
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
