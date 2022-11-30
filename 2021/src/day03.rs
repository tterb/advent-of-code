
pub fn part1(input: String) {
    let mut count = 0;
    let mut bit_counts = vec![0; 0];
    for line in input.split('\n') {
        let bits = line.trim().chars().map(|c| c.to_digit(10).unwrap() as i32).collect();
        if bit_counts.len() == 0 {
            bit_counts = bits;
        } else {
            for (i, b) in bits.iter().enumerate() {
                bit_counts[i] += b;
            }
        }
        count += 1;
    }
    let mut gamma_rate = 0;
    let mut epsilon_rate = 0;
    let mut index = bit_counts.len() as u32;
    for b in bit_counts {
        let val = i32::pow(2, index - 1);
        if b > count/2 {
            gamma_rate += val;
        } else {
            epsilon_rate += val;
        }
        index -= 1;
    }
    println!("Gamma rate: {}", gamma_rate);
    println!("Epsilon rate: {}", epsilon_rate);
    println!("Power consumption: {}", gamma_rate * epsilon_rate);
}

fn get_common_bit(values: &Vec<Vec<i32>>, index: &usize, default: i32) -> i32 {
    let mut bit_count = 0.0;
    for val in values {
        bit_count += val[*index] as f32;
    }
    let half = (values.len() as f32)/2.0;
    if bit_count == half {
        return default;
    }
    return (bit_count > half) as i32;
}

pub fn part2(input: String) {
    let mut values: Vec<Vec<i32>> = Vec::new();
    for line in input.split('\n') {
        let bits = line.trim().chars().map(|c| c.to_digit(10).unwrap() as i32).collect();
        values.push(bits);
    }
    let mut index = 0;
    let mut generator_values = values.clone();
    while generator_values.len() > 1 {
        let common = get_common_bit(&generator_values, &index, 1);
        generator_values.retain(|val| {
            return val[index] == common;
        });
        index += 1;
    }
    index = 0;
    let mut scrubber_values = values.clone();
    while scrubber_values.len() > 1 {
        let common = get_common_bit(&scrubber_values, &index, 1);
        scrubber_values.retain(|val| {
            return val[index] != common;
        });
        index += 1;
    }
    let generator_rating = generator_values[0].iter().fold(0, |acc, &b| acc*2 + b as i32);
    let scrubber_rating = scrubber_values[0].iter().fold(0, |acc, &b| acc*2 + b as i32);
    println!("Generator rating: {}", generator_rating);
    println!("Scrubber rating: {}", scrubber_rating);
    println!("Life support rating: {}", generator_rating * scrubber_rating);
}
