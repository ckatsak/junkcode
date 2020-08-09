//! Build and run:
//!
//!     $ rustc t201.rs && ./t201
//!
//! 1. Test comparison of &str.
//!
//! 2. Test derivation of Default trait for types that contain &str.

pub static s1: &'static str = "yolo";
pub static s2: &'static str = "yolo";
pub static s3: &'static str = "kolo";
pub static s4: &'static str = "kOLo";

#[derive(Debug, Default)]
struct St<'a> {
    s: &'a str,
    x: usize,
}

fn main() {
    eprintln!("s1: {}\ns2: {}\n s1 == s2 : {}\n", s1, s2, s1 == s2);
    eprintln!("s1: {}\ns3: {}\n s1 == s3 : {}\n", s1, s3, s1 == s3);
    eprintln!("s1: {}\ns4: {}\n s1 == s4 : {}\n", s1, s4, s1 == s4);
    eprintln!("s2: {}\ns3: {}\n s2 == s3 : {}\n", s2, s3, s2 == s3);
    eprintln!("s2: {}\ns4: {}\n s2 == s4 : {}\n", s2, s4, s2 == s4);
    eprintln!("s3: {}\ns4: {}\n s3 == s4 : {}\n", s3, s4, s2 == s4);
    eprintln!(
        "s3: {}\ns4: {}\n s3.eq_..._case(s4) : {}\n",
        s3,
        s4,
        s3.eq_ignore_ascii_case(s4)
    );

    let st = St::default();
    assert_eq!(st.s, "");
    eprintln!("St::default() = {:?}", st);
}
