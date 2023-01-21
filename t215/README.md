# t215

## Run

```console
$ cargo b --release
$ target/release/server
```
```console
$ target/release/client 30000
```

## Configuration

To see the client being throttled at different points down the request sequence:
- play with server's flow control options;
- play with server's internal channel's size;
- ...

## Memory usage

Monitor memory stats for the two processes:

```console
$ hwatch -n 1 -- "grep -E '^(Vm|Rss)' /proc/$(pidof target/release/server)/status"
```

```console
$ hwatch -n 1 -- "grep -E '^(Vm|Rss)' /proc/$(pidof target/release/client)/status"
```
