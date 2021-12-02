use std::collections::HashMap;
use std::collections::HashSet;


pub fn part2(input: String) {
    let target = 2020;
    let mut values = vec![];
    for line in input.split('\n') {
        values.push(line
            .trim()
            .parse::<i32>().unwrap());
    }
    for i in 0..values.len()-1 {
        let mut set = HashSet::new();
        let current = target - values[i];
        for j in i+1..values.len() {
            let target = current - values[j];
            if set.contains(&target) {
                println!("{} * {} * {} = {}", values[j], values[j], target, values[j] * values[j] * target);
                return;
            }
            set.insert(values[j]);
        }
    }
}

pub fn part1(input: String) {
    let target = 2020;
    let mut values = HashMap::new();
    for line in input.split('\n') {
        let num = line.parse::<i32>().unwrap();
        match values.get(&num) {
            Some(&val) => println!("{} * {} = {}", num, val, num * val),
            _ => {
                values.insert(target - num, num);
            }
        }
    }
}