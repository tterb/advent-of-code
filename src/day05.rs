use std::str;

fn follow_instruction(stacks: &mut Vec<Vec<&str>>, instr: &str, reverse: bool) {
  let op = instr
    .split(' ')
    .filter_map(|a| a.parse::<usize>().ok())
    .map(|i| i - 1)
    .collect::<Vec<usize>>();
  let mut shift: Vec<&str> = stacks[op[1]].drain(0..op[0] + 1).collect();
  if reverse {
    shift = shift.into_iter().rev().collect()
  }
  for (i, s) in shift.iter().enumerate() {
    stacks[op[2]].insert(i, s);
  }
}

fn get_top_stacks(stacks: Vec<Vec<&str>>) -> String {
  return stacks
    .iter()
    .filter_map(|s| None.unwrap_or(s.first().cloned()))
    .map(|s| s.chars().collect::<Vec<char>>()[1])
    .collect::<String>();
}

pub fn part1(input: String) {
  let mut stacks: Vec<Vec<&str>> = std::iter::repeat(vec![]).take(10).collect::<Vec<_>>();
  for line in input.split('\n') {
    if line.contains("move") {
      follow_instruction(&mut stacks, line, true);
    } else if line.contains("[") {
      let parts = line
        .as_bytes()
        .chunks(4)
        .map(|s| str::from_utf8(s).unwrap().trim())
        .collect::<Vec<&str>>();
      for (i, p) in parts.iter().enumerate() {
        if p.len() > 0 {
          stacks[i].push(*p)
        }
      }
    }
  }
  println!("Top stacks: {:?}", get_top_stacks(stacks));
}

pub fn part2(input: String) {
  let mut stacks: Vec<Vec<&str>> = std::iter::repeat(vec![]).take(10).collect::<Vec<_>>();
  for line in input.split('\n') {
    if line.contains("move") {
      follow_instruction(&mut stacks, line, false);
    } else if line.contains("[") {
      let parts = line
        .as_bytes()
        .chunks(4)
        .map(|s| str::from_utf8(s).unwrap().trim())
        .collect::<Vec<&str>>();
      for (i, p) in parts.iter().enumerate() {
        if p.len() > 0 {
          stacks[i].push(*p)
        }
      }
    }
  }
  println!("Top stacks: {:?}", get_top_stacks(stacks));
}
