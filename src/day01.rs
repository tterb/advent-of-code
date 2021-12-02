
pub fn part1(input: String) {
    let mut prev = i32::MAX;
    let mut increase_count = 0;
    for line in input.split('\n') {
        let num = line.parse::<i32>().unwrap();
        if prev < num {
            increase_count += 1
        }
        prev = num;
    }
    println!("{} increases", increase_count)
}

pub fn part2(input: String) {
    let mut values = vec![];
    let mut increase_count = 0;
    for (i, line) in input.split('\n').enumerate() {
        values.push(line
            .trim()
            .parse::<i32>().unwrap());
        if i >= 3 && values[i-3] < values[i] {
            increase_count += 1
        }
    }

    println!("{} sliding increases", increase_count)
}