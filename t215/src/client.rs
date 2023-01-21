use std::task::Poll;

use futures::stream::poll_fn;
use tonic::{transport::Channel, Request};
use tracing::{error, info, warn};

use t215::{
    pb::{echo_client::EchoClient, EchoRequest},
    ADDRESS,
};

async fn do_it(c: &mut EchoClient<Channel>, count: u64) -> anyhow::Result<()> {
    let mut curr = 0;
    let req = Request::new(poll_fn(move |_| -> Poll<Option<EchoRequest>> {
        if curr == count {
            warn!("Done sending requests!");
            Poll::Ready(None)
        } else {
            warn!("Sending request id = {curr} ...");
            curr += 1;
            Poll::Ready(Some(EchoRequest { id: curr - 1 }))
        }
    }));

    match c.client_streaming_echo(req).await {
        Ok(resp) => {
            //info!("Just received {resp:?} from the server");
            let resp = resp.into_inner();
            info!("Received response for id = {}", resp.id);
        }
        Err(err) => {
            error!("Received error from server: {err}");
        }
    }

    Ok(())
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let count: u64 = std::env::args()
        .nth(1)
        .expect("usage: client <count>")
        .parse()
        .expect("failed to parse u64 from argv[1]");

    tracing_subscriber::fmt()
        .with_target(true)
        .with_thread_ids(true)
        .with_level(true)
        .with_file(true)
        .with_line_number(true)
        .init();

    let mut client = EchoClient::connect(format!("http://{ADDRESS}")).await?;
    do_it(&mut client, count).await?;

    Ok(())
}
