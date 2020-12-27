//! Creates a number of random Moves, serializes them into a CBOR file and then deserializes them
//! again from the file.
//!
//! This takes place in two ways:
//!     1. By serializing them into a BSON array (serde3.bson);
//!     2. By serializing them as separate BSON objects, indexed by integers in (0..SIZE)
//!        (serde3b.bson).

use std::fs::File;
use std::io::Write;

use anyhow::Result;
use rand::{distributions::Standard, prelude::*};
use serde::{Deserialize, Serialize};

const SIZE: usize = 100;
const BUF_SZ: usize = 1 << 4;

#[derive(Debug, Clone, Copy, Serialize, Deserialize, Eq, PartialEq, Hash)]
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
    // Serialize & deserialize all Moves one-by-one
    //

    let mut f = File::create("serde5.cbor")?;
    for m in moves.iter() {
        eprintln!("to_writer({:?})", m);
        serde_cbor::to_writer(&f, m)?;
    }
    f.flush()?;
    eprintln!("Serialization to 'serde5.cbor' has now been completed!");

    let mut h = std::collections::HashSet::with_capacity(SIZE);
    let f = File::open("serde5.cbor")?;

    //while let Ok(datum) = serde_cbor::from_reader(&f) {
    //    eprintln!("Just deserialized datum {:?}", datum);
    //    h.insert(datum);
    //}
    let de = serde_cbor::Deserializer::from_reader(&f);
    for datum in de.into_iter::<Move>() {
        let datum = datum.unwrap();
        eprintln!("Just deserialized datum {:?}", datum);
        h.insert(datum);
    }

    h.iter().for_each(|datum| {
        assert!(moves.contains(&datum));
        eprintln!("'moves' indeed contains {:?}", datum);
    });
    eprintln!("Deserialization from 'serde5.cbor' has now been completed!");

    //
    //    // Serialize them into a single Document and into the same bson file
    //    let f = File::create("serde3b.bson")?;
    //    {
    //        let mut bw = BufWriter::new(f);
    //        let mut doc = Document::new();
    //        moves.iter().enumerate().for_each(|(i, m)| {
    //            //eprintln!("(i, m) = {:?}", (i, m));
    //            doc.insert(
    //                i.to_string(),
    //                bson::to_bson(m).expect("could not convert to Bson"),
    //            );
    //            // NOTE: If you move `doc.to_writer()` here, you will have written a separate Document
    //            // for each Move, each containing all previous moves; hence the deserialization code
    //            // below will have to loop SIZE iterations to read them all (and again, each of them
    //            // will be the same as the previous one, but additionally augmented by the latest Move)
    //        });
    //        doc.to_writer(&mut bw)
    //            .expect("could not write Document to BufWriter");
    //    }
    //    eprintln!("Serialization to 'serde3b.bson' is now completed!");
    //
    //    // Deserialize the Document(s) from the file
    //    let f = File::open("serde3b.bson")?;
    //    let mut br = BufReader::with_capacity(BUF_SZ, f);
    //    let mut i = 0;
    //    while let Ok(doc) = Document::from_reader(&mut br) {
    //        i += 1;
    //        eprintln!("iteration {}:\nDeserialized Document: {}", i, doc);
    //    }
    //    assert_eq!(i, 1);
    //    //assert_eq!(i, SIZE);
    //    // NOTE:  ^^  A separate, new Document would be read on each iteration if the above
    //    // `doc.to_writer()` had been kept inside the loop -- this is obviously not what we want here.
    //    eprintln!("Deserialization from 'serde3b.bson' is now completed!");
    //

    Ok(())
}
