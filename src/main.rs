use std::io::*;

fn main() {
    init();
}

fn init() {
    println!("Welcome to blnote!");
    'infinite: loop {
        println!("Please enter a command:");
        let mut input = String::new();
        stdin().read_to_string(&mut input).expect("Unable to read input");
        match input.trim_end() {
            x if x.contains("new") => {
                let name = {
                    let words = input.split(" ").map(|x| x.to_string()).collect::<Vec<String>>();
                    if words.len() > 1 { words[1].clone() }
                    else { "Untitled".to_string() }
                };
                new_note(name);
            },
            "show" => show_commands(),
            "exit" => {
                println!("Shutting down...\nGoodbye!");
                break 'infinite;
            },
            "help" => help_menu(),
            _ => println!("Unknown command"),
        }
    }
}

fn new_note(name: String) {
    println!("Please enter the text for your note:\n");
    let mut input = String::new();
    stdin().read_to_string(&mut input).expect("Unable to read input");
    save_note(name, input);
}

fn save_note(name: String, input: String) {
    println!("Saving note: {}, with content: {}", name.trim_end(), input);
    println!("(Nothin happens yet.)");
}

fn help_menu() {
    println!("Welcome to the help menu!");
    show_commands();
}

fn show_commands() {
    let commands = vec!["new <note name>", "help", "exit", "show"];
    println!("The available commands are:");
    for command in commands {
        println!("{}", command);
    }
}
