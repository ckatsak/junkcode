//! Build and run:
//!
//!     $ rustc a.rs && ./a
//!
//! Test using enum (multiple variant flavours) as a key in a HashMap.

use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq, Hash)]
enum State {
    A,
    B(i32),
    C(St),
}

#[derive(Debug, PartialEq, Eq, Hash)]
struct St {
    x: i32,
}

fn main() {
    let mut m: HashMap<State, String> = HashMap::new();

    assert_eq!(None, m.insert(State::A, String::from("a")));
    assert_eq!(None, m.insert(State::B(0), String::from("b")));
    assert_eq!(None, m.insert(State::C(St { x: 1 }), String::from("c")));

    eprintln!("{:?}", m);

    eprintln!("\nm.get(&State::A) = {:?}", m.get(&State::A));

    //eprintln!("m.get(&State::B) = {:?}", m.get(&State::B)); /// error
    eprintln!("\nm.get(&State::B(0)) = {:?}", m.get(&State::B(0)));
    eprintln!("m.get(&State::B(1)) = {:?}", m.get(&State::B(1)));

    //eprintln!("m.get(&State::C) = {:?}", m.get(&State::C)); /// error
    eprintln!(
        "\nm.get(&State::C(St{{ x: 0 }})) = {:?}",
        m.get(&State::C(St { x: 0 })),
    );
    eprintln!(
        "m.get(&State::C(St{{ x: 1 }})) = {:?}",
        m.get(&State::C(St { x: 1 })),
    );

    let mut st1 = St { x: 0 };
    let mut st2 = St { x: 42 };
    eprintln!("\nst1 = {:?}, st2 = {:?}", st1, st2);
    let mut c1 = State::C(st1);
    let mut c2 = State::C(st2);
    eprintln!("c1 = {:?}, c2 = {:?}", c1, c2);
    eprintln!("m.get(&c1) = {:?}", m.get(&c1));
    eprintln!("m.get(&c2) = {:?}", m.get(&c2));
    //st1.x = 1;
    //st2.x = 1;
    //eprintln!("st1 = {:?}, st2 = {:?}", st1, st2);
    //eprintln!("m.get(&State::C(st1)) = {:?}", m.get(&State::C(st1)));
    //eprintln!("m.get(&State::C(st2)) = {:?}", m.get(&State::C(st2)));
}
