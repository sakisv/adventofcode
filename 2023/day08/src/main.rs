use std::{fs::read_to_string, collections::HashMap, ops::Index};
use regex::Regex;

fn get_instructions_and_nodes(input: &String) -> (String, HashMap<String, Vec<String>>) {
    let sections: Vec<String>  = input.split("\n\n").map(String::from).collect();
    let instructions_string = sections[0].clone();

    let mut nodes: HashMap<String, Vec<String>> = HashMap::new();

    let regx = Regex::new(r"(?<key>[A-Z1-9]{3}) = \((?<dest_l>[A-Z1-9]{3}), (?<dest_r>[A-Z1-9]{3})\)").unwrap();
    for line in sections[1].lines() {
        let caps = regx.captures(line).unwrap();
        let k = caps.name("key").unwrap().as_str();
        let dest_l = caps.name("dest_l").unwrap().as_str();
        let dest_r = caps.name("dest_r").unwrap().as_str();

        nodes.insert(k.to_string(), vec![dest_l.to_string(), dest_r.to_string()]);
    }

    (instructions_string, nodes)
}

fn solve_part1(instructions: String, nodes: HashMap<String, Vec<String>>) -> u32 {
    let mut current_node = nodes.get_key_value("AAA").unwrap();
    let mut steps = 0;

    let instructions_count = instructions.len();
    let instructions_vec: Vec<char> = instructions.chars().collect();
    while current_node.0 != "ZZZ" {
        let instruction_index = steps % instructions_count;
        let instruction = instructions_vec.index(instruction_index).to_owned();
        steps += 1;

        if instruction == 'L' {
            current_node = nodes.get_key_value(current_node.1.index(0)).unwrap();
        } else {
            current_node = nodes.get_key_value(current_node.1.index(1)).unwrap();
        }
    }

    steps.try_into().unwrap()
}

fn are_all_nodes_final(nodes: Vec<String>) -> bool {
    nodes.iter().all(|i| i.ends_with("Z"))
}

fn solve_part2(instructions: String, nodes: HashMap<String, Vec<String>>) -> u32 {
    let mut current_nodes: Vec<(&String, &Vec<String>)> = nodes.iter().filter(|k| k.0.ends_with("A")).collect();
    let mut steps = 0;

    let instructions_count = instructions.len();
    let instructions_vec: Vec<char> = instructions.chars().collect();

    while ! are_all_nodes_final(current_nodes.iter().map(|k| k.0.to_owned()).collect()){
        let instruction_index = steps % instructions_count;
        let instruction = instructions_vec.index(instruction_index).to_owned();
        steps += 1;

        for n in 0..current_nodes.len() {
            let node = current_nodes.index(n);

            if instruction == 'L' {
                current_nodes[n] = nodes.get_key_value(node.1.index(0)).unwrap();
            } else {
                current_nodes[n] = nodes.get_key_value(node.1.index(1)).unwrap();
            }
        }
    }

    steps.try_into().unwrap()
}

fn main() {
    let input = read_to_string("input.txt").unwrap();
    let (instructions, nodes) = get_instructions_and_nodes(&input);
    let part1 = solve_part1(instructions, nodes);
    println!("Part 1: {part1}");

    let (instructions, nodes) = get_instructions_and_nodes(&input);
    let part2 = solve_part2(instructions, nodes);
    println!("Part 2: {part2}");
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_part1() {
        let input = "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)";

        let (instructions, nodes) = super::get_instructions_and_nodes(&input.to_string());
        println!("{:#?}", nodes);
        let steps = super::solve_part1(instructions, nodes);
        assert_eq!(steps, 6);
    }

    #[test]
    fn solve_part2() {
        let input = "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)";

        let (instructions, nodes) = super::get_instructions_and_nodes(&input.to_string());
        println!("{:#?}", nodes);
        let steps = super::solve_part2(instructions, nodes);
        assert_eq!(steps, 6);
    }
}
