use pathfinding::prelude::bfs;

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Pos(usize, usize);

impl Pos {
  fn neighbors(&self, height_map: &Vec<Vec<i32>>) -> Vec<Pos> {
    let &Pos(x, y) = self;
    let mut options: Vec<Pos> = vec![];
    if x > 0 && (height_map[x - 1][y] <= height_map[x][y] + 1) {
      options.push(Pos(x - 1, y));
    }
    if y > 0 && (height_map[x][y - 1] <= height_map[x][y] + 1) {
      options.push(Pos(x, y - 1));
    }
    if x < height_map.len() - 1 && (height_map[x + 1][y] <= height_map[x][y] + 1) {
      options.push(Pos(x + 1, y));
    }
    if y < height_map[x].len() - 1 && (height_map[x][y + 1] <= height_map[x][y] + 1) {
      options.push(Pos(x, y + 1));
    }
    return options;
  }
}

pub fn part1(input: String) {
  let mut start: Pos = Pos(0, 0);
  let mut dest: Pos = Pos(0, 0);
  let mut height_map: Vec<Vec<i32>> = input
    .lines()
    .map(|l| l.trim().chars().map(|c| c as i32).collect())
    .collect();
  for (i, col) in height_map.clone().iter().enumerate() {
    for (j, val) in col.iter().enumerate() {
      if *val == 83 {
        start = Pos(i, j);
        height_map[i][j] = 97;
      } else if *val == 69 {
        dest = Pos(i, j);
        height_map[i][j] = 122;
      }
    }
  }
  let result = bfs(&start, |p| p.neighbors(&height_map), |p| *p == dest).unwrap_or(vec![Pos(0, 0)]);
  println!("Result ({}): {:?}", result.len() - 1, result);
}

pub fn part2(input: String) {
  let mut starts: Vec<Pos> = vec![];
  let mut dest: Pos = Pos(0, 0);
  let mut height_map: Vec<Vec<i32>> = input
    .lines()
    .map(|l| l.trim().chars().map(|c| c as i32).collect())
    .collect();
  for (i, col) in height_map.clone().iter().enumerate() {
    for (j, val) in col.iter().enumerate() {
      if *val == 97 {
        starts.push(Pos(i, j));
      } else if *val == 83 {
        starts.push(Pos(i, j));
        height_map[i][j] = 97;
      } else if *val == 69 {
        dest = Pos(i, j);
        height_map[i][j] = 122;
      }
    }
  }
  let mut paths: Vec<usize> = vec![];
  for start in starts {
    let path = bfs(&start, |p| p.neighbors(&height_map), |p| *p == dest).unwrap_or(vec![]);
    if path.len() > 0 {
      paths.push(path.len() - 1);
    }
  }
  println!("Result: {}", paths.iter().min().unwrap());
}
