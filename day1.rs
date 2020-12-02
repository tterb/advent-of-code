use std::fs::File;
use std::io::{BufRead, BufReader};
use std::collections::HashMap;


fn main() {
    let target = 2020;
    let mut values = HashMap::new();
    let reader = BufReader::new(File::open("input/input1.txt").unwrap());
    for line in reader.lines() {
        let line = line.unwrap();
        let num = line.parse::<i32>().unwrap();
        match values.get(&num) {
            Some(&val) => println!("{} * {} = {}", num, val, num * val),
            _ => {
                values.insert(target - num, num);
            }
        }
    }
}