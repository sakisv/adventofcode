use std::fs;

struct Report {
    levels: Vec<u32>,
}

impl Report {
    fn new(levels_str: &String) -> Self {
        Report {
            levels: levels_str
                .split_ascii_whitespace()
                .map(|l| l.parse::<u32>().unwrap())
                .collect(),
        }
    }

    fn check_rules_against(levels: &Vec<u32>) -> bool {
        // if it's not sorted
        if !(levels.is_sorted_by(|a, b| a <= b) || levels.is_sorted_by(|a, b| a >= b)) {
            return false;
        }

        // if the diff between the elements are < 1 or > 3
        let mut abs_diff;
        for i in 1..levels.len() {
            abs_diff = levels[i - 1].abs_diff(levels[i]);
            if abs_diff < 1 || abs_diff > 3 {
                return false;
            }
        }

        true
    }

    fn is_safe(&self) -> bool {
        Report::check_rules_against(&self.levels)
    }

    fn is_safe_part2(&self) -> bool {
        if Report::check_rules_against(&self.levels) {
            return true;
        }

        let mut poped_levels: Vec<u32>;
        for i in 0..self.levels.len() {
            poped_levels = self.levels.clone();
            poped_levels.remove(i);
            if Report::check_rules_against(&poped_levels) {
                return true;
            }
        }

        false
    }
}

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

fn conver_to_reports(input: &Vec<String>) -> Vec<Report> {
    input.iter().map(|l| Report::new(l)).collect()
}

fn solve_part1(input: Vec<String>) -> u32 {
    let reports = conver_to_reports(&input);

    reports
        .iter()
        .filter(|i| i.is_safe())
        .collect::<Vec<_>>()
        .len()
        .try_into()
        .unwrap()
}

fn solve_part2(input: Vec<String>) -> u32 {
    let reports = conver_to_reports(&input);

    reports
        .iter()
        .filter(|i| i.is_safe_part2())
        .collect::<Vec<_>>()
        .len()
        .try_into()
        .unwrap()
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
            "7 6 4 2 1".to_string(),
            "1 2 7 8 9".to_string(),
            "9 7 6 2 1".to_string(),
            "1 3 2 4 5".to_string(),
            "8 6 4 4 1".to_string(),
            "1 3 6 7 9".to_string(),
        ];
        assert_eq!(2, super::solve_part1(input));
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            "7 6 4 2 1".to_string(),
            "1 2 7 8 9".to_string(),
            "9 7 6 2 1".to_string(),
            "1 3 2 4 5".to_string(),
            "8 6 4 4 1".to_string(),
            "1 3 6 7 9".to_string(),
        ];
        assert_eq!(4, super::solve_part2(input));
    }
}
