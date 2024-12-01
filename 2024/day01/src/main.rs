use std::{collections::HashMap, fs};

fn get_input_lines(filename: &str) -> Vec<String> {
    fs::read_to_string(filename)
        .unwrap_or(format!("Could not load {filename}").to_string())
        .lines()
        .map(|l| String::from(l).trim().to_string())
        .collect()
}

fn get_input_string(filename: &str) -> String {
    fs::read_to_string(filename)
        .unwrap_or(format!("Could not load {filename}"))
        .trim()
        .to_string()
}

fn parse_list(input: &Vec<String>, column: usize) -> Vec<u32> {
    input
        .iter()
        .map(|l| l.split_ascii_whitespace().collect::<Vec<&str>>()[column])
        .map(|i| i.parse::<u32>().unwrap())
        .collect()
}

fn solve_part1(input: Vec<String>) -> u32 {
    let mut first_column = parse_list(&input, 0);
    let mut second_column = parse_list(&input, 1);

    first_column.sort();
    second_column.sort();

    let mut sum = 0;
    let mut diff;
    for i in 0..first_column.len() {
        diff = first_column[i].abs_diff(second_column[i]);
        sum += diff
    }
    sum
}

fn solve_part2(input: Vec<String>) -> u32 {
    let first_column = parse_list(&input, 0);
    let second_column = parse_list(&input, 1);

    let mut second_column_map: HashMap<u32, u32> = HashMap::new();
    let mut frequency: u32 = 0;
    for i in second_column {
        frequency = *second_column_map.get(&i).or(Some(&0)).unwrap();
        second_column_map.insert(i, frequency + 1);
    }

    let mut sum = 0;
    for i in first_column {
        sum += i * *second_column_map.get(&i).or(Some(&0)).unwrap();
    }

    sum
}

fn main() {
    let input = get_input_lines("input.txt");

    let part1 = solve_part1(input.clone());
    println!("Part 1: {}", part1);

    let part2 = solve_part2(input.clone());
    println!("Part 2: {}", part2);
}

#[cfg(test)]
mod tests {

    #[test]
    fn solve_part1() {
        let input = vec![
            "3   4".to_string(),
            "4   3".to_string(),
            "2   5".to_string(),
            "1   3".to_string(),
            "3   9".to_string(),
            "3   3".to_string(),
        ];
        assert_eq!(11, super::solve_part1(input));
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            "3   4".to_string(),
            "4   3".to_string(),
            "2   5".to_string(),
            "1   3".to_string(),
            "3   9".to_string(),
            "3   3".to_string(),
        ];
        assert_eq!(31, super::solve_part2(input));
    }
}
