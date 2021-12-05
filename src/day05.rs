
fn read_lines(input: &str, max_dim: &mut i32) -> Vec<Vec<Vec<i32>>> {
    let mut lines = vec![vec![vec![0; 2]; 0]; 0];
    for (_, line) in input.split('\n').enumerate() {
        let mut coords: Vec<Vec<i32>> = Vec::new();
        let split_string = line.trim().split(" -> ").collect::<Vec<&str>>();
        for coord in split_string {
            let temp = coord.split(",").map(|s| s.parse().unwrap()).collect::<Vec<i32>>();
            for val in temp.iter() {
                if *max_dim < *val {
                    *max_dim = *val;
                }
            }
            coords.push(temp);
        }
        lines.push(coords);
    }
    return lines;
}

fn mark_line(map: &mut Vec<Vec<i32>>, line: &Vec<Vec<i32>>, overlap_count: &mut i32, horiz: bool) {
    let mut x1 = line[0][0];
    let mut y1 = line[0][1];
    let x2 = line[1][0];
    let y2 = line[1][1];
    let mut dx = (x2 - x1);
    let mut dy = (y2 - y1);
    if x1 != x2 && y1 != y2 && (!horiz && (x1 - y1).abs() != (x2 - y2).abs()) {
        return;
    }
    if dx != 0 {
        if dx < 0 {
            dx = -1
        } else {
            dx = 1;
        }
    }
    if dy != 0 {
        if dy < 0 {
            dy = -1
        } else {
            dy = 1;
        }
    }

    while x1 != x2 || y1 != y2 {
        if map[y1 as usize][x1 as usize] == 1 {
            *overlap_count += 1;
        }
        map[y1 as usize][x1 as usize] += 1;
        x1 += dx;
        y1 += dy;
    }
    if map[y1 as usize][x1 as usize] == 1 {
        *overlap_count += 1;
    }
    map[y1 as usize][x1 as usize] += 1;
}

pub fn part1(input: String) {
    let mut max_dim = 0;
    let lines = read_lines(&input, &mut max_dim);

    let mut overlap_count = 0;
    let bounds = (max_dim+1) as usize;
    let mut map = vec![vec![0; bounds]; bounds];
    for line in lines {
        mark_line(&mut map, &line, &mut overlap_count, false);
    }
    println!("Overlaps: {}", overlap_count);
}

pub fn part2(input: String) {
    let mut max_dim = 0;
    let lines = read_lines(&input, &mut max_dim);

    let mut overlap_count = 0;
    let bounds = (max_dim+1) as usize;
    let mut map = vec![vec![0; bounds]; bounds];
    for line in lines {
        mark_line(&mut map, &line, &mut overlap_count, true);
    }
    println!("Overlaps: {}", overlap_count);
}
