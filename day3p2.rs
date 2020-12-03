use std::fs::File;
use std::io::{BufRead, BufReader};


fn main() {
    let reader = BufReader::new(File::open("input/input3.txt").unwrap());
    let slopes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]];
    let mut tree_count = vec![0; slopes.len()];
    for (y, ln) in reader.lines().enumerate() {
        let line = ln.unwrap();
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
    println!("{} total trees hit", total_count);
}
