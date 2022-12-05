use std::collections::HashSet;

pub fn part1(input: String) {
  let mut subsets = 0;
  for line in input.split('\n') {
    let mut sections: Vec<HashSet<i32>> = Vec::new();
    for range in line.split(',') {
      let bounds: Vec<i32> = range
        .split('-')
        .map(|s| s.parse::<i32>().unwrap())
        .collect();
      let mut assign: HashSet<i32> = HashSet::new();
      for n in bounds[0]..bounds[1] + 1 {
        assign.insert(n);
      }
      sections.push(assign);
    }
    if sections[0].is_subset(&sections[1]) || sections[0].is_superset(&sections[1]) {
      subsets += 1;
    }
  }
  println!("Subsets: {}", subsets);
}

pub fn part2(input: String) {
  let mut subsets = 0;
  for line in input.split('\n') {
    let mut sections: Vec<HashSet<i32>> = Vec::new();
    for range in line.split(',') {
      let bounds: Vec<i32> = range
        .split('-')
        .map(|s| s.parse::<i32>().unwrap())
        .collect();
      let mut assign: HashSet<i32> = HashSet::new();
      for n in bounds[0]..bounds[1] + 1 {
        assign.insert(n);
      }
      sections.push(assign);
    }
    let overlap: HashSet<_> = sections[0].intersection(&sections[1]).collect();
    if overlap.len() > 0 {
      subsets += 1;
    }
  }
  println!("Overlapping: {}", subsets);
}
