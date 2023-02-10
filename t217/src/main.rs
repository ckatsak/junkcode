use std::{collections::HashMap, error::Error, time::Duration};

use humantime::format_duration;
use rand::{
    distributions::{Alphanumeric, DistString},
    thread_rng,
};
use tokio::{task::JoinSet, time::Instant};

const ROUNDS: usize = 1 << 30;

#[::tokio::main]
async fn main() -> Result<(), Box<dyn Error + 'static>> {
    let num_cpus = num_cpus::get();
    let rounds_per_task: usize = ROUNDS / num_cpus * 4;

    let mut set = JoinSet::new();
    for _ in 0..num_cpus / 4 {
        set.spawn(::tokio::spawn(async move {
            let mut m = HashMap::with_capacity(rounds_per_task);

            let mut rng = thread_rng();
            let dist = Alphanumeric;

            let mut total_time = Duration::ZERO;
            for _ in 0..rounds_per_task {
                let t_s = Instant::now();
                let s = dist.sample_string(&mut rng, 15);
                total_time += Instant::now() - t_s;

                m.entry(s).and_modify(|count| *count += 1).or_insert(1u64);
            }

            //m.retain(|_s, &mut count| count > 1); // NOTE: not here; after aggregating them all
            (m, total_time / rounds_per_task as u32)
        }));
    }

    let mut aggr_m = HashMap::with_capacity(ROUNDS + num_cpus);
    while let Some(join_res) = set.join_next().await {
        if let Err(join_err) = join_res {
            eprintln!("external join error: {join_err}");
            continue;
        }
        let join_res = join_res.unwrap();
        if let Err(join_err) = join_res {
            eprintln!("internal join error: {join_err}");
            continue;
        }
        let (m, dur) = join_res.unwrap();
        eprintln!("Average time per sample = {}", format_duration(dur),);
        merge_map(&mut aggr_m, m);
    }

    aggr_m.retain(|_s, &mut count| count > 1);
    eprintln!("# collisions = {}", aggr_m.len());
    eprintln!("Collisions: {aggr_m:#?}");

    Ok(())
}

fn merge_map(acc: &mut HashMap<String, u64>, m: HashMap<String, u64>) {
    eprintln!("len(acc) = {}; len(m) = {}", acc.len(), m.len());
    acc.reserve(m.len());
    m.into_iter().for_each(|(s, new_c)| {
        acc.entry(s)
            .and_modify(|acc_c| *acc_c += new_c)
            .or_insert(new_c);
    });
    eprintln!("len(acc) = {}\n", acc.len());
}
