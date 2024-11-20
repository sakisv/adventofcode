use std::fs;

fn get_input(filename: &str) -> String {
    fs::read_to_string(filename)
        .unwrap_or(format!("Could not load {filename}"))
        .trim()
        .to_string()
}

fn solve(input: &String, starts_with: String) -> u32 {
    let mut number: u32 = 0;
    while true {
        let current = format!("{}{}", input, number);
        let digest = md5::compute(current);
        let digest_string = format!("{:x}", digest);
        if digest_string.starts_with(&starts_with) {
            break;
        }
        if number % 1000000 == 0 {
            println!("Number: {}", number);
        }
        number += 1
    }
    return number;
}

fn main() {
    let filename = "input.txt";
    let input = get_input(filename);
    let part1 = solve(&input, String::from("00000"));
    println!("Part 1: {part1}");

    let part2 = solve(&input, String::from("000000"));
    println!("Part 2: {part2}");
}

#[cfg(test)]
mod tests {

    #[test]
    fn solve_part1() {
        let input = "abcdef".to_string();
        let result = super::solve(&input, String::from("00000"));
        assert_eq!(result, 609043);
    }
}
