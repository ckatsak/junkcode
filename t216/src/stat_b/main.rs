use std::{cell::RefCell, collections::HashMap, error::Error, rc::Rc};

///////////////////////////////////////////////////////////////////////////////////////////////////
// Strategy
///////////////////////////////////////////////////////////////////////////////////////////////////

pub trait Strategy {
    fn apply(&mut self, yolo_map: &mut HashMap<u64, String>) -> u64
    where
        Self: Sized;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// S1
///////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Debug, Clone, Default)]
pub struct S1 {
    count: u64,
    s: String,
}

impl Strategy for S1 {
    fn apply(&mut self, yolo_map: &mut HashMap<u64, String>) -> u64 {
        self.count += 1;
        yolo_map.insert(self.count, format!("{}", self.count));
        self.s = format!("{}", yolo_map.len());
        self.count
    }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// S2
///////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Debug, Clone, Default)]
pub struct S2 {
    count: u64,
    s: String,
}

impl Strategy for S2 {
    fn apply(&mut self, yolo_map: &mut HashMap<u64, String>) -> u64 {
        self.count += 1;
        yolo_map.insert(self.count, format!("{}", 2 * self.count));
        self.s = format!("{}", 2 * yolo_map.len());
        2 * self.count
    }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// Yolo
///////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Debug, Clone, Default)]
pub struct Yolo<S: Strategy> {
    m: HashMap<u64, String>,

    s: Rc<RefCell<S>>,
}

impl<S: Strategy> Yolo<S> {
    pub fn strategy(&mut self) -> u64 {
        let s = self.s.clone();
        let ret = s.borrow_mut().apply(&mut self.m);
        ret
    }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// main
///////////////////////////////////////////////////////////////////////////////////////////////////

fn main() -> Result<(), Box<dyn Error + 'static>> {
    let mut y1 = Yolo::<S1>::default();
    let mut y2 = Yolo::<S2>::default();

    for i in 0..10 {
        assert!(y2.strategy() == 2 * y1.strategy());

        println!("{i:02}:\ny1 = {y1:#?}\ny2 = {y2:#?}\n\n");
    }

    Ok(())
}
