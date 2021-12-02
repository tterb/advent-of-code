
pub fn part1(input: String) {
    let mut horiz = 0;
    let mut depth = 0;
    for line in input.split('\n') {
        let inp = line.trim().split(" ").collect::<Vec<&str>>();
        let dir = inp[0];
        let val = inp[1].parse::<i32>().unwrap();
        if dir == "up" {
            depth -= val;
        } else if dir == "down" {
            depth += val;
        } else if dir == "forward" {
            horiz += val;
        }
    }
    println!("(horizontal, depth): ({}, {})", horiz, depth);
    println!("Total: {}", horiz * depth);
}

pub fn part2(input: String) {
    let mut horiz = 0;
    let mut depth = 0;
    let mut aim = 0;
    for line in input.split('\n') {
        let inp = line.trim().split(" ").collect::<Vec<&str>>();
        let dir = inp[0];
        let val = inp[1].parse::<i32>().unwrap();
        if dir == "up" {
            aim -= val;
        } else if dir == "down" {
            aim += val;
        } else if dir == "forward" {
            horiz += val;
            depth += aim * val
        }
    }
    println!("(horizontal, depth): ({}, {})", horiz, depth);
    println!("Total: {}", horiz * depth);
}