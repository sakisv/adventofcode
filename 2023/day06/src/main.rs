use std::{fs::read_to_string, collections::HashSet, i32};
use regex::Regex;

fn get_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
    .unwrap()  // panic on possible file-reading errors
    .lines()  // split the string into an iterator of string slices
    .map(String::from)  // make each slice into a string
    .collect()  // gather them together into a vector
}

fn get_numbers_vector(line: &str) -> Vec<i32> {
    let re = Regex::new(r"\d+").unwrap();
    let ms = re.find_iter(&line).map(|m| m.as_str()).map(|i| i.parse().unwrap()).collect();
    ms
}

fn get_numbers_as_string_vector(line: &str) -> Vec<String> {
    let re = Regex::new(r"\d+").unwrap();
    let ms = re.find_iter(&line).map(|m| m.as_str()).map(String::from).collect();
    ms
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let race_durations = get_numbers_vector(&input[0]);
    let distances = get_numbers_vector(&input[1]);

    let mut winning_ways: Vec<i32> = vec![0; race_durations.len()];

    for i in 0..race_durations.len() {
        let race_duration = race_durations[i];
        let distance_to_beat = distances[i];

        for t in 1..race_duration+1 {
            let travel_distance = t * (race_duration - t);
            if travel_distance > distance_to_beat{
                winning_ways[i] += 1;
            }
        }
    }

    winning_ways.into_iter().product()
}

fn solve_part2(input: &Vec<String>) -> i32 {
    let mut winning_ways = 0;
    let race_duration: u128 = get_numbers_as_string_vector(&input[0]).join("").parse().unwrap();
    let distance_to_beat: u128 = get_numbers_as_string_vector(&input[1]).join("").parse().unwrap();

    for t in 1..race_duration+1 {
        let travel_distance = t * (race_duration - t);
        if travel_distance > distance_to_beat {
            winning_ways += 1;
        }
    }

    winning_ways
}


fn main() {
    let input = get_input("input.txt");

    let part1 = solve_part1(&input);
    println!("Part 1: {part1}");

    let part2 = solve_part2(&input);
    println!("Part 2: {part2}");
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_part1() {
        let input = vec![
            String::from("Time:      7  15   30"),
            String::from("Distance:  9  40  200"),
        ];

        let result = super::solve_part1(&input);
        assert_eq!(result, 288);
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            String::from("Time:      7  15   30"),
            String::from("Distance:  9  40  200"),
        ];

        let result = super::solve_part2(&input);
        assert_eq!(result, 71503);
    }
}
