
fn update_board(board: &mut Vec<Vec<i32>>, num: &i32) -> bool {
    let mut has_win = false;
    for row in 0..5 {
        for col in 0..5 {
            if board[row][col] == *num {
                board[row][col] = -1;
                if check_win(&board, row, col) {
                    has_win = true;
                }
            }
        }
    }
    return has_win;
}

fn check_win(board: &Vec<Vec<i32>>, row: usize, col: usize) -> bool {
    let mut row_sum = 0;
    let mut col_sum = 0;
    for i in 0..5 {
        row_sum += board[row][i];
        col_sum += board[i][col];
    }
    return (row_sum == -5 || col_sum == -5) as bool;
}

fn calculate_sum(board: &Vec<Vec<i32>>) -> i32 {
    let mut total = 0;
    for i in 0..5 {
        for j in 0..5 {
            if board[i][j] > 0 {
                total += board[i][j];
            }
        }
    }
    return total;
}

fn print_board(board: &Vec<Vec<i32>>) {
    for i in 0..5 {
        println!("{:?}", board[i]);
    }
    println!("");
}

pub fn part1(input: String) {
    let mut draw_numbers: Vec<i32> = Vec::new();
    let mut boards: Vec<Vec<Vec<i32>>> = Vec::new();
    let mut board_row = 0;
    let mut board = vec![vec![0; 5]; 5];
    for (i, line) in input.split('\n').enumerate() {
        if i == 0 {
            draw_numbers = line.trim().split(",").map(|s| s.parse().unwrap()).collect();
        } else if line.len() > 0 {
            let row = line.trim().split_whitespace().map(|s| s.parse::<i32>().unwrap()).collect();
            board[board_row%5] = row;
            board_row += 1;
            if (board_row - 1)/5 != board_row/5 {
                boards.push(board.clone());
                board = vec![vec![0; 5]; 5];
            }
        }
    }

    'drawing: for num in draw_numbers {
        for index in 0..boards.len() {
            let has_win = update_board(&mut boards[index], &num);
            if has_win {
                print_board(&boards[index]);
                let unmarked_sum = calculate_sum(&boards[index]);
                let final_score = unmarked_sum * num;
                println!("Board {} wins!", index+1);
                println!("Unmarked sum: {}", unmarked_sum);
                println!("Final score: {}", final_score);
                println!("");
                break 'drawing;
            }
        }
    }
}

pub fn part2(input: String) {
    let mut draw_numbers: Vec<i32> = Vec::new();
    let mut boards: Vec<Vec<Vec<i32>>> = Vec::new();
    let mut board_row = 0;
    let mut board = vec![vec![0; 5]; 5];
    for (i, line) in input.split('\n').enumerate() {
        if i == 0 {
            draw_numbers = line.trim().split(",").map(|s| s.parse().unwrap()).collect();
        } else if line.len() > 0 {
            let row = line.trim().split_whitespace().map(|s| s.parse::<i32>().unwrap()).collect();
            board[board_row%5] = row;
            board_row += 1;
            if (board_row - 1)/5 != board_row/5 {
                boards.push(board.clone());
                board = vec![vec![0; 5]; 5];
            }
        }
    }

    let mut winning_boards = vec![0; boards.len()];
    'drawing: for num in draw_numbers {
        for index in 0..boards.len() {
            let has_win = update_board(&mut boards[index], &num);
            if has_win {
                if winning_boards[index] == 0 && winning_boards.iter().sum::<i32>() == (boards.len() as i32) - 1 {
                    let unmarked_sum = calculate_sum(&boards[index]);
                    let final_score = unmarked_sum * num;
                    println!("Board {} wins last!", index+1);
                    println!("Last number: {}", num);
                    println!("Unmarked sum: {}", unmarked_sum);
                    println!("Final score: {}", final_score);
                    break 'drawing;
                }
                winning_boards[index] = 1;
            }
        }
    }
}