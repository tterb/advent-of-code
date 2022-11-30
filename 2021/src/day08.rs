use std::collections::HashMap;
use std::collections::HashSet;

// let mut segments = vec![
    //     //   0, 1, 2, 3, 4, 5, 6, 7, 8, 9
    //     vec![6, 2, 4, 4, 3, 4, 5, 3, 6, 5], // 0
    //     vec![2, 2, 1, 2, 2, 1, 1, 2, 2, 2], // 1
    //     vec![4, 1, 5, 4, 2, 3, 4, 2, 5, 4], // 2
    //     vec![4, 2, 4, 5, 3, 4, 4, 3, 5, 5], // 3
    //     vec![3, 2, 2, 3, 4, 3, 3, 2, 4, 4], // 4
    //     vec![4, 1, 3, 4, 3, 5, 5, 2, 5, 5], // 5
    //     vec![5, 1, 4, 3, 3, 5, 6, 2, 6, 5], // 6
    //     vec![3, 2, ], // 7
    // ];
    // let mut segment_vec = vec![
    //     vec!['a', 'b', 'c', 'e', 'f', 'g'], // 0
    //     vec!['c', 'f'], // 1
    //     vec!['a', 'c', 'd', 'e', 'g'], // 2
    //     vec!['a', 'c', 'd', 'f', 'g'], // 3
    //     vec!['b', 'c', 'd', 'f'], // 4
    //     vec!['a', 'b', 'd', 'f', 'g'], // 5
    //     vec!['a', 'b', 'd', 'e', 'f', 'g'], // 6
    //     vec!['a', 'c', 'f'], // 7
    //     vec!['a', 'b', 'c', 'd', 'e', 'f', 'g'], // 8
    //     vec!['a', 'b', 'c', 'd', 'f', 'g'], // 9
    // ];

