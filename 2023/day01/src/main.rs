use std::fs::read_to_string;

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut sum = 0;
    for line in input {
        let mut line_digits = String::from("");
        for c in line.chars() {
            if c.is_ascii_digit() {
                line_digits.push(c);
            }
        }
        let line_numbers: i32;
        let first_last: String;
        let first_digit = line_digits.remove(0);

        if line_digits.len() > 0 {
            let last_digit = line_digits.remove(line_digits.len()-1);
            first_last = format!("{}{}", first_digit, last_digit);
            line_numbers = first_last.parse().unwrap();
        } else {
            first_last = format!("{}{}", first_digit, first_digit);
            line_numbers = first_last.parse().unwrap();
        }
        sum = sum + line_numbers;
        println!("Line: {line} - Digits: {line_digits} - first_last: {first_last} - As int: {line_numbers} - Sum: {sum}");
    }

    sum
}

fn solve_part2(input: &Vec<String>) -> i32 {
    let sum = 0;
    sum
}

fn get_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
    .unwrap()  // panic on possible file-reading errors
    .lines()  // split the string into an iterator of string slices
    .map(String::from)  // make each slice into a string
    .collect()  // gather them together into a vector
}

fn main() {
    let filename = "input.txt";
    let input = get_input(filename);
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
            String::from("1abc2"),
            String::from("pqr3stu8vwx"),
            String::from("a1b2c3d4e5f"),
            String::from("treb7uchet"),
        ];
        let result = super::solve_part1(&input);
        assert_eq!(result, 142);
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            String::from("two1nine"),
            String::from("eightwothree"),
            String::from("abcone2threexyz"),
            String::from("xtwone3four"),
            String::from("4nineeightseven2"),
            String::from("zoneight234"),
            String::from("7pqrstsixteen"),
        ];
        let result = super::solve_part2(&input);
        assert_eq!(result, 281);
    }
}
