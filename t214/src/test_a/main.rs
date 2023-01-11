use std::{sync::Arc, time::Duration};

use anyhow::{anyhow, Context, Result};
use tokio::{sync::Semaphore, time::sleep};
use tracing::{debug, info};
use tracing_subscriber::{fmt::format::FmtSpan, EnvFilter};

#[tokio::main]
async fn main() -> Result<()> {
    ::tracing_subscriber::fmt()
        .with_test_writer()
        .with_env_filter(
            EnvFilter::from_default_env().add_directive(
                "test_a=trace"
                    .parse()
                    .with_context(|| "failed to parse filtering directive")?,
            ),
        )
        .with_thread_ids(true)
        .with_span_events(FmtSpan::NEW | FmtSpan::CLOSE)
        .try_init()
        .map_err(|err| anyhow!("failed to initialize tracing subscriber: {err}"))?;

    ///////////////////////////////////////////////////////////////////////////////////////////////

    let semaphore = Arc::new(Semaphore::new(3));
    let mut join_handles = Vec::new();

    for i in 0..10 {
        info!("[main] iteration {i}");
        let permit = semaphore.clone().acquire_owned().await.unwrap();
        join_handles.push(tokio::spawn(async move {
            // perform task...
            debug!("[task {i}] hey!");
            sleep(Duration::from_secs(2)).await;
            //
            drop(permit);
        }));
    }

    for handle in join_handles {
        handle.await.unwrap();
    }

    Ok(())
}
