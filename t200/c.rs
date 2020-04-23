use std::cell::RefCell;
use std::fmt;
use std::rc::Rc;

trait State: fmt::Debug {
    fn process(&mut self, w: &mut W) -> bool;
}

#[derive(Debug)]
struct A;

#[derive(Debug)]
struct B(i32);

#[derive(Debug)]
struct C {
    x: i32,
}

impl State for A {
    fn process(&mut self, w: &mut W) -> bool {
        w.data += 1;
        w.curr = Some(Rc::clone(&w.states[ONE]));

        if w.data > 50 {
            false
        } else {
            true
        }
    }
}

impl State for B {
    fn process(&mut self, w: &mut W) -> bool {
        self.0 += 1;

        w.data += 1;
        w.curr = Some(Rc::clone(&w.states[TWO]));

        if w.data > 50 {
            false
        } else {
            true
        }
    }
}

impl State for C {
    fn process(&mut self, w: &mut W) -> bool {
        self.x += 1;

        w.data += 1;
        w.curr = Some(Rc::clone(&w.states[ZERO]));

        if w.data > 50 {
            false
        } else {
            true
        }
    }
}

#[derive(Debug)]
struct W {
    data: u8,
    curr: Option<Rc<RefCell<dyn State>>>,
    states: [Rc<RefCell<dyn State>>; 3],
}

const ZERO: usize = 0;
const ONE: usize = 1;
const TWO: usize = 2;

impl W {
    fn init(&mut self) {
        self.curr = Some(Rc::clone(&self.states[ZERO]));
    }

    fn run(&mut self) -> bool {
        eprintln!("{:?}", self);

        while self.curr.as_ref().unwrap().borrow_mut().process(self) {}

        //let mut ret = true;
        //while ret {
        //    let state = self.curr.as_ref().unwrap();
        //    ret = state.borrow_mut().process(self);
        //}

        true
    }
}

fn main() {
    let curr = None;
    let states: [Rc<RefCell<dyn State>>; 3] = [
        Rc::new(RefCell::new(A {})),
        Rc::new(RefCell::new(B(1))),
        Rc::new(RefCell::new(C { x: 42 })),
    ];

    let mut w = W {
        data: 0,
        curr,
        states,
    };

    w.init();
    w.run();

    for s in &w.states {
        eprintln!("{:?}", s);
    }
    eprintln!("{:#?}", w);
}
