//! Test std::path::PathBuf.extension
//!
//! $ rustc t197.rs
//! $ ./t197

use std::path::PathBuf;

fn main() {
    let dez = vec![
        PathBuf::from("A.a"),
        PathBuf::from("B.b.b"),
        PathBuf::from("C"),
        PathBuf::from("D_d"),
        PathBuf::from(".E"),
        PathBuf::from(".F.f_f"),
        PathBuf::from("G.g_g"),
        PathBuf::from(".H.h."),
        PathBuf::from("I.i."),
        PathBuf::from("J."),
    ];

    for e in &dez {
        eprintln!("{:?}\t--> {:?}", e, e.extension());
    }
}
