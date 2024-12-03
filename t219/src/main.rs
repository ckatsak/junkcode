use std::{
    any::type_name,
    time::{Duration, SystemTime},
};

use humantime::format_duration;
use num_traits::{PrimInt, ToPrimitive};

fn test_type_for_nanoseconds<T: PrimInt>() {
    eprintln!("  **\t{}\t**", type_name::<T>());
    let max_as_u64 = T::max_value().to_u64().expect("T::MAX.to_u64()");
    let max_as_u128 = T::max_value().to_u128().expect("T::MAX.to_u128()");
    eprintln!(" - max_as_u64  = {max_as_u64}\n - max_as_u128 = {max_as_u128}");

    eprintln!(
        " - {}::MAX nanoseconds == {}",
        type_name::<T>(),
        format_duration(Duration::from_nanos(max_as_u64)),
    );

    let nanos_epoch = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap()
        .as_nanos();
    eprintln!(
        " - Since EPOCH: {nanos_epoch}ns == {}",
        format_duration(Duration::from_nanos(nanos_epoch as _)),
    );

    if nanos_epoch > max_as_u128 {
        let since = Duration::from_nanos(
            (nanos_epoch - max_as_u128)
                .to_u64()
                .expect("(EPOCH-max).to_u64()"),
        );
        eprintln!(
            " - Nanos since EPOCH cannot fit in a {} for the past {}!",
            type_name::<T>(),
            format_duration(since),
        );
    } else {
        let rem = Duration::from_nanos(
            (max_as_u128 - nanos_epoch)
                .to_u64()
                .expect("(max-EPOCH).to_u64()"),
        );
        eprintln!(
            " - Time until EPOCH cannot fit in {}: {}",
            type_name::<T>(),
            format_duration(rem),
        );
    }
}

fn main() {
    //eprintln!("---------------------------------------------------------------------------------");
    //test_u64();
    eprintln!("---------------------------------------------------------------------------------");
    test_type_for_nanoseconds::<u64>();
    eprintln!("---------------------------------------------------------------------------------");
    test_type_for_nanoseconds::<u32>();
    //eprintln!("---------------------------------------------------------------------------------");
    //test_type_for_nanoseconds::<u16>();
    //eprintln!("---------------------------------------------------------------------------------");
    //test_type_for_nanoseconds::<u8>();
    eprintln!("---------------------------------------------------------------------------------");
    test_type_for_nanoseconds::<i64>();
    eprintln!("---------------------------------------------------------------------------------");
    test_type_for_nanoseconds::<i32>();
    //eprintln!("---------------------------------------------------------------------------------");
    //test_type_for_nanoseconds::<i16>();
    //eprintln!("---------------------------------------------------------------------------------");
    //test_type_for_nanoseconds::<i8>();
    eprintln!("---------------------------------------------------------------------------------");
}

#[allow(dead_code)]
fn test_u64() {
    let d = Duration::from_nanos(u64::MAX).as_secs_f64();
    eprintln!(
        " - u64::MAX nanoseconds == {:.3} years",
        d / 60. / 60. / 24. / 365.25
    );

    let nanos_epoch = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap()
        .as_nanos();
    eprintln!(
        " - {nanos_epoch}ns == {:.3} years since EPOCH",
        Duration::from_nanos(nanos_epoch as _).as_secs_f64() / 3600. / 24. / 365.25
    );
    if nanos_epoch > u64::MAX as _ {
        eprintln!(" - Nanos since EPOCH cannot fit in a u64!")
    } else {
        let rem = Duration::from_nanos(u64::MAX - nanos_epoch as u64).as_secs_f64()
            / 60.
            / 60.
            / 24.
            / 365.25;
        eprintln!(" - {rem:.3} years remaining until EPOCH cannot fit in u64");
    }
}
