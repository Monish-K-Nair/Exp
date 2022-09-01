use std::io;

fn main() {
    let mut str_input = String::new();
    io::stdin()
        .read_line(&mut str_input)
        .expect("failed to read");

    let ar: Vec<u32> = str_input.split(" ")
        .map(|x| x.parse().expect("Not an integer!"))
        .collect();
    let sum: u32 = ar.iter().sum();
    
    println!("{}", sum);
}