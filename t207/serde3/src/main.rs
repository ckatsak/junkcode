//! Creates a number of random Moves, serializes them into a BSON file and then deserializes them
//! again from the file.
//!
//! This takes place in two ways:
//!     1. By serializing them into a BSON array (serde3.bson);
//!     2. By serializing them as separate BSON objects, indexed by integers in (0..SIZE)
//!        (serde3b.bson).

use std::fs::File;
use std::io::{BufReader, BufWriter};

use anyhow::Result;
use bson::{Bson, Document};
use rand::{distributions::Standard, prelude::*};
use serde::{Deserialize, Serialize};

const SIZE: usize = 100;
const BUF_SZ: usize = 1 << 4;

#[derive(Debug, Clone, Copy, Serialize, Deserialize)]
#[serde(tag = "direction", content = "by")]
enum Move {
    North(i8),
    East(i8),
    South(i8),
    West(i8),
}

impl Distribution<Move> for Standard {
    fn sample<R: Rng + ?Sized>(&self, rng: &mut R) -> Move {
        let by = rng.gen();
        match rng.gen::<u8>() % 4 {
            0 => Move::North(by),
            1 => Move::East(by),
            2 => Move::South(by),
            3 => Move::West(by),
            _ => unreachable!(),
        }
    }
}

fn main() -> Result<()> {
    // Create the random Moves
    let mut moves = Vec::<Move>::with_capacity(SIZE);
    for _ in 0..SIZE {
        moves.push(thread_rng().gen());
    }

    //
    // Serialize & deserialize all Moves into a single Bson::Array
    //

    // Serialize the Moves into an Array, and insert the array into a new Document
    let mut doc = Document::new();
    let mut arr = bson::Array::with_capacity(SIZE);
    moves
        .iter()
        .for_each(|m| arr.push(bson::to_bson(m).unwrap()));
    doc.insert("array".to_string(), Bson::Array(arr));

    // Write the Document to the file
    {
        let f = File::create("serde3.bson")?;
        let mut bw = BufWriter::new(f);
        doc.to_writer(&mut bw)?;
    }
    eprintln!("Serialization to 'serde3.bson' is now completed!");

    // Open the file and deserialize the Document here
    let f = File::open("serde3.bson")?;
    let mut br = BufReader::with_capacity(BUF_SZ, f);
    let doc = Document::from_reader(&mut br)?;
    eprintln!("Deserialized document: {}", doc);
    //loop {
    //    match Document::from_reader(&mut br) {
    //        Ok(de) => {
    //            eprintln!("Deserialized chunk: {:?}", de);
    //        }
    //        Err(e) => {
    //            eprintln!("Gamithike pali to Document::from_reader(&br): '{}'", e);
    //            break;
    //        }
    //    }
    //}
    eprintln!("Deserialization from 'serde3.bson' is now completed!");

    //
    // Serialize & deserialize all Moves one-by-one
    //

    // Serialize them into a single Document and into the same bson file
    let f = File::create("serde3b.bson")?;
    {
        let mut bw = BufWriter::new(f);
        let mut doc = Document::new();
        moves.iter().enumerate().for_each(|(i, m)| {
            //eprintln!("(i, m) = {:?}", (i, m));
            doc.insert(
                i.to_string(),
                bson::to_bson(m).expect("could not convert to Bson"),
            );
            // NOTE: If you move `doc.to_writer()` here, you will have written a separate Document
            // for each Move, each containing all previous moves; hence the deserialization code
            // below will have to loop SIZE iterations to read them all (and again, each of them
            // will be the same as the previous one, but additionally augmented by the latest Move)
        });
        doc.to_writer(&mut bw)
            .expect("could not write Document to BufWriter");
    }
    eprintln!("Serialization to 'serde3b.bson' is now completed!");

    // Deserialize the Document(s) from the file
    let f = File::open("serde3b.bson")?;
    let mut br = BufReader::with_capacity(BUF_SZ, f);
    let mut i = 0;
    while let Ok(doc) = Document::from_reader(&mut br) {
        i += 1;
        eprintln!("iteration {}:\nDeserialized Document: {}", i, doc);
    }
    assert_eq!(i, 1);
    //assert_eq!(i, SIZE);
    // NOTE:  ^^  A separate, new Document would be read on each iteration if the above
    // `doc.to_writer()` had been kept inside the loop -- this is obviously not what we want here.
    eprintln!("Deserialization from 'serde3b.bson' is now completed!");

    Ok(())
}
