use std::fs::read_to_string;
use regex::Regex;

fn get_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
    .unwrap()  // panic on possible file-reading errors
    .lines()  // split the string into an iterator of string slices
    .map(String::from)  // make each slice into a string
    .collect()  // gather them together into a vector
}

fn get_game_number_and_color_picks(line: &str) -> (i32, String) {
    let line_split: Vec<&str> = line.split(":").collect();
    let _game_split: Vec<&str> = line_split[0].split(" ").collect();
    let game_number = _game_split[1].parse().unwrap();
    let color_picks = String::from(line_split[1]);

    (game_number, color_picks)
}

fn get_color_counts(color_picks: String) -> Vec<(i32, String)> {
    let re = Regex::new(r"(?<count>\d+)\s(?<color>red|green|blue)").unwrap();
    let color_counts: Vec<(i32, String)> = re.captures_iter(&color_picks).map(|cap|{
        let color = String::from(cap.name("color").unwrap().as_str());
        let count: i32 = cap.name("count").unwrap().as_str().parse().unwrap();
        (count, color)
    }).collect();

    color_counts
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut sum = 0;
    let max_red = 12;
    let max_green = 13;
    let max_blue = 14;

    for line in input {
        let (game_number, color_picks) = get_game_number_and_color_picks(line);
        let color_counts = get_color_counts(color_picks);

        let mut possible_game = true;
        for (count, color) in color_counts {
            if color == "red" && count > max_red {
                possible_game = false;
                break;
            }
            if color == "green" && count > max_green {
                possible_game = false;
                break;
            }
            if color == "blue" && count > max_blue {
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


fn solve_part2(input: &Vec<String>) -> i32 {
    let mut cubes_power = 0;
    for line in input {
        let (game_number, color_picks) = get_game_number_and_color_picks(line);
        let color_counts = get_color_counts(color_picks);

        let mut max_red = 1;
        let mut max_green = 1;
        let mut max_blue = 1;
        for (count, color) in color_counts {
            if color == "red" && count > max_red {
                max_red = count;
            }
            if color == "green" && count > max_green {
                max_green = count;
            }
            if color == "blue" && count > max_blue {
                max_blue = count;
            }
        }

        cubes_power = cubes_power + (max_red * max_green * max_blue);
    }
    cubes_power
}


fn main() {
    let input = get_input("input.txt");

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
            String::from("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
            String::from("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"),
            String::from("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"),
            String::from("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"),
            String::from("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"),
        ];

        let result = super::solve_part1(&input);
        assert_eq!(result, 8);
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            String::from("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
            String::from("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"),
            String::from("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"),
            String::from("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"),
            String::from("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"),
        ];

        let result = super::solve_part2(&input);
        assert_eq!(result, 2286);
    }
}
