use std::sync::{Arc, Mutex};
use std::time::Duration;

use anyhow::{anyhow, Result};
use futures::try_join;
use tokio::{
    signal::unix::{signal, Signal, SignalKind},
    sync::broadcast::{self, Receiver, Sender},
    time,
};
use tracing::{error, info};

#[tracing::instrument(skip(counter, quit_rx))]
async fn task_loop(id: usize, counter: Arc<Mutex<i32>>, mut quit_rx: Receiver<()>) -> Result<()> {
    let mut ticker = time::interval(Duration::from_secs(2));
    loop {
        tokio::select! {
            _ = quit_rx.recv() => {
                info!("Just received a () from the signal handling task!");
                break;
            }
            _ = ticker.tick() => {
                let mut counter = counter.lock().unwrap();
                *counter += 1;
                info!("{:?} ...", counter);
            }
        }
    }
    Ok::<_, anyhow::Error>(())
}

#[tracing::instrument(skip(sigterm, quit_tx))]
async fn sigterm_handler(mut sigterm: Signal, quit_tx: Sender<()>) -> Result<usize> {
    sigterm.recv().await;
    info!("Just received a SIGTERM!");
    quit_tx
        .send(())
        .map_err(|err| anyhow!("broadcasting `()` upon SIGTERM delivery: {}", err))
}

#[tokio::main]
async fn main() -> Result<()> {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::TRACE)
        .init();
    info!("[main] PID = {}", std::process::id());

    let cnt = Arc::new(Mutex::new(0));

    // Set up the signal handler...
    let sigterm = signal(SignalKind::terminate())?;
    // ...and a channel for the signal handling task to broadcast to the rest of the tasks.
    let (quit_tx, quit_rx1) = broadcast::channel(1);
    let quit_rx2 = quit_tx.subscribe();
    let quit_rx3 = quit_tx.subscribe();

    // Spawn the signal handling task
    let sh = tokio::spawn(async move { sigterm_handler(sigterm, quit_tx).await });

    // Spawn dummy tasks
    let t1 = tokio::spawn({
        let c = Arc::clone(&cnt);
        async move { task_loop(1, c, quit_rx1).await }
    });
    let t2 = tokio::spawn({
        let c = Arc::clone(&cnt);
        async move { task_loop(2, c, quit_rx2).await }
    });
    let t3 = tokio::spawn({
        let c = Arc::clone(&cnt);
        async move { task_loop(3, c, quit_rx3).await }
    });

    if let Err(err) = try_join!(sh, t1, t2, t3) {
        error!("One of the tasks terminated unsuccessfully: {}", err);
    }
    Ok(())
}
