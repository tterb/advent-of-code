// Days
pub mod day01;
pub mod day02;
pub mod day03;
pub mod day04;
pub mod day05;
pub mod day06;
pub mod day07;
pub mod day08;
pub mod day09;
pub mod day10;
pub mod day11;
pub mod day12;

pub fn noop(_inp: String) {}

pub type DayFn = fn(String);

pub fn get_day(day: u32) -> (DayFn, DayFn) {
  return match day {
    1 => (day01::part1, day01::part2),
    2 => (day02::part1, day02::part2),
    3 => (day03::part1, day03::part2),
    4 => (day04::part1, day04::part2),
    5 => (day05::part1, day05::part2),
    6 => (day06::part1, day06::part2),
    // 7 => (day07::part1, noop),
    8 => (day08::part1, day08::part2),
    9 => (day09::part1, day09::part2),
    10 => (day10::part1, day10::part2),
    // 11 => (day11::part1, noop),
    12 => (day12::part1, day12::part2),
    _ => {
      println!("Unknown day: {}", day);
      return (noop, noop);
    }
  };
}
