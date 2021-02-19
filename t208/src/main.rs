enum Adj {
    Pred = -1,
    Succ = 1,
}

fn main() {
    eprintln!("\nPositive");
    let x: isize = 7;
    eprintln!("x = {:#?}", x);
    eprintln!("x % 2 = {:#?}", x % 2);
    eprintln!("x % (-2) = {:#?}", x % (-2));
    eprintln!("x % 3 = {:#?}", x % 3);
    eprintln!("x % (-3) = {:#?}", x % (-3));
    eprintln!("x % 4 = {:#?}", x % 4);
    eprintln!("x % (-4) = {:#?}", x % (-4));

    eprintln!("\nNegative");
    let x: isize = -7;
    eprintln!("x = {:#?}", x);
    eprintln!("x % 2 = {:#?}", x % 2);
    eprintln!("x % (-2) = {:#?}", x % (-2));
    eprintln!("x % 3 = {:#?}", x % 3);
    eprintln!("x % (-3) = {:#?}", x % (-3));
    eprintln!("x % 4 = {:#?}", x % 4);
    eprintln!("x % (-4) = {:#?}", x % (-4));

    eprintln!("\nAdj");
    let x: isize = 0 + (Adj::Pred as isize);
    eprintln!("x = {:#?}", x);
    eprintln!("x % 2 = {:#?}", x % 2);
    eprintln!("x % (-2) = {:#?}", x % (-2));
    eprintln!("x % 3 = {:#?}", x % 3);
    eprintln!("x % (-3) = {:#?}", x % (-3));
    eprintln!("x % 4 = {:#?}", x % 4);
    eprintln!("x % (-4) = {:#?}", x % (-4));

    let index = 2;
    eprintln!("\nStraight Iterator");
    let v = vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
    let (mut j, len) = (0, v.len());
    let iter = v.iter().enumerate().cycle().skip(index + 1);
    for (i, v) in iter {
        if j == len {
            break;
        }
        eprintln!("i = {}, v = {}", i, v);
        j += 1;
    }

    eprintln!("\nReverse Iterator");
    let v = vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
    let (mut j, len) = (0, v.len());
    let iter = v.iter().rev().enumerate().cycle().skip(len - index);
    for (i, v) in iter {
        if j == len {
            break;
        }
        eprintln!("i = {}, v = {}", i, v);
        j += 1;
    }
}
