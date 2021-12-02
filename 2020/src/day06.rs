use std::collections::HashSet;


pub fn part1(input: String) {
    let mut answer_count = 0;
    let mut group_answers: HashSet<char> = HashSet::new();
    for line in input.split('\n') {
        if line.len() > 0 {
            let answers: HashSet<char> = line.chars().collect::<Vec<char>>().into_iter().collect();
            group_answers = group_answers.union(&answers).map(|c| *c).collect();
        } else {
            answer_count += group_answers.len();
            group_answers.clear();
        }
    }
    answer_count += group_answers.len();
    println!("Answer count: {}", answer_count);
}

pub fn part2(input: String) {
    let mut answer_count = 0;
    let mut start_group = true;
    let mut group_answers: HashSet<char> = HashSet::new();
    for line in input.split('\n') {
        if line.len() > 0 {
            let answers: HashSet<char> = line.chars().collect::<Vec<char>>().into_iter().collect();
            if start_group {
                group_answers = answers;
                start_group = false;
            } else {
                group_answers = group_answers.intersection(&answers).map(|c| *c).collect();
            }
        } else {
            answer_count += group_answers.len();
            group_answers.clear();
            start_group = true;
            continue;
        }
    }
    answer_count += group_answers.len();
    println!("Shared-Answer count: {}", answer_count);
}
