use std::{collections::HashMap, fs};

fn get_input(filename: &str) -> String {
    fs::read_to_string(filename).unwrap_or(format!("Could not load {filename}").to_string())
}

struct GiftGiver {
    x: i32,
    y: i32,
}

impl GiftGiver {
    fn new() -> Self {
        GiftGiver { x: 0, y: 0 }
    }

    fn make_move(&mut self, c: char) {
        match c {
            '^' => self.y += 1,
            'v' => self.y -= 1,
            '<' => self.x -= 1,
            '>' => self.x += 1,
            _ => (),
        }
    }

    fn get_current_coords(&self) -> String {
        format!("{}-{}", self.x, self.y)
    }
}

fn solve_part1(input: &String) -> u32 {
    let mut px = 0;
    let mut py = 0;
    let mut houses = HashMap::new();

    houses.insert(format!("{px}-{py}"), 1);
    for c in input.chars() {
        if c == '^' {
            py += 1;
        } else if c == 'v' {
            py -= 1;
        } else if c == '<' {
            px -= 1;
        } else if c == '>' {
            px += 1;
        } else {
            continue;
        }

        houses
            .entry(format!("{px}-{py}"))
            .and_modify(|visits| *visits += 1)
            .or_insert(1);
    }
    return houses.len().try_into().unwrap();
}

fn solve_part2(input: &String) -> u32 {
    let mut santa = GiftGiver::new();
    let mut robot = GiftGiver::new();
    let mut houses = HashMap::new();

    houses.insert(format!("0-0"), 2);
    let mut cur: &mut GiftGiver;
    for (i, c) in input.chars().enumerate() {
        if (i + 1) % 2 == 0 {
            cur = &mut santa;
        } else {
            cur = &mut robot;
        }
        cur.make_move(c);
        houses
            .entry(cur.get_current_coords())
            .and_modify(|visits| *visits += 1)
            .or_insert(1);
    }
    return houses.len().try_into().unwrap();
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
        let input = "^v^v^v^v^v".to_string();
        let result = super::solve_part1(&input);
        assert_eq!(result, 2);
    }

    #[test]
    fn solve_part2() {
        let input = "^v^v^v^v^v".to_string();
        let result = super::solve_part2(&input);
        assert_eq!(result, 11);
    }
}
