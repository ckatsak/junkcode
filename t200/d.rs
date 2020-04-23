//use std::cell::RefCell;
use std::fmt;
//use std::rc::Rc;

trait State: fmt::Debug {
    fn process(&mut self) -> bool;
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
    fn process(&mut self) -> bool {
        true
    }
}

impl State for B {
    fn process(&mut self) -> bool {
        true
    }
}

impl State for C {
    fn process(&mut self) -> bool {
        true
    }
}

#[derive(Debug)]
struct StateMachine<S: State> {
    count: u8,
    state: S,
}

impl<S: State> StateMachine<S> {
    fn process(&mut self) -> bool {
        self.state.process()
    }
}

impl StateMachine<A> {
    fn new(count: u8) -> Self {
        StateMachine {
            count: count,
            state: A {},
        }
    }
}

impl From<StateMachine<A>> for StateMachine<B> {
    fn from(old: StateMachine<A>) -> StateMachine<B> {
        StateMachine::<B> {
            count: old.count + 1,
            state: B(42),
        }
    }
}

impl From<StateMachine<B>> for StateMachine<C> {
    fn from(old: StateMachine<B>) -> StateMachine<C> {
        StateMachine::<C> {
            count: old.count + 1,
            state: C { x: 69 },
        }
    }
}

impl From<StateMachine<C>> for StateMachine<A> {
    fn from(old: StateMachine<C>) -> StateMachine<A> {
        StateMachine::<A> {
            count: old.count + 1,
            state: A {},
        }
    }
}

#[derive(Debug)]
struct W {
    count: u8,
}

impl W {
    fn new(count: u8) -> Self {
        W { count }
    }

    fn run(&mut self) {
        eprintln!("{:?}", self);

        let mut sm = StateMachine::new(0);
        eprintln!("{:#?}", sm);
    }
}

fn main() {
    let mut w = W::new(0);

    w.run();

    eprintln!("{:#?}", w);
}
