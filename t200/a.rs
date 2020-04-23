use std::cell::RefCell;
use std::fmt;
use std::rc::Rc;

trait State: fmt::Debug {
    fn process(&mut self) -> bool;
}

#[derive(Debug)]
struct A {}

#[derive(Debug)]
struct B(i32);

#[derive(Debug)]
struct C {
    x: i32,
}

impl State for A {
    fn process(&mut self) -> bool {
        false
    }
}

impl State for B {
    fn process(&mut self) -> bool {
        self.0 += 1;
        true
    }
}

impl State for C {
    fn process(&mut self) -> bool {
        self.x += 1;
        true
    }
}

struct W<'a> {
    state: Option<RefCell<Rc<&'a dyn State>>>,
    states: [RefCell<Rc<&'a dyn State>>; 3],
    //states: [Box<dyn State>; 3],
}

const ZERO: usize = 0;

impl<'a> W<'a> {
    fn init(&mut self) {
        //self.state = Some(RefCell::clone(Rc::clone(self.states[ZERO])));
        //self.state = Some(RefCell::clone(Rc::clone(&self.states[ZERO].borrow_mut())));
        self.state = Some(RefCell::clone(&Rc::clone(&self.states[ZERO].borrow_mut())));
        //self.state = Some(&mut self.states[ZERO]);
    }
}

fn main() {
    let mut state = None;
    //let mut states: [Box<dyn State>; 3] = [Box::new(A {}), Box::new(B(1)), Box::new(C { x: 42 })];
    let mut states: [RefCell<Rc<&dyn State>>; 3] = [
        RefCell::new(Rc::new(&A {})),
        RefCell::new(Rc::new(&B(1))),
        RefCell::new(Rc::new(&C { x: 42 })),
    ];

    let w = W { state, states };

    for s in &w.states {
        eprintln!("{:?}", s);
    }
}
