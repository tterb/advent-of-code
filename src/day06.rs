
pub fn part1(input: String) {
    let mut fish: Vec<i32> = Vec::new();
    let num_days = 80;
    for line in input.split('\n') {
        fish = line.trim().split(",").map(|s| s.parse::<i32>().unwrap()).collect();
    }
    for _ in 0..num_days {
        let mut new_fish = vec![0; 0];
        for f in &mut fish {
            if *f == 0 {
                new_fish.push(8);
                *f = 6;
            } else {
                *f -= 1;
            }
        }
        fish.append(&mut new_fish);
    }
    println!("{} fish after {} days", fish.len(), num_days);
}

pub fn part2(input: String) {
    let mut fish: Vec<i64> = vec![0; 9];
    let num_days = 256;
    for line in input.split('\n') {
        let inp: Vec<i32> = line.trim().split(",").map(|s| s.parse::<i32>().unwrap()).collect();
        for f in inp {
            fish[f as usize] += 1;
        }
    }
    for _ in 0..num_days {
        let new_fish = fish[0];
        fish.rotate_left(1);
        fish[6] += fish[8];
        fish[8] = new_fish;
    }
    println!("{} fish after {} days", fish.iter().sum::<i64>(), num_days);
}