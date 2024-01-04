use std::time::Duration;

use anyhow::{Context, Result};
use tokio::{
    sync::{broadcast, mpsc},
    time::sleep,
};

async fn task(mut quit_rx: broadcast::Receiver<()>, mut rx: mpsc::Receiver<u8>) {
    eprintln!("[task] Spawning...");

    loop {
        tokio::select! {
            //biased;
            // NOTE:
            // - If we make this `biased`, this select will always check `quit_rx` first. So
            // once `quit_rx` is closed, `quit_rx.recv()` will always return `Err(Closed)`, thus
            // starving `rx.recv()` and making the loop infinite.
            // - Commenting out the `biased` mark, all branches are checked "fairly", so `rx` is
            // not starved and the loop finishes when `rx.recv()` returns `None`.

            res = quit_rx.recv() => {
                eprintln!("[task] QUIT CHANNEL: {res:?}");
            }

            n = rx.recv() => {
                eprintln!("[task] rx: received {n:?}");
                match n {
                    Some(n) => eprintln!("[task] num {n:03}"),
                    None => break,
                }
            }

        }
    }

    eprintln!("[task] Exiting...");
}

#[tokio::main]
async fn main() -> Result<()> {
    let (tx, rx) = mpsc::channel(4);
    let (quit_tx, quit_rx) = broadcast::channel::<()>(1);

    eprintln!("[main] Spawning task...");
    let h = tokio::spawn(async move { task(quit_rx, rx).await });
    eprintln!("[main] Spawned task");

    play(quit_tx, tx).await.context("play failed")?;

    h.await.context("[main] Failed to join task")
}

async fn play(quit_tx: broadcast::Sender<()>, tx: mpsc::Sender<u8>) -> Result<()> {
    for i in 0..5 {
        sleep(Duration::from_secs(1)).await;
        tx.send(i)
            .await
            .with_context(|| format!("failed to send {i:03?}"))?;
    }

    quit_tx.send(()).context("failed to send quit signal")?;

    for i in 5..10 {
        sleep(Duration::from_millis(500)).await;
        tx.send(i)
            .await
            .with_context(|| format!("failed to send {i:03?}"))?;
    }

    Ok(())
}
