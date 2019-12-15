pub type Word = i64;

pub struct Program {
    p: Vec<Word>,
    ip: usize,
    base: Word,
    input: Vec<Word>,
}

const TRACE: bool = false;
const PROGSIZE : usize = 1000;

impl Program {
    pub fn poke(&mut self, pos:usize, val:Word) {
        self.p[pos]=val;
    }
    pub fn start_at(&mut self, ip:usize) {
        self.ip = ip;
    }

    fn arg_info(&self, n:u32) -> (Word, Word) {
        let mode = (self.p[self.ip] / 10_i64.pow(n + 1)) % 10;
        let val = self.p[self.ip + n as usize];
        (mode,val)
    }

    fn arg(&self, n: u32) -> Word {
        let (mode,val) = self.arg_info(n);

        let x = match mode {
            0 => self.p[ val as usize ],
            1 => val,
            2 => self.p[ (self.base + val) as usize],
            _ => {
                println!("unknown mode {}", mode);
                0
            }
        };
        if TRACE {
            println!("   arg {}/mode{}=>{}", n, mode, x);
        }
        x
    }

    // set v to the address at addr (indirect mode)
    fn set_at(&mut self, n: u32, v: Word) {
        let (mode,val) = self.arg_info(n);
        let k = match mode {
            0 => val,
            2 => (self.base + val),
            _ => { println!("unknown mode {}", mode); 0 }
        }  as usize;
        self.p[k] = v;
        if TRACE {
            println!("   SET {} to {}", k, v)
        }
    }

    pub fn input(&mut self, n: Word) {
        self.input.push(n);
    }

    pub fn new(prog_str :&str) -> Program {
        let line = prog_str.trim();

        let split_line = line.split(',');
        let mut pro_vec : Vec<Word> = split_line.map(|x| x.parse::<Word>().expect("cannot parse integer")).collect();
        pro_vec.extend_from_slice(&[0;PROGSIZE]);

        Program {
            p: pro_vec,
            ip: 0,
            input: vec![],
            base: 0
        }
    }
}

impl Clone for Program {
    fn clone(&self) -> Program {
        Program {
            p: self.p.clone(),
            ip: self.ip,
            input: self.input.clone(),
            base: self.base,
        }
    }
}

impl Iterator for Program {
    type Item = Word;

    fn next(&mut self) -> Option<Word> {
        let mut output = None;

        loop {
            let instr = self.p[self.ip] % 100;
            if TRACE {
                println!(
                    "{} op={} {:?}",
                    self.ip,
                    self.p[self.ip],
                    match instr {
                        1 => "ADD",
                        2 => "MUL",
                        3 => "INPUT",
                        4 => "OUTPUT",
                        5 => "JE",
                        6 => "JNE",
                        7 => "LT",
                        8 => "EQ",
                        9 => "BASE",
                        99 => "END",
                        _ => "WTF",
                    }
                );
            }
            match instr {
                // Add
                1 => {
                    let v = self.arg(1) + self.arg(2);
                    self.set_at(3, v);
                    self.ip += 4;
                }

                // MUL
                2 => {
                    let v = self.arg(1) * self.arg(2);
                    self.set_at(3, v);
                    self.ip += 4;
                }

                // input
                3 => {
                    let v = self.input.remove(0);
                    self.set_at(1, v);
                    self.ip += 2;
                }

                // output
                4 => {
                    output = Some(self.arg(1));
                    self.ip += 2;
                }

                // Jump if True
                5 => {
                    self.ip = if self.arg(1) != 0 { self.arg(2) as usize } else { self.ip + 3 }
                } 

                // Jump if False
                6 => {
                    self.ip = if self.arg(1) == 0 { self.arg(2) as usize } else { self.ip + 3 }
                } 

                // LT
                7 => {
                    let v = if self.arg(1) < self.arg(2) { 1 } else { 0 };
                    self.set_at(3, v);
                    self.ip += 4;
                } 

                // EQ
                8 => {
                    let v = if self.arg(1) == self.arg(2) { 1 } else { 0 };
                    self.set_at(3, v);
                    self.ip += 4;
                }

                // move relative base
                9 => {
                    self.base += self.arg(1);
                    self.ip += 2;
                }

                // END
                99 => return None,

                // WTF
                _ => {
                    println!("Error: unknown op {}", self.p[self.ip]);
                    return None;
                }
            }

            // something to output ?
            if output.is_some() {
                return output;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn test_day05_test() {
        let mut p = Program::new("3,9,1001,9,23,9,4,9,99,0");
        p.input(100);
        assert_eq!(p.next(),Some(123));
    }

    const PROG_TEST8 : &str = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99";

    #[test]
    fn test_day05_test8_a() {
        let mut p = Program::new(PROG_TEST8);
        p.input(7);
        let res: Vec<Word> = p.collect();
        assert_eq!(res, vec![999]);
    }

    #[test]
    fn test_day05_test8_b() {
        let mut p = Program::new(PROG_TEST8);
        p.input(8);
        let res: Vec<Word> = p.collect();
        assert_eq!(res, vec![1000]);
    }

    #[test]
    fn test_day05_test8_c() {
        let mut p = Program::new(PROG_TEST8);
        p.input(9);
        let res: Vec<Word> = p.collect();
        assert_eq!(res, vec![1001]);
    }

    #[test]
    fn test_09a_quine(){
        let prog = "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99";
        let p = Program::new(&prog);
        let res: Vec<Word> = p.collect();
        assert_eq!(res, vec![109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99]);
    }

    #[test]
    fn test_09a_overflow(){
        let mut prog = Program::new("1102,34915192,34915192,7,4,7,99,0");
        assert!(prog.next().expect("no output?") > 1000000000000000_i64); // 16-digit number
    }

    #[test]
    fn test_09a_bignum(){
        let mut prog = Program::new("104,1125899906842624,99");
        assert_eq!(prog.next().expect("no output?") , 1125899906842624_i64); // 16-digit number
    }
}