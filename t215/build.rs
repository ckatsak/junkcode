fn main() {
    tonic_build::compile_protos("proto/echo.proto").expect("proto compilation");
}
