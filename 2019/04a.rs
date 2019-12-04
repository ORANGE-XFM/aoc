fn test(n: &u32) -> bool {
    let mut x = *n;
    let mut digit = 10; // high value to decrease
    let mut onesame = false;

    for _i in 0..6 {
        let prevdigit = digit;
        digit = x % 10; // units to milions
        if digit == prevdigit {
            onesame = true
        } else if digit > prevdigit {
            return false;
        };
        x /= 10;
    }
    return onesame;
}

fn main() {
    let tests = [(111111,true), (223450,false), (123789,false), (135679,false), (135677,true)];
    for (i,r) in tests.iter() {
        if test(i) != *r {println!("😱 {}", i)};
    }

    println!("total ok {}", (156218..=652527).filter(test).count());   
}
