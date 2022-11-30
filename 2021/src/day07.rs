
fn calculate_fuel(pos: &Vec<i32>, target: i32, part_two: bool) -> i64 {
    let mut fuel: i64 = 0;
    for num in pos.iter() {
        let diff = (target as i32 - num).abs();
        if part_two {
            fuel += (((diff*diff) + diff)/2) as i64;
        } else {
            fuel += diff as i64;
        }
    }
    return fuel;
}

pub fn part1(input: String) {
    let mut pos: Vec<i32> = Vec::new();
    for line in input.split('\n') {
        pos = line.trim().split(",").map(|s| s.parse::<i32>().unwrap()).collect();
    }
    let max_val = pos.iter().max().unwrap();
    let mut fuel_cost = vec![0; *max_val as usize];
    for index in 0..*max_val {
        fuel_cost[index as usize] = calculate_fuel(&pos, index as i32, false);
    }
    let min_fuel = fuel_cost.iter().min().unwrap();
    println!("Fuel {}", min_fuel);
}

pub fn part2(input: String) {
    let mut pos: Vec<i32> = Vec::new();
    for line in input.split('\n') {
        pos = line.trim().split(",").map(|s| s.parse::<i32>().unwrap()).collect();
    }
    let max_val = pos.iter().max().unwrap();
    let mut fuel_cost = vec![0; *max_val as usize];
    for index in 0..*max_val {
        fuel_cost[index as usize] = calculate_fuel(&pos, index as i32, true);
    }
    let min_fuel = fuel_cost.iter().min().unwrap();
    println!("Fuel {}", min_fuel);
}
