pub mod pb {
    tonic::include_proto!("grpc.examples.echo");
}

pub const ADDRESS: &str = "127.0.0.1:50052";
