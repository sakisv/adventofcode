use std::fs;

use regex::Regex;

enum Action {
    TurnOn,
    TurnOff,
    Toggle,
}

struct Instruction {
    action: Action,
    from_row: u16,
    from_col: u16,
    end_row: u16,
    end_col: u16,
}

impl Instruction {
    fn new(instruction: &String) -> Self {
        let regex = Regex::new(
            r"(?<action>toggle|turn off|turn on) (?<from_row>\d+),(?<from_col>\d+) through (?<end_row>\d+),(?<end_col>\d+)",
        )
        .unwrap();
        let captures = regex.captures(&instruction).unwrap();
        let action: Action = match &captures["action"] {
            "turn on" => Action::TurnOn,
            "turn off" => Action::TurnOff,
            "toggle" => Action::Toggle,
            _ => panic!("wtf"),
        };

        Self {
            action: action,
            from_row: captures["from_row"].parse().unwrap(),
            from_col: captures["from_col"].parse().unwrap(),
            end_row: captures["end_row"].parse().unwrap(),
            end_col: captures["end_col"].parse().unwrap(),
        }
    }
}

#[derive(Debug)]
struct Grid {
    grid_size: u16,
    grid: Vec<Vec<u32>>,
    total_brightness: u32,
}

impl Grid {
    fn new(grid_size: u16) -> Self {
        let size = usize::try_from(grid_size).unwrap();
        Grid {
            grid_size: grid_size,
            grid: vec![vec![0; size]; size],
            total_brightness: 0,
        }
    }

    fn execute(&mut self, instruction: Instruction) {
        for row in instruction.from_row..=instruction.end_row {
            for col in instruction.from_col..=instruction.end_col {
                let u_row = usize::try_from(row).unwrap();
                let u_col = usize::try_from(col).unwrap();
                match instruction.action {
                    Action::TurnOn => self.grid[u_row][u_col] = 1,
                    Action::TurnOff => self.grid[u_row][u_col] = 0,
                    Action::Toggle => {
                        if self.grid[u_row][u_col] == 0 {
                            self.grid[u_row][u_col] = 1
                        } else {
                            self.grid[u_row][u_col] = 0
                        }
                    }
                }
            }
        }
    }

    fn execute_2(&mut self, instruction: Instruction) {
        for row in instruction.from_row..=instruction.end_row {
            for col in instruction.from_col..=instruction.end_col {
                let u_row = usize::try_from(row).unwrap();
                let u_col = usize::try_from(col).unwrap();
                match instruction.action {
                    Action::TurnOn => {
                        self.grid[u_row][u_col] += 1;
                        self.total_brightness += 1;
                    }
                    Action::TurnOff => {
                        if self.grid[u_row][u_col] > 0 {
                            self.grid[u_row][u_col] -= 1;
                            self.total_brightness -= 1;
                        }
                    }
                    Action::Toggle => {
                        self.grid[u_row][u_col] += 2;
                        self.total_brightness += 2;
                    }
                }
            }
        }
    }

    fn find_lit_lights(&self) -> u32 {
        let mut count = 0;
        for row in 0..self.grid_size {
            for col in 0..self.grid_size {
                let u_row = usize::try_from(row).unwrap();
                let u_col = usize::try_from(col).unwrap();
                if self.grid[u_row][u_col] == 1 {
                    count += 1;
                }
            }
        }
        count
    }
}

fn get_input(filename: &str) -> Vec<String> {
    fs::read_to_string(filename)
        .unwrap_or(format!("Could not load {filename}").to_string())
        .lines()
        .map(|l| String::from(l).trim().to_string())
        .collect()
}

fn solve_part1(input: &Vec<String>, grid_size: u16) -> u32 {
    let mut grid = Grid::new(grid_size);
    for line in input {
        let ins = Instruction::new(line);
        grid.execute(ins);
    }

    grid.find_lit_lights()
}

fn solve_part2(input: &Vec<String>, grid_size: u16) -> u32 {
    let mut grid = Grid::new(grid_size);
    for line in input {
        let ins = Instruction::new(line);
        grid.execute_2(ins);
    }

    grid.total_brightness
}

fn main() {
    let filename = "input.txt";
    let input = get_input(filename);
    let part1 = solve_part1(&input, 1000);
    println!("Part 1: {part1}");

    let part2 = solve_part2(&input, 1000);
    println!("Part 2: {part2}");
}

#[cfg(test)]
mod tests {

    #[test]
    fn solve_part1() {
        let input = vec!["toggle 998,998 through 999,999".to_string()];
        let result = super::solve_part1(&input, 1000);
        assert_eq!(result, 1000);
    }

    // #[test]
    // fn solve_part2() {
    //     let input = vec![
    //         "qjhvhtzxzqqjkmpb".to_string(),
    //         "xxyxx".to_string(),
    //         "uurcxstgmygtbstg".to_string(),
    //         "ieodomkazucvgmuy".to_string(),
    //         "aaa".to_string(),
    //     ];
    //     let result = super::solve_part2(&input);
    //     assert_eq!(result, 2);
    // }
}
