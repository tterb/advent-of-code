use std::collections::HashMap;
use std::collections::HashSet;

fn distance(head: (i32, i32), tail: (i32, i32)) -> i32 {
  let d = i32::pow(tail.0 - head.0, 2) + i32::pow(tail.1 - head.1, 2);
  return f32::sqrt(d as f32).round() as i32;
}

fn get_direction_step(dir: &str) -> (i32, i32) {
  let direction_steps: HashMap<&str, (i32, i32)> =
    [("L", (-1, 0)), ("R", (1, 0)), ("U", (0, 1)), ("D", (0, -1))]
      .iter()
      .cloned()
      .collect();
  return direction_steps[dir];
}

fn print_map(head: (i32, i32), tail: Vec<(i32, i32)>, limit: usize) {
  let mut board: Vec<Vec<String>> = vec![vec![".".to_string(); limit + 1]; limit + 1];
  for (i, t) in tail.iter().enumerate() {
    board[t.1 as usize][t.0 as usize] = i.to_string();
  }
  board[head.1 as usize][head.0 as usize] = "H".to_string();
  for i in board {
    println!(
      "{:?}",
      i.iter()
        .map(|s| s.to_owned().to_string() + " ")
        .collect::<String>()
    );
  }
  println!();
}

fn print_snake_map(snake: &Vec<(i32, i32)>, limit: usize) {
  let mut board: Vec<Vec<String>> = vec![vec![".".to_string(); limit + 1]; limit + 1];
  for (i, t) in snake.iter().enumerate() {
    if i == 0 {
      board[t.1 as usize][t.0 as usize] = "H".to_string();
    } else {
      board[t.1 as usize][t.0 as usize] = i.to_string();
    }
  }
  for i in board {
    println!(
      "{:?}",
      i.iter()
        .map(|s| s.to_owned().to_string() + " ")
        .collect::<String>()
    );
  }
  println!();
}

pub fn part2(input: String) {
  let snake_length = 10;
  let mut snake: Vec<(i32, i32)> = vec![(12, 12); snake_length];
  let mut tail_pos: HashSet<(i32, i32)> = HashSet::new();

  let instr = input
    .lines()
    .map(|s| {
      let chars: Vec<&str> = s.split(' ').collect();
      let step: (i32, i32) = get_direction_step(chars[0]);
      let count = chars[1].parse::<i32>().unwrap();
      return (step, count);
    })
    .collect::<Vec<((i32, i32), i32)>>();

  for (step, count) in &instr {
    let mut i = 0;
    // print_snake_map(&snake, 26);
    while i < *count {
      snake[0].0 += step.0;
      snake[0].1 += step.1;
      for (i, t) in &mut snake.clone().iter().enumerate() {
        if i > 0 {
          let s1 = snake[i - 1].clone();
          let mut s2 = t.clone();
          if distance(s1, s2) > 1 {
            if s1.0 > s2.0 {
              s2.0 += 1;
            } else if s1.0 < s2.0 {
              s2.0 -= 1;
            }
            if s1.1 > s2.1 {
              s2.1 += 1;
            } else if s1.1 < s2.1 {
              s2.1 -= 1;
            }
            snake[i] = s2;
          }
          if i == snake.len() - 1 {
            tail_pos.insert(s2);
          }
        }
      }
      i += 1;
    }
  }
  println!("Tail positions: {}", tail_pos.len());
}

pub fn part1(input: String) {
  let mut head: (i32, i32) = (0, 0);
  let mut tail: (i32, i32) = (0, 0);
  let mut tail_pos: HashSet<(i32, i32)> = HashSet::new();
  tail_pos.insert(tail);
  let instr = input
    .lines()
    .map(|s| {
      let chars: Vec<&str> = s.split(' ').collect();
      let step: (i32, i32) = get_direction_step(chars[0]);
      let count = chars[1].parse::<i32>().unwrap();
      return (step, count);
    })
    .collect::<Vec<((i32, i32), i32)>>();

  for (step, count) in &instr {
    let mut i = 0;
    while i < *count {
      // print_map(head, tail, 50);
      head.0 += step.0;
      head.1 += step.1;
      if distance(head, tail) > 1 {
        if head.0 > tail.0 {
          tail.0 += 1;
        } else if head.0 < tail.0 {
          tail.0 -= 1;
        }
        if head.1 > tail.1 {
          tail.1 += 1;
        } else if head.1 < tail.1 {
          tail.1 -= 1;
        }
        tail_pos.insert(tail);
      }
      i += 1;
    }
  }
  println!("Tail positions: {}", tail_pos.len());
}
