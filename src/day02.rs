pub fn part1(input: String) {
  let mut score = 0;
  for round in input.split('\n') {
    let p1 = round.chars().nth(0).unwrap() as u32 - 65;
    let p2 = round.chars().nth(2).unwrap() as u32 - 88;
    if (p1 + 1) % 3 == p2 {
      score += p2 + 7;
    } else if p1 == p2 {
      score += p2 + 4;
    } else {
      score += p2 + 1;
    }
  }
  println!("Score: {}", score);
}

pub fn part2(input: String) {
  let mut score = 0;
  for round in input.split('\n') {
    let p1 = round.chars().nth(0).unwrap() as i32 - 65;
    let result = round.chars().nth(2).unwrap() as i32 - 87;
    if result == 1 {
      // Should lose
      score += ((p1 + 2) % 3) + 1;
    } else if result == 2 {
      // Should tie
      score += p1 + 4;
    } else {
      // Should win
      score += ((p1 + 1) % 3) + 7;
    }
  }
  println!("Score: {}", score)
}
