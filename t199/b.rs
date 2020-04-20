//! Build and run:
//!
//!     $ rustc b.rs && ./b
//!
//! Test using enum (multiple variant flavours) as a key in a HashMap.

use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq, Hash)]
enum State<'a> {
    A,
    B(i32),
    C(&'a mut St),
}

#[derive(Debug, PartialEq, Eq, Hash)]
struct St {
    x: i32,
}

fn main() {
    let mut m: HashMap<State, String> = HashMap::new();

    // Populate the HashMap
    assert_eq!(None, m.insert(State::A, String::from("a")));
    assert_eq!(None, m.insert(State::B(0), String::from("b")));
    let mut st = St { x: 1 };
    assert_eq!(None, m.insert(State::C(&mut st), String::from("c")));

    // Print the HashMap
    eprintln!("{:?}", m);

    // Check various enums as keys
    eprintln!("\nm.get(&State::A) = {:?}", m.get(&State::A));

    eprintln!("\nm.get(&State::B(0)) = {:?}", m.get(&State::B(0)));
    eprintln!("m.get(&State::B(1)) = {:?}", m.get(&State::B(1)));

    eprintln!(
        "\nm.get(&State::C(&mut St {{ x: 0 }})) = {:?}",
        m.get(&State::C(&mut St { x: 0 })),
    );
    eprintln!(
        "m.get(&State::C(&mut St {{ x: 1 }})) = {:?}",
        m.get(&State::C(&mut St { x: 1 })),
    );

    // This time, using State::C's ability to mutate its embedded type, mutate
    // and check more :D
    let mut st1 = St { x: 0 };
    let mut st2 = St { x: 42 };
    eprintln!("\nst1 = {:?}, st2 = {:?}", st1, st2);
    let mut c1 = State::C(&mut st1);
    let mut c2 = State::C(&mut st2);
    eprintln!("c1 = {:?}, c2 = {:?}", c1, c2);
    eprintln!("m.get(&c1) = {:?}", m.get(&c1));
    eprintln!("m.get(&c2) = {:?}", m.get(&c2));

    //st1.x = 1; /// error: &mut ref to st1 still alive
    //st2.x = 1; /// error: &mut ref to st2 still alive
    // Destructure each of the enum instances to mutate their contents using
    // the embedded references to the original St values, since st1 and st2
    // cannot be used while the mutable references are through c1 and c2.
    match c1 {
        State::C(ref mut st) => {
            st.x = 1; // Notice that it's a mutation, not a new allocation
        }
        _ => (),
    }
    match c2 {
        State::C(ref mut st) => {
            st.x = 1; // Notice that it's a mutation, not a new allocation
        }
        _ => (),
    }
    //eprintln!("\nst1 = {:?}, st2 = {:?}", st1, st2); /// error: &mut ref to st2 still alive
    // Having mutated st1 and st2 through their &mut references in c1 and c2
    // respectively, check the HashMap again, and be amazed.
    eprintln!("c1 = {:?}, c2 = {:?}", c1, c2);
    eprintln!("m.get(&c1) = {:?}", m.get(&c1));
    eprintln!("m.get(&c2) = {:?}", m.get(&c2));
}
