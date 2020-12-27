//! Creates a number of random Moves, serializes them into a BSON Vec<u8> and then deserializes
//! them again from it.
//!
//! This takes place in two ways:
//!     1. By serializing them into a BSON array;
//!     2. By serializing them as separate BSON objects, indexed by integers in (0..SIZE).

use std::io::Cursor;

use anyhow::Result;
use bson::{Bson, Document};
use rand::{distributions::Standard, prelude::*};
use serde::{Deserialize, Serialize};

const SIZE: usize = 100;

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

    // Write the Document to the Vec<u8>
    let mut vec = Vec::new();
    doc.to_writer(&mut vec)?; // Apparently, there is no need to wait for a Drop in this case
    eprintln!("Serialization of BSON Array to Vec<u8> is now completed!");

    // Deserialize the Document from the Vec<u8>
    let mut cursor = Cursor::new(vec); // Wrap Vec<u8> with a std::io::Read
    let doc = Document::from_reader(&mut cursor)?;
    // NOTE: We had to use a std::io::Curson because Vec<u8> does not implement std::io::Read.
    // However, [u8] does; therefore, alternatively we could have explicitly coerced the Vec<u8>
    // into a u8 slice, as shown right below. See: https://stackoverflow.com/q/42240663/2304215
    //let doc = Document::from_reader(&mut &vec[..])?; // Alternative to std::io::Cursor
    eprintln!("Deserialized document: {}", doc);
    eprintln!("Deserialization of the BSON Array from Vec<u8> is now completed!");

    //
    // Serialize & deserialize all Moves one-by-one
    //

    // Serialize them into a single Document and into the same bson file
    let mut vec = vec![];
    let mut doc = Document::new();
    moves.iter().enumerate().for_each(|(i, m)| {
        doc.insert(
            i.to_string(),
            bson::to_bson(m).expect("could not convert to Bson"),
        );
        // NOTE: If you move `doc.to_writer()` here, you will have written a separate Document
        // for each Move, each containing all previous moves; hence the deserialization code
        // below will have to loop SIZE iterations to read them all (and again, each of them
        // will be the same as the previous one, but additionally augmented by the latest Move).
    });
    doc.to_writer(&mut vec)
        .expect("could not write Document to BufWriter");
    eprintln!("Serialization (separately) into the Vec<u8> is now completed!");

    // Deserialize the Document(s) from the Vec<u8>
    let mut cursor = Cursor::new(vec);
    let mut i = 0;
    while let Ok(doc) = Document::from_reader(&mut cursor) {
        i += 1;
        eprintln!("iteration {}:\nDeserialized Document: {}", i, doc);
    }
    assert_eq!(i, 1);
    //assert_eq!(i, SIZE);
    // NOTE:  ^^  A separate, new Document would be read on each iteration if the above
    // `doc.to_writer()` had been kept inside the loop -- this is obviously not what we want here.
    eprintln!("Deserialization (separately) from the Vec<u8> is now completed!");

    Ok(())
}
