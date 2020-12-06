use std::collections::HashSet;


fn get_pivot(length: usize) -> usize {
    return (((length as f32) / 2.0).ceil() as i32) as usize;
}

fn get_seat_id(ident: &Vec<char>) -> usize {
    let mut row_range: Vec<usize> = (0..128).collect();
    let mut column_range: Vec<usize> = (0..8).collect();
    let mut index = 0;
    while index < ident.len() {
        if index < 7 {
            let pivot = get_pivot(row_range.len());
            if ident[index] == 'F' {
                row_range = row_range[..pivot].to_vec();
            } else if ident[index] == 'B' {
                row_range = row_range[pivot..].to_vec();
            }
        } else {
            let pivot = get_pivot(column_range.len());
            if ident[index] == 'L' {
                column_range = column_range[..pivot].to_vec();
            } else if ident[index] == 'R' {
                column_range = column_range[pivot..].to_vec();
            }
        }
        index += 1;
    }
    return (row_range[0] * 8) + column_range[0];
}

fn find_seat(open_seats: &HashSet<usize>) -> usize {
    // Sort open seats
    let mut sorted_seats = open_seats.iter().cloned().collect::<Vec<usize>>();
    sorted_seats.sort();
    let mut index = 1;
    while sorted_seats[index]-1 == sorted_seats[index-1] 
        || sorted_seats[index]+1 == sorted_seats[index+1] {
        index += 1;
    }
    return sorted_seats[index];
}

pub fn part1(input: String) {
    let mut max_id = 0;
    for line in input.split('\n') {
        let seat: Vec<char> = line.chars().collect();
        let seat_id = get_seat_id(&seat);
        if max_id < seat_id {
            max_id = seat_id;
        }
    }
    println!("Max Seat ID: {}", max_id);
}

pub fn part2(input: String) {
    let mut seats: HashSet<usize> = HashSet::new();
    for row in 0..128 {
        for col in 0..8 {
            seats.insert((row * 8) + col);
        }
    }
    for line in input.split('\n') {
        let seat: Vec<char> = line.chars().collect();
        if seat.len() > 0 {
            let seat_id = get_seat_id(&seat);
            seats.remove(&seat_id);
        }
    }
    let my_seat = find_seat(&seats);
    println!("My Seat ID: {}", my_seat);
}
