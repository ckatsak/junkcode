# t180


## Problem Description

```text
$ go run prob.go
```

In `prob.go`, the main-goroutine produces some workload (`int`s,
representing a "job") and fills a buffered channel (to be used as a
blocking queue for multiple consumer-goroutines).

The channel is not closed by main-goroutine, because processing of a
job (by a consumer-goroutine) may fail, in which case it has to be
re-enqueued (re-sent to the end of the buffered channel) to allow for
attempting to process it again later (either by the same or by some
other consumer-goroutine).
A job may fail and need to be re-enqueued many times (however, in this
example it is re-enqueued exactly once per consumer-goroutine that
grabs it).

If the time-waster was not up, the runtime would detect a deadlock, and
panic, which is convenient for such a small example.
However, in my actual project, there are multiple goroutines up, hence
the deadlock is not even detected at all.
The role of the time-waster is to just be up doing nothing other than
reproducing this inability to detect the deadlock, and we don't care
what happens to it after `THE END` is printed out.

**NOTE** regarding `prob.go`: the `bool` returned when receiving from
a channel, doesn't help at all, unless the channel has been closed.
But the channel cannot be closed in this case, because there are `N+1`
senders and `N` receivers (where `N` is the number of
consumer-goroutines).


## tl;dr GOAL

Modify `prob.go` so that when all the workload is properly processed by
the consumer-goroutines, the main-goroutine prints `THE END` to stdout.


## Solutions?

### 1. Using `len()` on the chan

Yeah.. I know, normally `len()` is weird on channels, and stuff...
but in this case, it looks good to me...(?)

Files:
- `sol_len.go`
- `sol_len_L.go` is the same as `sol_len.go`, but with larger number of
workers and larger workload, and less processing time per job.
- `sol_len_L_nowait.go` is the same as `sol_len_L.go`, but completely
eliminates the processing time of each job, to make things a bit
tighter..

#### Thoughts

- If there were goroutines leaking, wouldn't `WaitGroup.Wait()`
never return?
- As the number of jobs remaining to be processed gets lesser than
the number of consumer-goroutines, the latter will get `0` returned
by `len()`, and exit.
From this point, the number of consumer-goroutines being up and
processing jobs, should decrease as the number of the remaining
jobs decreases. (?)

#### Verifying the problem with `len()`...

Let `g1` and `g2` be the two last consumer-goroutines, and `v0` be the
last job remaining to be processed, and both `g1` and `g2` have just
finished processing jobs `v1` and `v2`, so they are re-entering the
loop now.

Possible problematic interleaving causing leakage:
- t0: `g1` sees `len(q) == 1`, so it proceeds;
- t1: `g2` sees `len(q) == 1`, so it proceeds;
- t2: `g1` dequeues `v0` (the last remaining job) and begins processing it.

In this case, `g2` will just remain blocked forever on a receive from
`q`.

This is hard to reproduce, and (according to my observations) it occurs
fewer times as the size of the workload and the number of
consumer-goroutines get smaller.

To test it, run one of:
```text
$ for i in $(seq 1 1000); do echo $i ; go run sol_len_L_nowait.go >/dev/null ; done
```

```text
$ for i in $(seq 1 1000); do echo $i ; go run sol_len_XL_nowait.go >/dev/null ; done
```
You should notice that at some point, one of the runs will just hang
there forever, unless interrupted by force (i.e. `SIGINT` for me).

### 2. Using `len()` + `select` to receive

To mitigate the above problem, maybe we could use a `select` after
getting `len()`'s result, to only receive the next job if it is
immediately available, or otherwise go back to the beginning of the
loop, where `len()` will be checked again.

Files:
- `sol_len_sel_XL_nowait.go`, which uses the same workload size and
number of consumer-goroutines as `sol_len_XL_nowait.go`, but also
includes the `select` fix.

This seems to work pretty well... or, at least, I cannot reproduce the
deadlock anymore (even when trying with "XL" workload & #consumers).

To test it, run:
```text
$ for i in $(seq 1 1000); do echo $i ; go run sol_len_sel_XL_nowait.go | grep "HIT" ; done
```
which also prints stuff to stdout every time `select`'s default case is
visited (it should be these times that solution #1 would hang).
