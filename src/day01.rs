pub fn part1(input: String) {
  let mut max_calories = 0;
  let mut total = 0;
  for val in input.split('\n') {
    if val.len() == 0 {
      if total > max_calories {
        max_calories = total;
      }
      total = 0;
    } else {
      total += val.parse::<i64>().unwrap();
    }
  }
  println!("Max calories: {}", max_calories);
}

pub fn part2(input: String) {
  let mut total = 0;
  let mut values = vec![];
  for val in input.split('\n') {
    if val.len() == 0 {
      values.push(total);
      total = 0;
    } else {
      total += val.parse::<i64>().unwrap();
    }
  }
  values.push(total);
  values.sort();
  let top_calories: i64 = values[values.len()-3 .. values.len()].iter().sum();
  println!("Sum of top calories: {}", top_calories)
}