//! Serializes a Move (only unit variants here) into a JSON file, and deserializes it again.

use std::fs::File;
use std::io::{BufReader, BufWriter};

use anyhow::Result;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
#[serde(tag = "direction")]
enum Move {
    North,
    East,
    South,
    West,
}

fn main() -> Result<()> {
    let a = Move::North;

    // Write it to the file
    {
        let f = File::create("serde1.json")?;
        let wr = BufWriter::new(f);
        serde_json::to_writer_pretty(wr, &a)?;
    }
    // NOTE: Writing occurs in a separate scope so that the data can be flushed to the file as soon
    //       as the Writer is dropped.

    // Read it back from the file
    let f = File::open("serde1.json")?;
    let b: Move = serde_json::from_reader(BufReader::new(f))?;
    eprintln!("b = {:#?}", b);

    Ok(())
}
