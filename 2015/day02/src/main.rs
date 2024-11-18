use std::fs;

#[derive(Debug)]
struct Present {
    length: u32,
    width: u32,
    height: u32,
    _smallest: u32,
}

impl Present {
    fn new(dimensions: String) -> Self {
        let parts: Vec<String> = dimensions.split("x").map(String::from).collect();
        Present {
            length: parts[0].parse::<u32>().unwrap(),
            width: parts[1].parse::<u32>().unwrap(),
            height: parts[2].parse::<u32>().unwrap(),
            _smallest: parts
                .iter()
                .map(|s| s.parse::<u32>().unwrap())
                .collect::<Vec<u32>>()
                .iter()
                .min()
                .unwrap()
                .to_owned(),
        }
    }

    fn get_wrapping_paper(&self) -> u32 {
        let a: Vec<u32> = Vec::from([
            2 * self.length * self.width,
            2 * self.width * self.height,
            2 * self.height * self.length,
        ]);

        let min = a.iter().min().unwrap();

        return a.iter().sum::<u32>() + min / 2;
    }

    fn get_total_ribbon(&self) -> u32 {
        let mut a: Vec<u32> = Vec::from([2 * self.length, 2 * self.width, 2 * self.height]);
        a.sort();
        let volume: u32 = self.length * self.width * self.height;

        return a[0] + a[1] + volume;
    }
}

fn get_input(filename: &str) -> Vec<String> {
    fs::read_to_string(filename)
        .unwrap_or(format!("Could not load {filename}").to_string())
        .lines()
        .map(String::from)
        .collect()
}

fn solve_part1(input: &Vec<String>) -> u32 {
    let mut total_paper = 0;
    for line in input {
        total_paper += Present::new(line.to_owned()).get_wrapping_paper();
    }

    total_paper
}

fn solve_part2(input: &Vec<String>) -> u32 {
    let mut total_ribbon = 0;
    for line in input {
        total_ribbon += Present::new(line.to_owned()).get_total_ribbon();
    }

    total_ribbon
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
        let input = vec!["2x3x4".to_string()];
        let result = super::solve_part1(&input);
        assert_eq!(result, 58);
    }

    #[test]
    fn solve_part2() {
        let input = vec!["2x3x4".to_string()];
        let result = super::solve_part2(&input);
        assert_eq!(result, 34);
    }
}
