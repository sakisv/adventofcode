use std::{fs::read_to_string, collections::HashMap, cmp::Ordering};

#[derive(Debug, PartialEq, Eq, PartialOrd)]
enum HandType {
    FiveOfKind = 7, // 1 group (5)
    FourOfKind = 6, // 2 groups (4 + 1)
    FullHouse = 5, // 2 groups (3 + 2)
    ThreeOfKind = 4, // 3 groups (3 + 1 + 1)
    TwoPair = 3, // 3 groups (2 + 2 + 1)
    OnePair = 2, // 4 groups (2 + 1 + 1 + 1)
    HighCard = 1, // 5 groups (1 + 1 + 1 + 1 + 1)
}

#[derive(Debug)]
struct Hand {
    combo: String,
    hand_type: HandType,
    groups: HashMap<char, Vec<char>>,
    part2: bool,
}

impl Hand {
    fn find_groups(cards: &str) -> HashMap<char, Vec<char>> {
        let mut groups: HashMap<char, Vec<char>> = HashMap::new();
        for c in cards.chars() {
            groups.entry(c)
                .and_modify(|v| v.push(c))
                .or_insert(vec![c]);
        }
        groups
    }

    fn get_hand_type(groups: &HashMap<char, Vec<char>>) -> HandType {
        let group_count = groups.len();
        let group_v: Vec<&Vec<char>> = groups.values().collect();
        let mut ht = HandType::HighCard;

        if group_count == 5 {
            ht = HandType::HighCard;
        } else if group_count == 4 {
            ht = HandType::OnePair;
        } else if group_count == 3 {
            let max_v = group_v.into_iter().map(|v| v.len()).max().unwrap();
            ht = if max_v == 3 { HandType::ThreeOfKind } else { HandType::TwoPair };
        } else if group_count == 2 {
            let max_v = group_v.into_iter().map(|v| v.len()).max().unwrap();
            ht = if max_v == 4 { HandType::FourOfKind } else { HandType::FullHouse };
        } else {
           ht = HandType::FiveOfKind
        }

        ht
    }

    fn get_hand_type_part2(groups: &HashMap<char, Vec<char>>) -> HandType {
        // If there are no Joker cards, then reuse previous approach
        if ! groups.contains_key(&'J') {
            return Self::get_hand_type(groups);
        }

        let mut cloned_groups = groups.clone();
        let j_group = cloned_groups.remove(&'J').unwrap();
        let j_count = j_group.len();

        let group_count = cloned_groups.len();
        let group_v: Vec<&Vec<char>> = cloned_groups.values().collect();
        let mut ht = HandType::HighCard;

        if group_count == 4 {
            ht = HandType::OnePair; // actually 5 groups: 4 other cards + 1 Joker => HighCard becomes OnePair
        } else if group_count == 3 { // actually 4 groups - we add the jokers to the max_v in order to get the correct type
            let max_v = group_v.into_iter().map(|v| v.len()).max().unwrap();
            ht = if max_v + j_count == 3 { HandType::ThreeOfKind } else { HandType::TwoPair };
        } else if group_count == 2 { // actually 3 groups - but same as before
            let max_v = group_v.into_iter().map(|v| v.len()).max().unwrap();
            ht = if max_v + j_count == 4 { HandType::FourOfKind } else { HandType::FullHouse };
        } else { // 1 or 0 groups - either FourOfKind becomes FiveOfKind or we have 5 Jokers so we are at FiveOfKind anyway
           ht = HandType::FiveOfKind
        }

        ht
    }

    fn new(cards: &str, part2: bool) -> Self {
        let groups = Self::find_groups(cards);
        let ht = if part2 { Self::get_hand_type_part2(&groups) } else { Self::get_hand_type(&groups) } ;
        Hand { combo: String::from(cards), hand_type: ht, groups, part2 }
    }
}

impl PartialEq for Hand {
    fn eq(&self, other: &Self) -> bool {
        if self.hand_type != other.hand_type {
            return false;
        }
        (0..self.combo.len()).into_iter().all(|i| self.combo.chars().nth(i).unwrap() == other.combo.chars().nth(i).unwrap())
    }
}

