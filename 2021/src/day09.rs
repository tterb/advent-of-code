
pub fn part1(input: String) {
    let mut map = vec![vec![0; 0]; 0];
    for line in input.split('\n') {
        let row = line.trim().chars().map(|c| c.to_digit(10).unwrap() as i32).collect::<Vec<i32>>();
        map.push(row);
    }
    for r in map.clone() {
        println!("{:?}", r);
    }
    let mut total_risk = 0;
    let row_count = map.len();
    for i in 0..row_count {
        for j in 0..map[i].len() {
            let current = map[i][j];
            let mut is_low = true;
            let mut adj = vec![0; 0];
            if i > 0 {
                adj.push(map[i-1][j]);
            }
            if i < map.len()-1 {
                adj.push(map[i+1][j]);
            }
            if j > 0 {
                adj.push(map[i][j-1]);
            }
            if j < map[i].len()-1 {
                adj.push(map[i][j+1]);
            }
            for a in adj {
                if a <= current {
                    is_low = false;
                }
            }
            if is_low {
                println!("Low point ({}, {}): {}", i, j, current);
                total_risk += current+1;
            }
        }
    }
    println!("Total risk: {}", total_risk);
}