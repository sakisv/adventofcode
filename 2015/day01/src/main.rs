use std::fs;

fn get_input(filename: &str) -> String {
    fs::read_to_string(filename).unwrap_or(format!("Could not load {filename}").to_string())
}

fn solve_part1(input: &String) -> i32 {
    let mut floor = 0;
    for c in input.chars() {
        if c == '(' {
            floor += 1
        } else if c == ')' {
            floor -= 1
        }
    }
    floor
}

fn solve_part2(input: &String) -> i32 {
    let mut floor = 0;
    for (i, c) in input.chars().enumerate() {
        if c == '(' {
            floor += 1
        } else if c == ')' {
            floor -= 1
        }

        if floor == -1 {
            return (i + 1).try_into().unwrap();
        }
    }
    floor
}

fn main() {
    let filename = "input.txt";
    let input = get_input(filename);
    let part1 = solve_part1(&input);
    println!("Part 1: {part1}");

    let part2 = solve_part2(&input);
    println!("Part 2: {part2}");
}