impl PartialOrd for Hand {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Eq for Hand {}

impl Ord for Hand {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        let mut card_values = HashMap::from([
            ('A', 14),
            ('K', 13),
            ('Q', 12),
            ('J', 11),
            ('T', 10),
            ('9', 9),
            ('8', 8),
            ('7', 7),
            ('6', 6),
            ('5', 5),
            ('4', 4),
            ('3', 3),
            ('2', 2),
        ]);
        if self.part2 {
            card_values.insert('J', 1);
        }
        if self.hand_type > other.hand_type {
            return Ordering::Greater
        }
        if self.hand_type < other.hand_type {
            return Ordering::Less
        }
        if self.hand_type == other.hand_type {
            for i in 0..self.combo.len() {
                let c1 = self.combo.chars().nth(i).unwrap();
                let c2 = other.combo.chars().nth(i).unwrap();

                if card_values.get(&c1).unwrap() > card_values.get(&c2).unwrap() {
                    return Ordering::Greater;
                } else if card_values.get(&c1).unwrap() < card_values.get(&c2).unwrap() {
                    return Ordering::Less;
                }
            }
        }
        return Ordering::Equal;
    }
}

#[derive(Debug)]
struct Round {
    hand: Hand,
    bid: i32
}

impl Round {
    fn new(line: &str, part2: bool) -> Self {
        let splits: Vec<&str> = line.split_ascii_whitespace().collect();

        Round { hand: Hand::new(splits[0], part2), bid: splits[1].parse().unwrap() }
    }
}


fn get_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
    .unwrap()  // panic on possible file-reading errors
    .lines()  // split the string into an iterator of string slices
    .map(String::from)  // make each slice into a string
    .collect()  // gather them together into a vector
}

fn solve_part1(input: &Vec<String>) -> i32 {
    let mut winnings = 0;

    let mut rounds: Vec<Round> = vec![];

    for line in input {
        rounds.push(Round::new(line, false))
    }

    rounds.sort_by(|el1, el2| el1.hand.cmp(&el2.hand));

    for (i, v) in rounds.iter().enumerate() {
        let ii: i32 = i.try_into().unwrap();
        //println!("Hand {}: {} - Bid: {}", i+1, v.hand.combo, v.bid);
        winnings += (ii + 1) * v.bid;
    }

    winnings
}

fn solve_part2(input: &Vec<String>) -> i32 {
    let mut winnings = 0;

    let mut rounds: Vec<Round> = vec![];

    for line in input {
        rounds.push(Round::new(line, true))
    }

    rounds.sort_by(|el1, el2| el1.hand.cmp(&el2.hand));

    for (i, v) in rounds.iter().enumerate() {
        let ii: i32 = i.try_into().unwrap();
        //println!("Hand {}: {} ({:?}) - Bid: {}", i+1, v.hand.combo, v.hand.hand_type, v.bid);
        winnings += (ii + 1) * v.bid;
    }

    winnings
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
    use crate::Hand;

    #[test]
    fn solve_part1() {
        let input = vec![
            String::from("32T3K 765"),
            String::from("T55J5 684"),
            String::from("KK677 28"),
            String::from("KTJJT 220"),
            String::from("QQQJA 483"),
        ];

        let result = super::solve_part1(&input);
        assert_eq!(result, 6440);
    }

    #[test]
    fn solve_part2() {
        let input = vec![
            String::from("32T3K 765"),
            String::from("T55J5 684"),
            String::from("KK677 28"),
            String::from("KTJJT 220"),
            String::from("QQQJA 483"),
        ];

        let result = super::solve_part2(&input);
        assert_eq!(result, 5905);
    }

    #[test]
    fn find_hand_type_part2() {
        let h1 = Hand::new("32T3K", true);
        let h2 = Hand::new("T55J5", true);
        let h3 = Hand::new("KK677", true);
        let h4 = Hand::new("KTJJT", true);
        let h5 = Hand::new("QQQJA", true);

        assert_eq!(h1.hand_type, crate::HandType::OnePair);
        assert_eq!(h2.hand_type, crate::HandType::FourOfKind);
        assert_eq!(h3.hand_type, crate::HandType::TwoPair);
        assert_eq!(h4.hand_type, crate::HandType::FourOfKind);
        assert_eq!(h5.hand_type, crate::HandType::FourOfKind);
    }
    #[test]
    fn compare_hands_part1() {
        let h1 = Hand::new("KK677", false);
        let h2 = Hand::new("KTJJT", false);
        let h3 = Hand::new("32T3K", false);
        assert!(h1 > h2);
        assert!(h2 > h3);
        assert!(h1 > h3);

        let h4 = Hand::new("33332", false);
        let h5 = Hand::new("2AAAA", false);
        assert!(h4 > h5);

        let h6 = Hand::new("77888", false);
        let h7 = Hand::new("77788", false);
        assert!(h6 > h7);
    }

}
