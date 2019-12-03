use std::fs::File;
use std::io::{Result, BufReader, BufRead};

fn parse_element(s:&str) -> (char, i32) {
    let firstchar = s.chars().next().unwrap();
    let n = s[1..].parse::<i32>().expect("cannot parse");
    return (firstchar,n)
}

type LineVec = Vec::<(i32,i32,i32)>;

fn parse_line(ln:&str) -> (LineVec,LineVec) { // returns hlines, vlines
    let mut hlines = LineVec::new(); // y,x1,x2
    let mut vlines = LineVec::new(); // x,y1,y2

    let directions_iterator = ln.split(",").map(parse_element); 

    let mut pos = (0,0);
    for (dir,nb) in directions_iterator {
        match dir {
            'U' => {
                let npos = pos.1 - nb;
                vlines.push((pos.0,npos,pos.1));
                pos.1 = npos;
            },
            'D' => {
                let npos = pos.1 + nb;
                vlines.push((pos.0,pos.1,npos));
                pos.1 = npos;

            },
            'L' => {
                let npos = pos.0 - nb;
                hlines.push((pos.1,npos,pos.0));
                pos.0 = npos;
            },
            'R' => {
                let npos = pos.0 + nb;
                hlines.push((pos.1,pos.0,npos));
                pos.0 = npos;
            },               
            _   => println!("unknown direction {}",dir),
        }
        println!("prg = {:?} pos = {:?}", (dir,nb),pos);
    }

    (hlines,vlines)
}


// find all crosses between two elements and get the smallest from origin
fn get_intersect(hlines: &LineVec, vlines:&LineVec) -> i32 {
    let mut smallest = 999999999;
    for (ay,ax1,ax2) in hlines {
        for (bx,by1,by2) in vlines {
            if ay>=by1 && ay<=by2 && bx>=ax1 && bx<=ax2 {
                let dist = bx.abs()+ay.abs();
                if dist!=0 && dist<smallest {
                    smallest = dist;
                    println!("smaller ? cross at {},{}, dist: {}",bx,ay,dist);
                }
            }
        }
    }
    smallest
}

fn main() -> Result<()> {
    let file = File::open("03.txt")?;
    let reader = BufReader::new(file); 

    let mut vl = Vec::new();

    for line in reader.lines() {
        let ln = line?;
        let (hlines,vlines) = parse_line(&ln);
        println!("hlines {:?}",hlines);
        println!("vlines {:?}",vlines);
        vl.push((hlines,vlines));
    }

    get_intersect(&vl[0].0,&vl[1].1);
    get_intersect(&vl[0].1,&vl[1].0);

    Ok(())
}
