use std::collections::HashSet;
use std::iter::FromIterator;

fn get_priority_value(c: char) -> i32 {
  if c.is_uppercase() {
    return c as i32 - 38;
  } else {
    return c as i32 - 96;
  };
}

pub fn part1(input: String) {
  let mut shared: Vec<char> = Vec::new();
  for line in input.split('\n') {
    let segments: Vec<HashSet<char>> = line
      .chars()
      .collect::<Vec<char>>()
      .chunks(line.len() / 2)
      .into_iter()
      .map(|c| HashSet::from_iter(c.to_vec()))
      .collect();

    shared.push(
      *segments[0]
        .intersection(&segments[1])
        .collect::<HashSet<&char>>()
        .drain()
        .next()
        .unwrap(),
    );
  }
  let total: i32 = shared.into_iter().map(|c| get_priority_value(c)).sum();
  println!("Priority: {}", total);
}

pub fn part2(input: String) {
  let mut badges: Vec<char> = Vec::new();
  let groups: Vec<Vec<HashSet<char>>> = Vec::from_iter(input.split('\n'))
    .chunks(3)
    .map(|c| {
      c.clone()
        .to_vec()
        .iter()
        .map(|s| HashSet::from_iter(s.chars()))
        .collect()
    })
    .collect();

  for group in groups {
    badges.push(
      *group[0]
        .intersection(&group[1])
        .map(|c| c.clone())
        .collect::<HashSet<char>>()
        .intersection(&group[2])
        .collect::<HashSet<&char>>()
        .drain()
        .next()
        .unwrap(),
    );
  }
  let total: i32 = badges.into_iter().map(|c| get_priority_value(c)).sum();
  println!("Priority: {}", total);
}
