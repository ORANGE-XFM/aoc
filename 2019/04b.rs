fn test(n: &u32) -> bool {
    let mut x = *n;

    let mut digit = 10; // high value to decrease
    let mut onesame = false;
    let mut pdigit = digit;
    let mut ppdigit = pdigit;
    let mut pppdigit;

    for _i in 0..7 { // one more for last iteration : first digit is never a zero
        pppdigit = ppdigit;
        ppdigit = pdigit;
        pdigit = digit;
        digit = x % 10; // units to milions
        //println!("{:?}",(digit,onesame,digit != pdigit && ppdigit == pdigit && pppdigit != ppdigit) );
        if digit != pdigit && ppdigit == pdigit && pppdigit != ppdigit {
            onesame = true
        } 
        if digit > pdigit {
            return false;
        };
        x /= 10;
    }
    return onesame;
}

fn main() {
    // tests
    for i in [ 135677, 112233, 111122, 223333].iter() {
        if !test(i) {println!("😱 {}", i)};
    }
    for i in [ 111111, 223450, 123789, 135679, 123444].iter() {
        if test(i) {println!("😱 {}", i)};
    }

    println!("total ok {}", (156218..=652527).filter(test).count());   

}
