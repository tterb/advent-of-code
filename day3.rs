use std::fs::File;
use std::io::{BufRead, BufReader};


fn main() {
    let reader = BufReader::new(File::open("input/input3.txt").unwrap());
    let slope = [3, 1];
    let mut x = 0;
    let mut tree_count = 0;
    for (y, ln) in reader.lines().enumerate() {
        if y < slope[1] {
            x = slope[0];
            continue;
        }
        let line = ln.unwrap();
        let row = line.trim().chars().collect::<Vec<char>>();
        let index = x % row.len();
        if row[index] == '#' {
            tree_count += 1;
        }
        x += &slope[0];
    }
    println!("{} trees hit", tree_count);
}
