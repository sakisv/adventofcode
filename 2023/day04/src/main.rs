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

fn split_line(line: &String) -> (HashSet<i32>, HashSet<i32>) {
    let split1: Vec<&str> = line.split(":").collect();
    let split2: Vec<&str> = split1[1].split("|").collect();

    let winning_numbers = get_numbers_vector(split2[0]);
    let our_numbers = get_numbers_vector(split2[1]);

    (HashSet::from_iter(winning_numbers), HashSet::from_iter(our_numbers))
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut sum = 0;

    for line in input {
        let (our_numbers, winning_numbers) = split_line(line);
        if our_numbers.is_disjoint(&winning_numbers) {
            continue;
        }
        let intersection: Vec<_> = our_numbers.intersection(&winning_numbers).collect();

        // 1 match => 2 ^ (1 - 1) => 2 ^ 0 = 1
        // 2 matches => 2 ^ (2 - 1) => 2 ^ 1 = 2
        // etc...
        sum += 2_i32.pow(u32::try_from(intersection.len()).unwrap() - 1_u32);
    }

    sum
}

fn solve_part2(input: &Vec<String>) -> i32 {
    let mut scratchcards_count = 0;
    let mut copies = vec![1; input.len()];

    for i in 0..input.len() {
        let (our_numbers, winning_numbers) = split_line(&input[i]);
        let matches: Vec<_> = our_numbers.intersection(&winning_numbers).collect();
        let match_count = matches.len();
        for _ in 0..copies[i] {
            scratchcards_count += 1;

            for j in i+1..i+1+match_count {
                copies[j] += 1;
            }
        }

    }

    scratchcards_count
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
            String::from("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"),
            String::from("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"),
            String::from("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"),
            String::from("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"),
            String::from("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"),
            String::from("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"),
        ];

        let result = super::solve_part1(&input);
        assert_eq!(result, 13);
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            String::from("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"),
            String::from("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"),
            String::from("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"),
            String::from("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"),
            String::from("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"),
            String::from("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"),
        ];

        let result = super::solve_part2(&input);
        assert_eq!(result, 30);
    }
}
