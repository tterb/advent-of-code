

pub fn part2(input: String) {
    let slopes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]];
    let mut tree_count = vec![0; slopes.len()];
    for (y, line) in input.split('\n').enumerate() {
        let row = line.trim().chars().collect::<Vec<char>>();
        for (index, step) in slopes.iter().enumerate() {
            if y >= step[1] && y % step[1] == 0 {
                let x_coord = (y / step[1]) * step[0];
                if row[x_coord % row.len()] == '#' {
                    tree_count[index] += 1;
                }
            }
        }
    }
    let total_count: i64 = tree_count.iter().product();
    println!("{} trees hit", total_count);
}


pub fn part1(input: String) {
    let slope = [3, 1];
    let mut x = 0;
    let mut tree_count = 0;
    for (y, line) in input.split('\n').enumerate() {
        if y < slope[1] {
            x = slope[0];
            continue;
        }
        let row = line.trim().chars().collect::<Vec<char>>();
        let index = x % row.len();
        if row[index] == '#' {
            tree_count += 1;
        }
        x += &slope[0];
    }
    println!("{} trees hit", tree_count);
}
