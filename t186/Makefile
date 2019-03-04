BIN= t186
CFLAGS= -Wall -Wextra $(shell pkg-config --cflags openssl)
LDFLAGS= $(shell pkg-config --libs openssl)

$(BIN):
	$(CC) $(CFLAGS) t186.c -o $(BIN) $(LDFLAGS)

clean:
	$(RM) $(BIN)
