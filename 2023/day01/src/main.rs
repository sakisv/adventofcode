use std::fs::read_to_string;

fn solve_part1(input: Vec<String>) -> i32 {
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
    let sum = solve_part1(input);
    println!("{sum}");
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
        let result = super::solve_part1(input);
        assert_eq!(result, 142);
    }
}
