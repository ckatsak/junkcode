use std::{net::ToSocketAddrs, time::Duration};

use rand::{rngs::SmallRng, Rng, SeedableRng};
use tokio::{
    sync::mpsc::{self, Sender},
    time::sleep,
};
use tokio_stream::StreamExt;
use tonic::{transport::Server, Request, Response, Status, Streaming};
use tracing::{error, info};

use t215::{
    pb::{self, EchoRequest, EchoResponse},
    ADDRESS,
};

const PROCESSOR_CHANNEL_SIZE: usize = 2;

struct SlowServer {
    tx: Sender<EchoRequest>,
}

#[tonic::async_trait]
impl pb::echo_server::Echo for SlowServer {
    async fn client_streaming_echo(
        &self,
        request: Request<Streaming<pb::EchoRequest>>,
    ) -> Result<Response<pb::EchoResponse>, Status> {
        let mut cnt = 0u64;

        let mut stream = request.into_inner();
        while let Some(req) = stream.next().await {
            info!("Just received {req:?} from the stream");
            let req = match req {
                Ok(req) => req,
                Err(err) => {
                    error!("Failed to retrieve an EchoRequest from the incoming Request: {err}");
                    sleep(Duration::from_millis(500)).await;
                    continue;
                }
            };
            info!("Sending {req:?} down the processor's channel...");
            match self.tx.send(req).await {
                Ok(()) => {}
                Err(err) => {
                    error!("Failed to send the request down the processor's channel: {err}");
                    sleep(Duration::from_millis(500)).await;
                    continue;
                }
            }
            cnt += 1;
            info!("Served {cnt} requests so far...");
        }

        Ok(Response::new(EchoResponse { id: cnt }))
    }
}

#[tokio::main]
async fn main() -> Result<(), tonic::transport::Error> {
    tracing_subscriber::fmt()
        .with_target(true)
        .with_thread_ids(true)
        .with_level(true)
        .with_file(true)
        .with_line_number(true)
        .init();

    let (tx, mut rx) = mpsc::channel::<EchoRequest>(PROCESSOR_CHANNEL_SIZE);
    let _handle = tokio::spawn(async move {
        let mut rng = SmallRng::from_entropy();
        while let Some(req) = rx.recv().await {
            info!("Processing id = {} ...", req.id);
            sleep(Duration::from_millis(rng.gen_range(1000..2000u64))).await;
            info!("Done processing id = {}", req.id);
        }
    });

    Server::builder()
        .max_concurrent_streams(1)
        //.initial_stream_window_size(65_535)
        //.initial_connection_window_size(65_535)
        .initial_stream_window_size(4096)
        .initial_connection_window_size(4096)
        .max_frame_size(None)
        .add_service(pb::echo_server::EchoServer::new(SlowServer { tx }))
        .serve(
            ADDRESS
                .to_socket_addrs()
                .expect("to_socket_addrs()")
                .next()
                .expect("to_socket_addrs().next()"),
        )
        .await
}
