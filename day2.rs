use std::fs::File;
use std::io::{BufRead, BufReader};


fn main() {
    let reader = BufReader::new(File::open("input/input2.txt").unwrap());
    let mut valid_count = 0;
    for ln in reader.lines() {
        let line = ln.unwrap();
        let arr = line.trim().split(" ").collect::<Vec<&str>>();
        // Get range, required charater, and password from line
        let rng = arr[0].split("-").map(|s| s.parse::<i32>().unwrap()).collect::<Vec<i32>>();
        let req = arr[1].replace(":", "");
        let pwd = arr[2].split("").collect::<Vec<&str>>();
        let mut char_count = 0;
        for c in pwd {
            if c == req {
                char_count += 1;
            }
        }
        if char_count >= rng[0] && char_count <= rng[1] {
            valid_count += 1;
        }
    }
    println!("{} valid passwords", valid_count);
}