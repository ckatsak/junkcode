# t66

---
---

## Run

---

```
$ ./tparent.sh
```

---
---

## Walk through (in detail)

---

1. `tparent.sh` sets trap for `SIGUSR1`

2. `tparent.sh` ` &`s `tinf.sh` and `wait`s to **it**

3. `tinf.sh` gets some work done

4. `tinf.sh` sends a `SIGUSR1` to its **current** parent

**WARNING:** `tinf.sh` *doesn't care who the parent is, see also t65/*

5. `tinf.sh` goes to sleep for some time

6. `tparent.sh` handles the `SIGUSR1` and dies

7. `tinf.sh` dies
