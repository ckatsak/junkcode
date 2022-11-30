//! Generics as function arguments: trait bounds & `impl` notation.

use std::path::{Path, PathBuf};

fn a<A: AsRef<Path>>(p1: A, p2: A) {
    eprintln!("a: {}, {}", p1.as_ref().display(), p2.as_ref().display())
}

fn b<A: AsRef<Path>, B: AsRef<Path>>(p1: A, p2: B) {
    eprintln!("b: {}, {}", p1.as_ref().display(), p2.as_ref().display())
}

fn c(p1: impl AsRef<Path>, p2: impl AsRef<Path>) {
    eprintln!("c: {}, {}", p1.as_ref().display(), p2.as_ref().display())
}

fn main() {
    a("/p/1/", &String::from("/p/2/"));
    a(&PathBuf::from("/p/1/"), &PathBuf::from("/p/2/"));
    //a(&PathBuf::from("/p/1/"), Path::new("/p/2/"));
    //-                          ^^^^^^^^^^^^^^^^^^ expected struct `PathBuf`, found struct `Path`
    //|
    //arguments to this function are incorrect
    //
    //= note: expected reference `&PathBuf`
    //        found reference `&Path`
    //
    //a("/p/1/", &PathBuf::from("/p/2/"));
    //-          ^^^^^^^^^^^^^^^^^^^^^^^ expected `str`, found struct `PathBuf`
    //|
    //arguments to this function are incorrect
    //
    //= note: expected reference `&str`
    //           found reference `&PathBuf`

    b("/p/1/", &String::from("/p/2/"));
    b(&PathBuf::from("/p/1/"), &PathBuf::from("/p/2/"));
    b(&PathBuf::from("/p/1/"), Path::new("/p/2/"));
    b("/p/1/", &PathBuf::from("/p/2/"));

    c("/p/1/", &String::from("/p/2/"));
    c(&PathBuf::from("/p/1/"), &PathBuf::from("/p/2/"));
    c(&PathBuf::from("/p/1/"), Path::new("/p/2/"));
    c("/p/1/", &PathBuf::from("/p/2/"));
}
