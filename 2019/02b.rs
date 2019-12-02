use std::fs::read_to_string;
use std::io::Result;

fn run(p0: &Vec<i32>, a: i32, b: i32) -> i32 {
    let mut p = p0.clone();
    p[1] = a;
    p[2] = b;

    let mut i = 0;
    loop {
        if p[i] == 99 {
            break;
        }
        let a = p[i + 1] as usize;
        let b = p[i + 2] as usize;
        let c = p[i + 3] as usize;
        p[c] = if p[i] == 1 { p[a] + p[b] } else { p[a] * p[b] };
        i += 4
    }
    return p[0];
}

fn main() -> Result<()> {
    let line = read_to_string("02.txt").expect("Something went wrong reading the file");
    let split_line = line.split(",");
    let pro_iter = split_line.map(|x| x.parse::<i32>().expect("cannot parse"));
    let program: Vec<i32> = pro_iter.collect();

    'outer: for i in 0..99 {
        for j in 0..99 {
            let res = run(&program, i, j);
            if res == 19690720 {
                println!("prog({},{})= {}", i, j, res);
                break 'outer;
            }
        }
    }

    Ok(())
}
