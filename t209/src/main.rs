use anyhow::{anyhow, Result};
use futures::try_join;
use log::{error, info};
use tokio::signal::unix::{signal, SignalKind};
use tokio::sync::broadcast;
use tokio::time;

#[tokio::main]
async fn main() -> Result<()> {
    env_logger::init();
    info!("[main] PID = {}", std::process::id());

    let mut sigterm = signal(SignalKind::terminate())?;
    let (quit_tx, mut quit_rx1) = broadcast::channel(1);
    let mut quit_rx2 = quit_tx.subscribe();
    let mut quit_rx3 = quit_tx.subscribe();

    // Spawn the signal handling task
    let sh = tokio::spawn(async move {
        sigterm.recv().await;
        info!("[signal-handler] Just received a SIGTERM!");
        quit_tx
            .send(())
            .map_err(|err| anyhow!("broadcasting `()` upon SIGTERM delivery: {}", err))
        //.expect("broadcasting () upon SIGTERM delivery");
    });

    // Spawn dummy tasks
    let t1 = tokio::spawn(async move {
        let mut c = 0;
        let mut ticker = time::interval(std::time::Duration::from_secs(2));
        loop {
            tokio::select! {
                _ = quit_rx1.recv() => {
                    info!("[t1] Just received a () from the signal handling task!");
                    break;
                }
                _ = ticker.tick() => {
                    c += 1;
                    info!("[t1] {} ...", c);
                }
            }
        }
        Ok::<_, anyhow::Error>(())
    });
    let t2 = tokio::spawn(async move {
        let mut c = 0;
        let mut ticker = time::interval(std::time::Duration::from_secs(2));
        loop {
            tokio::select! {
                _ = quit_rx2.recv() => {
                    info!("[t2] Just received a () from the signal handling task!");
                    break;
                }
                _ = ticker.tick() => {
                    c += 1;
                    info!("[t2] {} ...", c);
                }
            }
        }
        Ok::<_, anyhow::Error>(())
    });
    let t3 = tokio::spawn(async move {
        let mut c = 0;
        let mut ticker = time::interval(std::time::Duration::from_secs(2));
        loop {
            tokio::select! {
                _ = quit_rx3.recv() => {
                    info!("[t3] Just received a () from the signal handling task!");
                    break;
                }
                _ = ticker.tick() => {
                    c += 1;
                    info!("[t3] {} ...", c);
                }
            }
        }
        Ok::<_, anyhow::Error>(())
    });

    if let Err(err) = try_join!(sh, t1, t2, t3) {
        error!("One of the tasks terminated unsuccessfully: {}", err);
    }
    Ok(())
}
