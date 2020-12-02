use std::fs::File;
use std::io::{BufRead, BufReader};


fn main() {
    let reader = BufReader::new(File::open("input/input2.txt").unwrap());
    let mut valid_count = 0;
    for ln in reader.lines() {
        let line = ln.unwrap();
        let arr = line.trim().split(" ").collect::<Vec<&str>>();
        // Get indices, required charater, and password from line
        let indices = arr[0].split("-").map(|s| s.parse::<usize>().unwrap()).collect::<Vec<usize>>();
        let req = arr[1].replace(":", "");
        let pwd = arr[2].split("").collect::<Vec<&str>>();
        if pwd[indices[0]] == req && pwd[indices[1]] == req {
            valid_count += 1;
        }
    }
    println!("{} valid passwords", valid_count);
}