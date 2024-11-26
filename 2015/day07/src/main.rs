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

fn create_graph_as_hashmap(input: &Vec<String>) -> HashMap<String, String> {
    let mut graph: HashMap<String, String> = HashMap::new();

    for l in input {
        if l.trim() == "" {
            continue;
        }
        let parts = l.split("->").collect::<Vec<&str>>();
        let command = String::from(parts[0].trim());
        let target = String::from(parts[1].trim());

        graph.insert(target, command);
    }

    graph
}

fn is_numeric(s: &str) -> bool {
    s.parse::<u16>().is_ok()
}

fn get_number(s: &str) -> u16 {
    s.parse::<u16>().unwrap()
}

fn solve(graph: &HashMap<String, String>, command: &String) -> u16 {
    // println!("Solving for {:#?}", command);
    let cmd_parts = command
        .split_ascii_whitespace()
        .map(|p| p.trim())
        .filter(|p| p.to_string() != "".to_string())
        .collect::<Vec<&str>>();

    match cmd_parts.len() {
        // 1 match:
        // a -> b
        // 1 -> c
        1 => {
            let part = cmd_parts[0];
            if is_numeric(part) {
                return get_number(part);
            } else {
                return solve(graph, graph.get(part).unwrap());
            }
        }
        // 2 matches:
        // NOT a -> c
        2 => {
            if cmd_parts[0] != "NOT" {
                panic!("2 matches and no NOT: {:#?}", cmd_parts);
            }
            let value = cmd_parts[1];
            if is_numeric(value) {
                return !get_number(value);
            } else {
                return !solve(graph, graph.get(value).unwrap());
            }
        }
        // 3 matches:
        // x AND y -> z
        // 1 OR k -> l
        // xk LSHIFT 13 -> j
        // k RSHIFT 2 -> ll
        3 => {
            let left = {
                let part = cmd_parts[0];
                if !is_numeric(part) {
                    solve(graph, &part.to_string())
                } else {
                    get_number(part)
                }
            };
            let center = cmd_parts[1];
            let right = {
                let part = cmd_parts[2];
                if !is_numeric(part) {
                    solve(graph, &part.to_string())
                } else {
                    get_number(part)
                }
            };

            match center {
                "AND" => left & right,
                "OR" => left | right,
                "LSHIFT" => left << right,
                "RSHIFT" => left >> right,
                _ => {
                    panic!("3 parts but center is wrong: {:#?}", center);
                }
            }
        }
        _ => {
            panic!("wtf")
        }
    }
}

fn solve_part1(input: &Vec<String>, target: String) -> u16 {
    let graph = create_graph_as_hashmap(input);
    assert_eq!(input.len(), graph.len());

    solve(&graph, graph.get(&target).unwrap())
}

fn main() {
    let filename = "input.txt";
    let input = get_input_lines(filename);
    let part1 = solve_part1(&input, String::from("a"));
    println!("Part 1: {part1}");
}

#[cfg(test)]
mod tests {

    #[test]
    fn solve_part1() {
        let input = vec![
            "123 -> x".to_string(),
            "456 -> y".to_string(),
            "x AND y -> d".to_string(),
            "x OR y -> e".to_string(),
            "x LSHIFT 2 -> f".to_string(),
            "y RSHIFT 2 -> g".to_string(),
            "NOT x -> h".to_string(),
            "NOT y -> i".to_string(),
        ];
        println!("Solving d: ");
        let d = super::solve_part1(&input, String::from("d"));
        assert_eq!(d, 72);

        println!("Solving e: ");
        let e = super::solve_part1(&input, String::from("e"));
        assert_eq!(e, 507);

        println!("Solving f: ");
        let f = super::solve_part1(&input, String::from("f"));
        assert_eq!(f, 492);

        println!("Solving g: ");
        let g = super::solve_part1(&input, String::from("g"));
        assert_eq!(g, 114);

        println!("Solving h: ");
        let h = super::solve_part1(&input, String::from("h"));
        assert_eq!(h, 65412);

        println!("Solving i: ");
        let i = super::solve_part1(&input, String::from("i"));
        assert_eq!(i, 65079);

        println!("Solving x: ");
        let x = super::solve_part1(&input, String::from("x"));
        assert_eq!(x, 123);

        println!("Solving y: ");
        let y = super::solve_part1(&input, String::from("y"));
        assert_eq!(y, 456);
    }

    #[test]
    fn solve_part1_again() {
        let input = vec![
            "1 -> on".to_string(),
            "2 -> to".to_string(),
            "on AND to -> x".to_string(),
            "x OR on -> a".to_string(),
        ];
        let a = super::solve_part1(&input, String::from("a"));
        assert_eq!(a, 1);
    }
}
