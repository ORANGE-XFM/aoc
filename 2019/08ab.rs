use std::fs::File;
use std::io::{Read,Result};

const W : usize = 25;
const H : usize = 6;

type Layer = [u8;W*H];

fn disp(layer : &Layer) {
	for i in 0..H {
		for j in 0..W {
			print!("{}",match layer[W*i+j] as char {'0' =>' ','1' => '#', '2'=>'x',_=>'.'} );
		}
		println!("");
	}
}

fn nb_of(layer: &Layer, chr: char) -> usize {
	return layer.iter().filter(|x| **x==chr as u8).count()
}

fn paint (layer :&Layer, layer_to: &mut Layer) {
	for i in 0..W*H {
		if layer_to[i] == '2' as u8 {
			layer_to[i]=layer[i];
		}
	}
}

fn main() -> Result<()> {
    let mut file = File::open("08.txt")?;
    let mut layer_to : Layer = ['2' as u8;W*H];

    let mut min_0 = 99999;
    let mut nb_min = 0;
    loop {    	
    	let mut layer : Layer = [0;W*H];
    	if file.read_exact(&mut layer).is_err() { break }
    	//disp( &layer);
    	let nb_0 =  nb_of(&layer, '0');
    	if nb_0<min_0 {
    		min_0=nb_0;
    		nb_min = nb_of(&layer,'1')*nb_of(&layer,'2');
    	}
    	paint(&layer, &mut layer_to);
    	//println!("{:?}",  nb_0);
    }
	println!("{:?} {}",min_0,nb_min);
	disp(&layer_to);
    Ok(())
}
