pub fn part1(input: String) {
  let mut x = 1;
  let mut cycle = 0;
  let mut cycle_sum: Vec<i32> = vec![0];
  let mut signal_strengths: i32 = 0;
  let commands = input
    .lines()
    .map(|s| s.trim().split(' ').collect::<Vec<&str>>())
    .collect::<Vec<Vec<&str>>>();
  while cycle < 220 {
    cycle_sum.push(0);
    if cycle < commands.len() - 1 {
      if commands[cycle].len() > 1 && commands[cycle][0].eq("addx") {
        let val = commands[cycle][1].parse::<i32>().unwrap();
        cycle_sum.push(val);
      }
    }
    x += cycle_sum[cycle];
    cycle += 1;
    if cycle == 20 || (cycle > 20 && (cycle - 20) % 40 == 0) {
      let signal = (cycle as i32) * x;
      signal_strengths += signal;
    }
  }
  println!("Signal Strength: {}", signal_strengths);
}

pub fn part2(input: String) {
  let mut x = 1;
  let mut cycle = 1;
  let mut cycle_sum: Vec<i32> = vec![0];
  let commands = input
    .lines()
    .map(|s| s.trim().split(' ').collect::<Vec<&str>>())
    .collect::<Vec<Vec<&str>>>();
  while cycle < 240 {
    cycle_sum.push(0);
    if cycle < commands.len() - 1 {
      if commands[cycle - 1].len() > 1 && commands[cycle - 1][0].eq("addx") {
        let val = commands[cycle - 1][1].parse::<i32>().unwrap();
        cycle_sum.push(val);
      }
    }
    x += cycle_sum[cycle - 1];
    let offset = cycle as i32 % 40;
    if x == offset || x + 1 == offset || x + 2 == offset {
      print!("{0: <1}", "*");
    } else {
      print!("{0: <1}", ".");
    }
    if cycle > 0 && cycle % 40 == 0 {
      println!();
    }
    cycle += 1;
  }
  println!();
}
