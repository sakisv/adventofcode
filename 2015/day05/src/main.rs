use std::fs;

struct Line1 {
    l: String,
}

impl Line1 {
    fn new(line: String) -> Self {
        Line1 { l: line }
    }

    fn _has_three_vowels(&self) -> bool {
        let vowels = ['a', 'e', 'i', 'o', 'u'];
        let mut vowels_found: u32 = 0;
        for c in self.l.chars() {
            if vowels.contains(&c) {
                vowels_found += 1;
            }

            if vowels_found >= 3 {
                return true;
            }
        }
        false
    }

    fn _has_double_letters(&self) -> bool {
        for i in 1..self.l.len() {
            if self.l.chars().nth(i) == self.l.chars().nth(i - 1) {
                return true;
            }
        }
        false
    }

    fn _does_not_contain_bad_strings(&self) -> bool {
        let bad_strings = ["ab", "cd", "pq", "xy"];
        for bs in bad_strings {
            if self.l.contains(bs) {
                return false;
            }
        }
        true
    }

    fn is_good(self) -> bool {
        let res = [
            self._does_not_contain_bad_strings(),
            self._has_three_vowels(),
            self._has_double_letters(),
        ];

        println!("{:#?}", res);
        res.iter().all(|&i| i)
    }
}

struct Line2 {
    l: String,
}

impl Line2 {
    fn new(line: String) -> Self {
        Line2 { l: line }
    }

    fn _has_non_overlapping_double_pair(&self) -> bool {
        for i in 1..self.l.len() {
            let pair = format!(
                "{}{}",
                self.l.chars().nth(i - 1).unwrap(),
                self.l.chars().nth(i).unwrap()
            );
            let (_, end) = self.l.split_at(i + 1);
            if end.contains(&pair) {
                println!("String {}: {} exists in {}", self.l, pair, end);
                return true;
            }
        }
        false
    }

    fn _has_pair_with_one_letter_in_between(&self) -> bool {
        for i in 2..self.l.len() {
            if self.l.chars().nth(i) == self.l.chars().nth(i - 2) {
                return true;
            }
        }
        false
    }

    fn is_good(self) -> bool {
        let res = [
            self._has_non_overlapping_double_pair(),
            self._has_pair_with_one_letter_in_between(),
        ];

        println!("{:#?} -> {:#?}", self.l, res);
        res.iter().all(|&i| i)
    }
}

fn get_input(filename: &str) -> Vec<String> {
    fs::read_to_string(filename)
        .unwrap_or(format!("Could not load {filename}").to_string())
        .lines()
        .map(|l| String::from(l).trim().to_string())
        .collect()
}

fn solve_part1(input: &Vec<String>) -> u32 {
    let mut good_lines: u32 = 0;
    for line in input {
        let l = Line1::new(line.to_string());
        if l.is_good() {
            good_lines += 1;
        }
    }

    return good_lines;
}

fn solve_part2(input: &Vec<String>) -> u32 {
    let mut good_lines: u32 = 0;
    for line in input {
        let l = Line2::new(line.to_string());
        if l.is_good() {
            good_lines += 1;
        }
    }

    return good_lines;
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
            "ugknbfddgicrmopn".to_string(),
            "aaa".to_string(),
            "jchzalrnumimnmhp".to_string(),
            "haegwjzuvuyypxyu".to_string(),
            "dvszwmarrgswjxmb".to_string(),
        ];
        let result = super::solve_part1(&input);
        assert_eq!(result, 2);
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            "qjhvhtzxzqqjkmpb".to_string(),
            "xxyxx".to_string(),
            "uurcxstgmygtbstg".to_string(),
            "ieodomkazucvgmuy".to_string(),
            "aaa".to_string(),
        ];
        let result = super::solve_part2(&input);
        assert_eq!(result, 2);
    }
}
