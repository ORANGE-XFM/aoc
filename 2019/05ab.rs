use std::fs::read_to_string;

const TRACE : bool = true;

fn run(p0: &Vec<i32>, input:i32) -> i32 {
    let mut p = p0.clone();
    let mut output=-12345678; // special case , if same value at the end output p[0]

    if TRACE {println!("start ---")}

    let mut ip = 0;
    fn get_op(p:&Vec<i32>, ip:usize, n:u32) -> i32{
        let mode = (p[ip] / 10_i32.pow(n+1)) % 10;
        let x = match mode {
            0 => p[p[ip+n as usize] as usize],
            1 => p[ip+n as usize],
            _ => {println!("unknown mode {}",mode);0},
        };
        if TRACE {println!("   oper {}/mode{}=>{}", n,mode,x);}
        x
    }

    // set v to the address at addr (indirect mode)
    fn set_at(p:& mut Vec<i32>, addr:usize, v:i32) {
        let k = p[addr as usize] as usize;
        p[k] = v;
        if TRACE {println!("   SET {} to {}",addr,v )}
    }

    'outer: loop {
        let instr = p[ip]%100;
        if TRACE {
            println!("{} op={} {:?}",ip,p[ip],match instr {1=>"ADD",2=>"MUL",3=>"INPUT",4=>"OUTPUT",5=>"JE",6=>"JNE",7=>"LT",8=>"EQ",99=>"END",_=>"WTF"});
        }
        match instr {
            1 => { let v = get_op(&p,ip,1) + get_op(&p,ip,2) ; set_at(&mut p,ip+3,v); },
            2 => { let v = get_op(&p,ip,1) * get_op(&p,ip,2) ; set_at(&mut p,ip+3,v); },
            3 => set_at( & mut p,ip+1, input),
            4 => output = get_op(&p,ip,1),
            5 => if get_op(&p,ip,1)!=0 {ip=(get_op(&p,ip,2)-3) as usize}, // Jump if True, -1 : will be incremented afterards
            6 => if get_op(&p,ip,1)==0 {ip=(get_op(&p,ip,2)-3) as usize}, // Jump if False
            7 => { let v = if get_op(&p,ip,1) < get_op(&p,ip,2)  {1} else {0}; set_at(&mut p,ip+3,v); },  // LT
            8 => { let v = if get_op(&p,ip,1) == get_op(&p,ip,2) {1} else {0}; set_at(&mut p,ip+3,v); },  // EQ
            99 => break 'outer,
            _ => { println!("Error: unknown op {}",p[ip]);return -1},
        }
        
        ip += 1+match instr {99=>0, 1|2|7|8=>3, 3|4=>1, 5|6=>2,_=>{println!("wat");0}}; // Jumps already changed
    }
    return if output==-12345678 {p[0]} else {output}; // special case for program 02 that should still work but cannot output 
}

fn parse_prog(filename:&str) -> Vec<i32> {
    let line0 = read_to_string(filename).expect("Something went wrong reading the file");
    let line = line0.trim();

    let split_line = line.split(",");
    let pro_iter = split_line.map(|x| x.parse::<i32>().expect("cannot parse"));
    pro_iter.collect()
}

fn main() {
    let mut pr02 = parse_prog("02.txt");
    pr02[1] = 1; pr02[2] = 12;
    if run(&pr02,0)!=490655 {println!("😱 02.txt failed !")}

    if run(&parse_prog("05_test.txt"),100) != 123 {println!("😱 test05 failed!")}
    let testpr8 = parse_prog("05_test8.txt");
    if run(&testpr8,7)!=999  {println!("😱 {:?}",7)};
    if run(&testpr8,8)!=1000 {println!("😱 {:?}",8)};
    if run(&testpr8,9)!=1001 {println!("😱 {:?}",9)};

    let pr5 = parse_prog("05.txt");
    println!("prog(1)= {}", run(&pr5, 1));
    println!("prog(5)= {}", run(&pr5, 5));

    let pr_c = parse_prog("05_example.txt");
    println!("prog_asm(3)={}",run(&pr_c,3));
    println!("prog_asm(7)={}",run(&pr_c,7));
}
