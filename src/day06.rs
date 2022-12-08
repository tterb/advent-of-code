use std::collections::HashSet;

fn get_distinct_window_index(window_size: usize, data: Vec<char>) -> usize {
  for (i, window) in data.windows(window_size).enumerate() {
    let chars: HashSet<char> = window.iter().copied().collect();
    if chars.len() == window_size {
      return i + window_size;
    }
  }
  return 0;
}

pub fn part1(input: String) {
  let data = input.trim().chars().collect::<Vec<char>>();
  let marker_index = get_distinct_window_index(4, data);
  println!("Marker index: {}", marker_index);
}

pub fn part2(input: String) {
  let data = input.trim().chars().collect::<Vec<char>>();
  let marker_index = get_distinct_window_index(14, data);
  println!("Message Marker index: {}", marker_index);
}
