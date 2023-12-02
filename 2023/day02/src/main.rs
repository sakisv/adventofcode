use std::fs::read_to_string;
use regex::Regex;

fn get_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
    .unwrap()  // panic on possible file-reading errors
    .lines()  // split the string into an iterator of string slices
    .map(String::from)  // make each slice into a string
    .collect()  // gather them together into a vector
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut sum = 0;
    let max_red = 12;
    let max_green = 13;
    let max_blue = 14;

    for line in input {
        let line_split: Vec<&str> = line.split(":").collect();
        let _game_split: Vec<&str> = line_split[0].split(" ").collect();
        let game_number: i32 = _game_split[1].parse().unwrap();

        let re = Regex::new(r"(?<count>\d+)\s(?<color>red|green|blue)").unwrap();
        let color_counts: Vec<(&str, &str)> = re.captures_iter(line_split[1]).map(|cap|{
            let color = cap.name("color").unwrap().as_str();
            let count = cap.name("count").unwrap().as_str();
            (count, color)
        }).collect();

        let mut possible_game = true;
        for (count, color) in color_counts {
            let count_numb: i32 = count.parse().unwrap();
            if color == "red" && count_numb > max_red {
                possible_game = false;
                break;
            }
            if color == "green" && count_numb > max_green {
                possible_game = false;
                break;
            }
            if color == "blue" && count_numb > max_blue {
                possible_game = false;
                break;
            }
        }

        if possible_game {
            //println!("Game {} - Line {}", game_number, line_split[1]);
            sum += game_number;
        }
    }
    sum
}


fn main() {
    let input = get_input("input.txt");
    let part1 = solve_part1(&input);
    println!("Part 1: {part1}");
}

#[cfg(test)]
mod tests {
    #[test]
    fn solve_part1() {
        let input = vec![
            String::from("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
            String::from("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"),
            String::from("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"),
            String::from("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"),
            String::from("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"),
        ];

        let result = super::solve_part1(&input);
        assert_eq!(result, 8);
    }
}
