use std::fs::read_to_string;

fn first_and_last_digits_from_str(mut line_digits: String) -> String {
    let first_last: String;
    let first_digit = line_digits.remove(0);

    if line_digits.len() > 0 {
        let last_digit = line_digits.remove(line_digits.len()-1);
        first_last = format!("{}{}", first_digit, last_digit);
    } else {
        first_last = format!("{}{}", first_digit, first_digit);
    }

    return first_last;
}

fn extract_digits_as_str(line: String) -> String {
    let mut line_digits = String::from("");
    for c in line.chars() {
        if c.is_ascii_digit() {
            line_digits.push(c);
        }
    }
    line_digits
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut sum = 0;
    for line in input {
        let line_digits = extract_digits_as_str(line.to_string());
        let first_last = first_and_last_digits_from_str(line_digits.clone());
        let line_numbers: i32 = first_last.parse().unwrap();
        sum = sum + line_numbers;
        println!("Line: {line} - Digits: {line_digits} - first_last: {first_last} - As int: {line_numbers} - Sum: {sum}");
    }
    sum
}

fn extract_text_and_digits_as_str(line: String) -> String {
    let text_to_numbers = vec![
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9"),
        ("1", "1"),
        ("2", "2"),
        ("3", "3"),
        ("4", "4"),
        ("5", "5"),
        ("6", "6"),
        ("7", "7"),
        ("8", "8"),
        ("9", "9"),
    ];

    let mut first_digit = "";
    let mut last_digit = "";
    let mut first_position = line.len();
    let mut last_position = 0;

    for item in text_to_numbers {
        if line.contains(item.0) {
            let tmp_pos = line.find(item.0).unwrap();
            if tmp_pos < first_position {
                first_position = tmp_pos;
                first_digit = item.1;
            }
            if tmp_pos > last_position {
                last_position = tmp_pos;
                last_digit = item.1;
            }
        }
    }

    format!("{}{}", first_digit, last_digit)
}

fn solve_part2(input: &Vec<String>) -> i32 {
    let mut sum = 0;

    for line in input {
        let line_digits = extract_text_and_digits_as_str(line.to_string());
        let first_last = first_and_last_digits_from_str(line_digits.clone());
        let line_numbers: i32 = first_last.parse().unwrap();
        println!("Line: {line} - Digits: {line_digits} - first_last: {first_last} - As int: {line_numbers} - Sum: {sum}");
        sum = sum + line_numbers;
    }

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
            String::from("onethreetjpmpvpr6threethreeone"), // 11
        ];
        let result = super::solve_part2(&input);
        assert_eq!(result, 292); // 281 + 11
    }
}
