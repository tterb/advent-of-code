

pub fn part2(input: String) {
    let mut valid_count = 0;
    for line in input.split('\n') {
        let arr = line.trim().split(" ").collect::<Vec<&str>>();
        // Get indices, required charater, and password from line
        let indices = arr[0].split("-").map(|s| s.parse::<usize>().unwrap()).collect::<Vec<usize>>();
        let req = arr[1].replace(":", "");
        let pwd = arr[2].split("").collect::<Vec<&str>>();
        if pwd[indices[0]] == req && pwd[indices[1]] == req {
            valid_count += 1;
        }
    }
    println!("{} valid passwords", valid_count);
}

pub fn part1(input: String) {
    let mut valid_count = 0;
    for line in input.split('\n') {
        let arr = line.trim().split(" ").collect::<Vec<&str>>();
        // Get range, required charater, and password from line
        let rng = arr[0].split("-").map(|s| s.parse::<i32>().unwrap()).collect::<Vec<i32>>();
        let req = arr[1].replace(":", "");
        let pwd = arr[2].split("").collect::<Vec<&str>>();
        let mut char_count = 0;
        for c in pwd {
            if c == req {
                char_count += 1;
            }
        }
        if char_count >= rng[0] && char_count <= rng[1] {
            valid_count += 1;
        }
    }
    println!("{} valid passwords", valid_count);
}