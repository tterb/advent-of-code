use std::collections::HashSet;
use regex::Regex;


fn check_passport(passport: &HashSet<String>) -> i32 {
    let required_fields: HashSet<String> = vec!["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"].iter().copied().map(|s| s.to_string()).collect();
    let diff = required_fields.difference(&passport).collect::<Vec<&String>>();
    if diff.len() == 0 {
        return 1;
    }
    return 0;
}

fn is_field_valid(name: &String, value: &String) -> bool {
    let eye_colors: HashSet<String> = vec!["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].iter().copied().map(|s| s.to_string()).collect();
    if name == "byr" || name == "iyr" || name == "eyr" {
        let year = &value.parse::<i32>().unwrap();
        if name == "byr" {
            return (1920..=2002).contains(year);
        } else if name == "iyr" {
            return (2010..=2020).contains(year);
        } else if name == "eyr" {
            return (2010..=2030).contains(year);
        }
    } else if name == "hgt" && value.len() <= 3 {
        let unit = &value[(value.len()-2)..];
        let num = &value[..(value.len()-2)].parse::<i32>().unwrap();
        if unit == "cm" {
            return (150..=193).contains(num);
        } else if unit == "in" {
            return (59..=76).contains(num);
        }
    } else if name == "hcl" {
        let re = Regex::new(r"(?i)^#[0-9A-F]{6}$").unwrap();
        return re.is_match(&value);
    } else if name == "pid" {
        let re = Regex::new(r"(?i)^[0-9]{9}$").unwrap();
        return re.is_match(&value);
    } else if name == "ecl" {
        return eye_colors.contains(value);
    } else if name == "cid" {
        return true;
    }
    return false;
}

pub fn part1(input: String) {
    let mut valid_count = 0;
    let mut passport: HashSet<String> = HashSet::new();
    for ln in input.split('\n') {
        let line = ln.split(" ").map(|s| s.to_string()).collect::<Vec<String>>();
        for field in line {
            if field.len() > 0 {
                let pair = field.split(":").map(|s| s.to_string()).collect::<Vec<String>>();
                passport.insert(pair[0].to_string());
            } else {
                valid_count += check_passport(&passport);
                passport.clear();
            }
        }
    }
    valid_count += check_passport(&passport);
    println!("{} valid passports", valid_count);
}

pub fn part2(input: String) {
    let mut valid_count = 0;
    let mut has_error = false;
    let mut passport: HashSet<String> = HashSet::new();
    for ln in input.split('\n') {
        let line = ln.split(" ").map(|s| s.to_string()).collect::<Vec<String>>();
        for field in line {
            if field.len() <= 0 {
                valid_count += check_passport(&passport);
                passport.clear();
                has_error = false;
            } else if field.len() > 0 && !has_error {
                let pair = field.split(":").map(|s| s.to_string()).collect::<Vec<String>>();
                let (name, value) = (pair[0].to_string(), pair[1].to_string());
                if is_field_valid(&name, &value) {
                    passport.insert(name);
                } else {
                    has_error = true;
                }
            }
        }
    }
    if !has_error {
        valid_count += check_passport(&passport);
    }
    println!("{} valid passports", valid_count);
}
