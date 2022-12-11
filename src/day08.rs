fn get_forest(input: String) -> Vec<Vec<u32>> {
  return input
    .lines()
    .map(|line| {
      line
        .chars()
        .map(|c| c.to_digit(10).unwrap())
        .collect::<Vec<u32>>()
    })
    .collect::<Vec<Vec<u32>>>();
}

pub fn part1(input: String) {
  let forest = get_forest(input);
  let mut visible_trees = 0;
  for (i, col) in forest.iter().enumerate() {
    for (j, val) in col.iter().enumerate() {
      if (i == 0 || i == forest.len() - 1) || (j == 0 || j == col.len() - 1) {
        visible_trees += 1;
      } else {
        let is_visible: bool = forest[0..i].iter().all(|k| val > &k[j])
          || forest[i + 1..].iter().all(|k| val > &k[j])
          || forest[i][0..j].iter().all(|k| val > &k)
          || forest[i][j + 1..].iter().all(|k| val > &k);
        if is_visible {
          visible_trees += 1;
        }
      }
    }
  }
  println!("Visible trees: {}", visible_trees);
}

fn trees_in_view(tree: &u32, range: Vec<&u32>, reverse: bool) -> u32 {
  let mut trees = range.clone();
  if reverse {
    trees.reverse();
  }
  let mut count = 0;
  for val in &trees {
    count += 1;
    if tree <= val {
      break;
    }
  }
  return count;
}

pub fn part2(input: String) {
  let forest = get_forest(input);
  let mut highest_score = 0;
  for (i, col) in forest.iter().enumerate() {
    for (j, val) in col.iter().enumerate() {
      if (i != 0 && i != forest.len() - 1) && (j != 0 && j != col.len() - 1) {
        let scenic_score = trees_in_view(
          val,
          forest[0..i]
            .iter()
            .map(|tree| &tree[j])
            .collect::<Vec<&u32>>(),
          true,
        ) * trees_in_view(
          val,
          forest[i + 1..]
            .iter()
            .map(|tree| &tree[j])
            .collect::<Vec<&u32>>(),
          false,
        ) * trees_in_view(
          val,
          forest[i][0..j]
            .iter()
            .map(|tree| tree)
            .collect::<Vec<&u32>>(),
          true,
        ) * trees_in_view(
          val,
          forest[i][j + 1..]
            .iter()
            .map(|tree| tree)
            .collect::<Vec<&u32>>(),
          false,
        );
        if scenic_score > highest_score {
          highest_score = scenic_score;
        }
      }
    }
  }
  println!("Highest scenic score: {}", highest_score);
}
