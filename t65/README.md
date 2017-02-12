# t65

---
---

## Run

```
$ ./tparent.sh
```

---
---

## Walk through (in detail)

---

1. `tparent.sh` sets trap for `SIGUSR1`

2. `tparent.sh` ` &`s `tchild.sh`

3. `tparent.sh` `wait`s

4. `tchild.sh` gets its work done (just counting)

5. `tchild.sh` sends `SIGUSR1` to its **current** parent

**WANRING:** If `tparent.sh` is (for some reason) dead by then, somebody's
gonna have a, probably unpleasant, surprise.

6. `tchild.sh` dies

7. `tparent.sh` wakes up and unsets `SIGUSR1` trap

8. `tparent.sh` dies
