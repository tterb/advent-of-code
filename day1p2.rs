use std::fs::File;
use std::io::{BufRead, BufReader};
use std::collections::HashSet;


fn main() {
    let target = 2020;
    let reader = BufReader::new(File::open("input/input1.txt").unwrap());
    let mut vals = vec![];
    for line in reader.lines() {
        vals.push(line
            .unwrap()
            .trim()
            .parse::<i32>().unwrap());
    }
    
    for i in 0..vals.len()-1 {
        let mut set = HashSet::new();
        let current = target - vals[i];
        for j in i+1..vals.len() {
            let target = current - vals[j];
            if set.contains(&target) {
                println!("{} * {} * {} = {}", vals[j], vals[j], target, vals[j] * vals[j] * target);
                return;
            }
            set.insert(vals[j]);
        }
    }
}