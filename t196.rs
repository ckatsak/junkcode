//! Compile:
//!
//!     $ rustc t196.rs
//!
//! and then run:
//!
//!     $ ./t196
//!
//! Try uncommenting any of the last two lines (and attempt to recompile it) to observe the point
//! of this snippet.

#[derive(Debug)]
struct MyStruct {
    a: Box<i32>,
    b: Box<i32>,
}

impl MyStruct {
    fn incr_a(&mut self, by: i32) {
        *self.a += by;
    }
}

fn main() {
    let mut x = Box::new(MyStruct {
        a: Box::new(42),
        b: Box::new(69),
    });
    eprintln!("Initialization:\n{:#?}\n", x);

    x.incr_a(8);
    eprintln!("x.incr_a(8):\n{:#?}\n", x);

    drop(x.b);
    *x.a += 10;

    let y = *x.a;
    eprintln!("y = *x.a + 10 (== {:#?})", y);

    // Uncommenting any of the lines below makes the compiler whine about
    // borrowing a value (at lines 43 or 44) after partial move (line 33).
    //
    //eprintln!("{:#?}", x);
    //x.incr_a(10);
}
