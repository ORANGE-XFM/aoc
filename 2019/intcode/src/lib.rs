use std::fs::read_to_string;

pub struct Program {
    p: Vec<i32>,
    ip: usize,
    input: Vec<i32>,
}

const TRACE: bool = false;

impl Program {
    fn arg(&self, n: u32) -> i32 {
        let mode = (self.p[self.ip] / 10_i32.pow(n + 1)) % 10;
        let x = match mode {
            0 => self.p[self.p[self.ip + n as usize] as usize],
            1 => self.p[self.ip + n as usize],
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
    fn set_at(&mut self, ofs: usize, v: i32) {
        let k = self.p[self.ip + ofs] as usize;
        self.p[k] = v;
        if TRACE {
            println!("   SET {} to {}", k, v)
        }
    }

    pub fn input(&mut self, n: i32) {
        self.input.push(n);
    }

    pub fn new(filename: &str) -> Program {
        println!("reading {:?}", filename);
        let line0 = read_to_string(filename).expect("Something went wrong reading the file ");
        let line = line0.trim();

        let split_line = line.split(",");
        let pro_iter = split_line.map(|x| x.parse::<i32>().expect("cannot parse integer"));

        Program {
            p: pro_iter.collect(),
            ip: 0,
            input: vec![],
        }
    }
}

impl Clone for Program {
    fn clone(&self) -> Program {
        Program {
            p: self.p.clone(),
            ip: self.ip,
            input: self.input.clone(),
        }
    }
}

impl Iterator for Program {
    type Item = i32;

    fn next(&mut self) -> Option<i32> {
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
                        99 => "END",
                        _ => "WTF",
                    }
                );
            }
            match instr {
                1 => {
                    let v = self.arg(1) + self.arg(2);
                    self.set_at(3, v);
                }
                2 => {
                    let v = self.arg(1) * self.arg(2);
                    self.set_at(3, v);
                }
                3 => {
                    let v = self.input.remove(0);
                    self.set_at(1, v)
                }
                4 => output = Some(self.arg(1)),
                5 => {
                    if self.arg(1) != 0 {
                        self.ip = (self.arg(2) - 3) as usize
                    }
                } // Jump if True, -1 : will be incremented afterards
                6 => {
                    if self.arg(1) == 0 {
                        self.ip = (self.arg(2) - 3) as usize
                    }
                } // Jump if False
                7 => {
                    let v = if self.arg(1) < self.arg(2) { 1 } else { 0 };
                    self.set_at(3, v);
                } // LT
                8 => {
                    let v = if self.arg(1) == self.arg(2) { 1 } else { 0 };
                    self.set_at(3, v);
                } // EQ
                99 => return None,
                _ => {
                    println!("Error: unknown op {}", self.p[self.ip]);
                    return None;
                }
            }

            self.ip += 1 + match instr {
                99 => 0,
                1 | 2 | 7 | 8 => 3,
                3 | 4 => 1,
                5 | 6 => 2,
                _ => {
                    println!("wat");
                    0
                }
            };
            // something to output ?
            if output.is_some() {
                return output;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
