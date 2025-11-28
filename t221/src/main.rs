use std::collections::HashMap;

const IN_STR: &str = r#"{"nrow": 76, "ncol": 76}"#;

const PL: [u8; 24] = [
    123, 34, 110, 114, 111, 119, 34, 58, 32, 55, 54, 44, 32, 34, 110, 99, 111, 108, 34, 58, 32, 55,
    54, 125,
];

fn main() -> ::std::io::Result<()> {
    println!("{IN_STR}\nlen == {}\n", IN_STR.len());

    let input = ::serde_json::from_str::<HashMap<String, usize>>(IN_STR).expect("deser str");
    println!("json.loads(IN) == {input:?}\nlen == {}\n", input.len());

    let in_b = ::serde_json::to_vec(&input).expect("ser bytes");
    println!(
        "list(json.dumps(IN).encode()) == {in_b:?}\nlen == {}\n",
        in_b.len()
    );

    let in_str = ::serde_json::to_string(&input).expect("ser str");
    println!("json.dumps(IN)) == {in_str}\nlen == {}\n", in_str.len());

    let pl_dec = ::serde_json::from_slice::<HashMap<String, usize>>(&PL).expect("deser bytes");
    println!(
        "json.loads(bytes(PL)) == {pl_dec:?}\nlen == {}\n",
        pl_dec.len()
    );

    Ok(())
}
