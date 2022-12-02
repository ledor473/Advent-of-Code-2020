use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    let file_path = &args.get(1).expect("Missing argument for the path to the input file");

    let content = fs::read_to_string(file_path)
        .expect("Should have read the file");

    let inventory = content.split("\n\n").collect::<Vec<_>>();

    let mut calories = inventory.iter().map(|i| i.split("\n").map(|i| i.parse::<i32>().unwrap_or_default()).sum()).collect::<Vec<i32>>();
    calories.sort();
    calories.reverse();

    println!("Part 1: Most calories: {}", calories[0]);

    _ = calories.split_off(3);
    println!("Part 2: Sum of top 3: {}", calories.iter().sum::<i32>());
}