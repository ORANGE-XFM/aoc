use std::fs::File;
use std::io::{BufRead, BufReader, Result};

fn parse_element(s: &str) -> (bool, i32) {
    let firstchar = s.chars().next().unwrap();
    let nb = s[1..].parse::<i32>().expect("cannot parse");
    match firstchar {
        'U' => (true, -nb),
        'D' => (true, nb),
        'L' => (false, -nb),
        'R' => (false, nb),
        _ => {
            println!("unknown direction {}", firstchar);
            (false, 0)
        }
    }
}

type LineVec = Vec<(bool, i32)>; // is_vertical, dist

fn between(x: i32, a: i32, b: i32) -> bool {
    // x between a and b
    (a <= x && x <= b) || (a >= x && x >= b)
}

// find smallest crosses between two elements from origin
fn get_intersect(lines_a: &LineVec, lines_b: &LineVec) -> (i32, i32) {
    let mut ax = 0;
    let mut ay = 0;
    let mut smallest = 99999999;
    let mut smallest2 = 99999999;

    let mut dist_a = 0;
    for (dir_a, len_a) in lines_a {
        let mut bx = 0;
        let mut by = 0;
        let mut dist_b = 0;
        for (dir_b, len_b) in lines_b {
            if *dir_b != *dir_a {
                // -> not paralell (assumes not colinear)
                if *dir_a {
                    // A is vertical line, B should be horizontal
                    if between(ax, bx, bx + len_b) && between(by, ay, ay + len_a) {
                        let dist = ax.abs() + by.abs();
                        if dist != 0 && dist < smallest {
                            smallest = dist;
                        }

                        let dist2 = dist_a + (by - ay).abs() + dist_b + (ax - bx).abs();
                        if dist2 != 0 && dist2 < smallest2 {
                            smallest2 = dist2;
                        }
                    }
                } else {
                    if between(ay, by, by + len_b) && between(bx, ax, ax + len_a) {
                        let dist = bx.abs() + ay.abs();
                        if dist != 0 && dist < smallest {
                            smallest = dist;
                        }

                        let dist2 = dist_a + (bx - ax).abs() + dist_b + (ay - by).abs();
                        if dist2 != 0 && dist2 < smallest2 {
                            smallest2 = dist2;
                        }
                    }
                }
            }
            if *dir_b {
                by += len_b
            } else {
                bx += len_b
            };
            dist_b += len_b.abs();
        }
        if *dir_a {
            ay += len_a
        } else {
            ax += len_a
        };
        dist_a += len_a.abs();
    }
    (smallest, smallest2)
}

fn main() -> Result<()> {
    let file = File::open("03.txt")?;
    let reader = BufReader::new(file);

    let mut vl = Vec::new();

    for line in reader.lines() {
        let lines: LineVec = line?.split(",").map(parse_element).collect();
        //println!("lines {:?}",lines);
        vl.push(lines);
    }

    println!("result: {:?}", get_intersect(&vl[0], &vl[1]));

    Ok(())
}
