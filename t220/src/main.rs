fn main() {
    let mut ret = vec![(0, 0)];

    let mut v = Vec::new();
    let mut prev_cap = 0;
    for i in 0u32..1 << 31 {
        v.push((i & 0xff) as u8);

        let curr_cap = v.capacity();
        if curr_cap != prev_cap {
            ret.push((v.len(), curr_cap));
            prev_cap = curr_cap;
        }
    }

    eprintln!("[(length, cap)] = {ret:?}");
}
