//! Serializes a Move into JSON and RON files, but also into Vec<u8>, and deserializes them again.

use std::fs::File;
use std::io::BufWriter;

use anyhow::Result;
use ron::ser::PrettyConfig;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
#[serde(tag = "direction", content = "by")]
enum Move {
    Up(usize),
    Down(usize),
    Left(usize),
    Right(usize),
}

fn main() -> Result<()> {
    let a = Move::Left(42);

    // JSON & RON: String & Vec
    let jstr = serde_json::to_string_pretty(&a)?;
    let jvec = serde_json::to_vec(&a)?;
    let rstr = ron::ser::to_string_pretty(&a, PrettyConfig::default())?;
    let rvec: Vec<u8> = rstr.clone().into();
    eprintln!("jstr = {}", jstr);
    eprintln!("jvec = {:?}", jvec);
    eprintln!("rstr = {}", rstr);
    eprintln!("rvec = {:?}", rvec);

    // JSON & RON: Serialize to file
    let jf = File::create("serde2.json")?;
    let rf = File::create("serde2.ron")?;
    {
        let jbw = BufWriter::new(jf);
        let rbw = BufWriter::new(rf);

        serde_json::to_writer_pretty(jbw, &a)?;
        ron::ser::to_writer_pretty(rbw, &a, PrettyConfig::default())?;
    }

    // JSON & RON: Deserialize from file
    let jf = File::open("serde2.json")?;
    let rf = File::open("serde2.ron")?;
    let jb: Move = serde_json::from_reader(jf)?;
    let rb: Move = ron::de::from_reader(rf)?;
    eprintln!("jb = {:?}", jb);
    eprintln!("rb = {:?}", rb);

    Ok(())
}
