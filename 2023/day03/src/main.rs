use std::fs::read_to_string;
use regex::Regex;

fn get_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
    .unwrap()  // panic on possible file-reading errors
    .lines()  // split the string into an iterator of string slices
    .map(String::from)  // make each slice into a string
    .collect()  // gather them together into a vector
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut sum = 0;

    let re = Regex::new(r"\d+").unwrap();
    for i in 0..input.len() {
        let line = &input[i];
        if ! re.is_match(line) {
            continue
        }
        // if there are matches, check their surroundings for symbols
        let matches: Vec<_> = re.find_iter(line).collect();
        for m in matches {
            // ranges don't include the end part, i.e. they are [start, end)
            // matches already point to the next byte, so if we want to include that too, we point one more byte out
            // for rows, we need to end the range in current_row + 2 so that current_row + 1 is included
            let match_start = if m.start() > 0 { m.start() - 1} else { m.start()};
            let match_end = if m.end() == line.len() { m.end() } else { m.end() + 1 };

            let start_row = if i > 0 { i - 1 } else { i };
            let end_row = if i + 2 >= input.len() { input.len() } else {i + 2};
            let mut should_keep = false;
            for row in start_row..end_row {
                for col in match_start..match_end {
                    let cur_char = &input[row].chars().nth(col).expect(format!("failed for line {i} - row {row} col {col}\n{}", input[row]).as_str());
                    //println!("Line: {} - Match: {} - Row: {} - Col: {} - Char: {}", line, m.as_str(), row, col, cur_char);
                    if ! cur_char.is_numeric() && ! cur_char.eq_ignore_ascii_case(&'.') {
                        should_keep = true;
                    }
                }
            }

            if should_keep {
                let m_as_number: i32 = m.as_str().parse().unwrap();
                sum = sum + m_as_number;
                //println!("Keeping {} - Current sum: {}", m_as_number, sum);
            }
        }
    }

    sum
}

fn main() {
    let input = get_input("input.txt");

    let part1 = solve_part1(&input);
    println!("Part 1: {}", part1);
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_part1() {
        let input = vec![
            String::from("467..114.."),
            String::from("...*......"),
            String::from("..35..633."),
            String::from("......#..."),
            String::from("617*......"),
            String::from(".....+.58."),
            String::from("..592....."),
            String::from("......755."),
            String::from("...$.*...."),
            String::from(".664.598.."),
        ];

        let result = super::solve_part1(&input);
        assert_eq!(result, 4361);
    }
}
