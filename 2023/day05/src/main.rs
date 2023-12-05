use std::fs::read_to_string;
use regex::Regex;

#[derive(Debug)]
struct Instructions {
    destination_start: u32,
    destination_end: u32,
    source_start: u32,
    source_end: u32,
    range: u32,
}

impl Instructions {
    fn new(d: u32, s: u32, r: u32) -> Self {
        Self { destination_start: d, destination_end: d + (r-1), source_start: s, source_end: s + (r-1), range: r }
    }

    fn get_seed_mapping(&self, seed: u32) -> u32 {
        let mut new_seed = seed;
        if seed >= self.source_start && seed <= self.source_end {
            let diff = seed - self.source_start;
            new_seed = self.destination_start + diff;
        }
        new_seed
    }
}

fn get_sections(filename: &str) -> Vec<String> {
    read_to_string(filename).unwrap().split("\n\n").map(String::from).collect()
}

fn get_map(map_string: String) -> Vec<Instructions> {
    let map_lines: Vec<String> = map_string.lines().map(String::from).collect();
    let mut map: Vec<Instructions> = vec![];

    let get_numbers = Regex::new(r"\d+").unwrap();
    // first line is the name, we skip it
    for i in 1..map_lines.len() {
        let instruction: Vec<u32> = get_numbers.find_iter(&map_lines[i]).map(|i| i.as_str().parse().unwrap()).collect();

        map.push(Instructions::new(instruction[0], instruction[1], instruction[2] ));
    }

    map
}

fn solve_part1(sections: &Vec<String>) -> u32 {
    let get_numbers = Regex::new(r"\d+").unwrap();
    let seeds: Vec<u32> = get_numbers.find_iter(&sections[0]).map(|i| i.as_str().parse().unwrap()).collect();
    let mut maps: Vec<Vec<Instructions>> = vec![];

    for i in 1..sections.len() {
        maps.push(get_map(sections[i].clone()))
    }

    let mut last_seed_values: Vec<u32> = vec![];

    for i in 0..seeds.len() {
        let mut seed = seeds[i];
        //println!("Seed: {seed}");
        for map in &maps {
            for instruction in map {
                let new_seed = instruction.get_seed_mapping(seed);
                if new_seed != seed {
                    seed = new_seed;
                    break;
                }
            }
            //println!("End of map seed: {seed}");
        }
        last_seed_values.push(seed);
    }

    //println!("{:#?}", last_seed_values);
    last_seed_values.into_iter().reduce(u32::min).unwrap()
}

fn main() {
    let sections = get_sections("input.txt");

    let part1 = solve_part1(&sections);
    println!("Part 1: {part1}");
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_part1() {
        let input = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";

        let sections = input.split("\n\n").map(String::from).collect();

        let result = super::solve_part1(&sections);
        assert_eq!(result, 35);
    }
}
