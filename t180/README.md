# t180


## Problem Description

```text
$ go run prob.go
```

In `prob.go`, the main-goroutine produces some workload (`int`s,
representing a "job") and fills a buffered channel (to be used as a
blocking queue for multiple consumer-goroutines).

The channel is not closed by main-goroutine, because processing of a job
(by a consumer-goroutine) may fail, in which case it has to be re-enqueued
(re-sent to the end of the buffered channel) to allow for attempting to
process it again later (either by the same or by some other
consumer-goroutine).
A job may fail and need to be re-enqueued many times (however, in this
example it is re-enqueued exactly once per consumer-goroutine that grabs
it).

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
senders and `N` receivers (where `N` is the number of consumer-goroutines).


## tl;dr GOAL

Modify `prob.go` so that when all the workload is properly processed by
the consumer-goroutines, the main-goroutine prints `THE END` to stdout.


## Solution?

### using `len()` on the chan

Yeah.. I know, normally `len()` is weird on channels, and stuff...
but in this case, it looks good to me...(?)

- `sol.go`
- `sol-large.go` is the same as `sol.go`, but with larger number of
workers and larger workload, and less processing time per job.
- `sol-nowait.go` is the same as `sol-large.go`, but completely
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
