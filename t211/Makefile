CFLAGS = -Wall -Wextra -O2
BIN := thp

$(BIN): thp.c
	$(CC) $(CFLAGS) $< -o $@
	strip -s $(BIN)

clean:
	$(RM) $(BIN)