pub fn part2(input: String) {
    let mut digits = HashMap::from([
        (0, HashSet::from(['a', 'b', 'c', 'e', 'f', 'g'])),
        (1, HashSet::from(['c', 'f'])),
        (2, HashSet::from(['a', 'c', 'd', 'e', 'g'])),
        (3, HashSet::from(['a', 'c', 'd', 'f', 'g'])),
        (4, HashSet::from(['b', 'c', 'd', 'f'])),
        (5, HashSet::from(['a', 'b', 'd', 'f', 'g'])),
        (6, HashSet::from(['a', 'b', 'd', 'e', 'f', 'g'])),
        (7, HashSet::from(['a', 'c', 'f'])),
        (8, HashSet::from(['a', 'b', 'c', 'd', 'e', 'f', 'g'])),
        (9, HashSet::from(['a', 'b', 'c', 'd', 'f', 'g']))
    ]);
    // println!("{:?}", digits);
    let mut chars = vec!['a', 'b', 'c', 'd', 'e', 'f', 'g'];
    let mut letters: HashSet<char> = HashSet::from(['a', 'b', 'c', 'd', 'e', 'f', 'g']);
    let mut exclude = HashMap::from([
        (0, HashSet::from(['d'])),
        (1, HashSet::from(['a', 'b', 'd', 'e', 'g'])),
        (2, HashSet::from(['b', 'f'])),
        (3, HashSet::from(['b', 'e'])),
        (4, HashSet::from(['a', 'e', 'g'])),
        (5, HashSet::from(['c', 'e'])),
        (6, HashSet::from(['c'])),
        (7, HashSet::from(['b', 'd', 'e', 'g'])),
        (8, HashSet::from([])),
        (9, HashSet::from(['e']))
    ]);
    let mut lengths = HashMap::from([(0, 6), (1, 2), (2, 5), (3, 5), (4, 4), (5, 5), (6, 6), (7, 3), (8, 7), (9, 6)]);
    let mut unique_lengths = HashMap::from([(2, 1), (4, 4), (3, 7), (7, 8)]);
    let mut codes: HashMap<i32, HashSet<char>> = HashMap::new();
    let eight_seg = codes.entry(8).or_insert(HashSet::new());
    *eight_seg = letters.clone();
    let mut options: HashMap<char, HashSet<char>> = HashMap::new();
    for c in chars {
        options.insert(c, letters.clone());
    }

    let mut groups: Vec<HashSet<char>> = Vec::new();
    let mut output: Vec<HashSet<char>> = Vec::new();
    // let mut pairs: Vec<Vec<&str>> = Vec::new();
    for line in input.split('\n') {
        let pair = line.trim().split(" | ").collect::<Vec<&str>>();
        for (i, p) in pair.iter().enumerate() {
            let temp = p.split(" ").collect::<Vec<&str>>();
            for c in temp {
                let char_vec: HashSet<char> = c.chars().collect();
                if i == 1 {
                    output.push(char_vec.clone());
                }
                groups.push(char_vec);
            }
        }
    }

    let mut has_changes = true;
    while has_changes {
        for group in &groups {
            if group.len() == 2 {
                // for c in &group {
                //     for i in [0, 1, 4, 7, 8, 9] {
                //         let segs = codes.entry[i].or_insert(HashSet::new());
                //         segs.insert(*c);
                //     }
                // }
                for i in [0, 1, 3, 4, 7, 8, 9] {
                    let segs = codes.entry(i).or_insert(HashSet::new());
                    // let group_set = HashSet::from(group.clone());
                    segs.extend(HashSet::from(group.clone()));
                }
            } else if group.len() == 3 {
                for i in [0, 3, 7, 8, 9] {
                    let segs = codes.entry(i).or_insert(HashSet::new());
                    segs.extend(HashSet::from(group.clone()));
                }
                // let one_seg = codes.entry(1).or_insert(HashSet::new());
                let default = HashSet::new();
                let one_seg = codes.get(&1).unwrap_or(&default);
                if one_seg.len() == 2 {
                    // let seven_seg = &*codes.entry(7).or_insert(HashSet::new()).clone();
                    let seven_seg = &codes.get(&7).unwrap().clone();
                    // let diff_seg = seven_seg.difference(&one_seg).collect();
                    let diff_seg = &*seven_seg - one_seg;
                    // *opts = &*opts - temp;
                    for i in [2, 5, 6, 8, 9] {
                        let segs = codes.entry(i).or_insert(HashSet::new());
                        segs.extend(&diff_seg);
                        println!("Seg: {:?}", segs);
                        has_changes = false;
                    }
                }
            } else if group.len() == 4 {
                for i in [4, 8, 9] {
                    let segs = codes.entry(i).or_insert(HashSet::new());
                    segs.extend(HashSet::from(group.clone()));
                }
                let default = HashSet::new();
                let one_seg = codes.get(&1).unwrap_or(&default).clone();
                if one_seg.len() == 2 {
                    let temp = codes.get(&4).unwrap();
                    let four_seg = &temp.clone();
                    // let four_seg = &codes.get(&7).unwrap().clone();
                    // let diff_seg = seven_seg.difference(&one_seg).collect();
                    let diff_seg = &*four_seg - &one_seg;
                    // *opts = &*opts - temp;
                    for i in [2, 5, 6, 8, 9] {
                        let segs = codes.entry(i).or_insert(HashSet::new());
                        if i != 2 {
                            segs.extend(&*four_seg - &one_seg);
                        }
                        if i == 2 || i == 6 {
                            let diff = &letters.clone() - four_seg;
                            println!("4: {:?}", four_seg);
                            println!("8: {:?}", letters);
                            println!("8 - 4: {:?}", diff);
                            segs.extend(&diff);
                        }
                        println!("Seg: {:?}", segs);
                    }
                    // diff_seg = &letters.clone() - one_seg;
                    // for i in [2, 6] {
                    //     let segs = codes.entry(i).or_insert(HashSet::new());
                    //     segs.extend(&diff_seg);
                    //     println!("Seg: {:?}", segs);
                    // }
                }
            }
        }
    }
    for i in codes {
        println!("{:?}", i);
    }

    // for group in groups {
    //     // println!("Group: {:?}", group);
    //     if group.len() == 2 || group.len() == 3 {
    //         println!("Group: {:?}", group);
    //         for c in &letters {
    //             // println!("{}: {:?}", c, group);
    //             let index = unique_lengths.entry(group.len()).unwrap();
    //             if group.contains(c) {
    //                 let temp = exclude.get(&index).unwrap();
    //                 let opts = options.entry(*c).or_insert(HashSet::new());
    //                 *opts = &*opts - temp;
    //                 // println!("{:?}", opts);
    //             } else {
    //                 let temp = digits.get(index).unwrap();
    //                 let opts = options.entry(*c).or_insert(HashSet::new());
    //                 *opts = &*opts - temp;
    //             }
    //         }
    //     } else if group.len() == 4 {
    //         for c in group {
    //             let temp = exclude.get(&4).unwrap();
    //             let opts = options.entry(*c).or_insert(HashSet::new());
    //             *opts = &*opts - temp;
    //         }
    //     }
    // }
    // for i in options {
    //     println!("{:?}", i);
    // }
    // println!("{:?}", options);

    // println!("Groups: {:?}", groups);
    // println!("Output: {:?}", output);
}

pub fn part1(input: String) {
    let mut count = 0;
    for line in input.split('\n') {
        let pair: Vec<&str> = line.trim().split(" | ").collect::<Vec<&str>>();
        let output: Vec<&str> = (*pair[1].split(" ").collect::<Vec<&str>>()).to_vec();
        for val in output {
            println!("Val: {}", val);
            if (val.len() >= 2 && val.len() <= 4) || val.len() == 7 {
                count += 1;
            }
        }
    }

    println!("{} 1's, 4's, 7's, and 8's", count);
}