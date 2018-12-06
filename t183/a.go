package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

const (
	serverPort = 54242
	bufSize    = 1500
)

var (
	logFile  *os.File
	stdoutMu = sync.Mutex{}
)

func init() {
	var err error
	logFile, err = os.OpenFile("a.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[main] Error opening log file: %s\n", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	defer logFile.Close()
	go NewEchoServer().Run()
	go NewEchoClient(0).Run()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, unix.SIGTERM)
	<-s
}

type EchoServer struct {
	conn *net.UDPConn
}

func NewEchoServer() *EchoServer {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   []byte{127, 0, 0, 2},
		Port: serverPort},
	)
	if err != nil {
		log.Fatalf("[main] NewEchoServer: %s\n", err)
	}
	log.Printf("[server] listening to 127.0.0.2:%d", serverPort)

	return &EchoServer{conn: conn}
}

func (es *EchoServer) Run() {
	defer func() {
		if err := es.conn.Close(); err != nil {
			log.Panicf("[server] deferred conn.Close(): %s\n", err)
		}
	}()
	for {
		buf := make([]byte, bufSize)
		n, addr, err := es.conn.ReadFrom(buf)
		go es.handleDatagram(buf, n, addr, err)
	}
}

func (es *EchoServer) handleDatagram(b []byte, n int, addr net.Addr, err error) {
	out := bytes.NewBufferString(fmt.Sprintf("[server] received %d byte(s) from '%v':\n\t", n, addr))
	if m, e := out.Write(b[:n]); m != n || e != nil {
		log.Panicf("[server] error writing to out buffer: (written %d out of %d byte(s)) %s\n", m, n, e)
	}
	if e := out.WriteByte('\n'); e != nil {
		log.Panicf("[server] error writing '\\n\\n' to out buffer:", e)
	}
	if err != nil {
		if _, e := out.WriteString(fmt.Sprintf("\tand error: %v\n", err)); e != nil {
			log.Panicf("[server] error writing received error to out buffer: %s\n", e)
		}
	}
	log.Printf(out.String())

	stdoutMu.Lock()
	defer stdoutMu.Unlock()
	fmt.Printf("[server] < %q\n", b[:n])
}

/*
func (es *EchoServer) handleDatagram(b []byte, n int, addr net.Addr, err error) {
	out := bytes.NewBufferString(fmt.Sprintf("\n[server] Received %d bytes from '%v':\n\t", n, addr))
	if m, e := out.Write(b[:n]); m != n || e != nil {
		log.Printf("[server] error writing to out buffer: (written %d out of %d bytes) %s\n", m, n, e)
	}
	if m, e := out.Write([]byte{'\n', '\n'}); m != 2 || e != nil {
		//if e := out.WriteByte('\n'); e != nil {
		log.Printf("[server] error writing '\\n\\n' to out buffer:", e)
	}
	if err != nil {
		if _, e := out.WriteString(fmt.Sprintf("\nand error: %v\n\n", err)); e != nil {
			log.Printf("[server] error writing received error to out buffer: %s\n", e)
		}
	}

	stdoutMu.Lock()
	defer stdoutMu.Unlock()
	if _, err = out.WriteTo(os.Stdout); err != nil {
		log.Printf("[server] error writing to stdout: %s\n", err)
	}
}
*/

type EchoClient struct {
	id int
}

func NewEchoClient(id int) *EchoClient { return &EchoClient{id: id} }

func (ec *EchoClient) Run() {
	conn, err := net.Dial("udp", fmt.Sprintf("127.0.0.2:%d", serverPort))
	if err != nil {
		log.Fatalf("[client-%d] error connecting to the server: %s\n", ec.id, err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Panicf("[client-%d] deferred conn.Close(): %s\n", ec.id, err)
		}
	}()
	log.Printf("[client-%d] successfully dialed 127.0.0.2:%d\n", ec.id, serverPort)

	in := bufio.NewReader(os.Stdin)
	for {
		<-time.After(50 * time.Millisecond)
		stdoutMu.Lock()
		fmt.Printf("\r[client-%d] > ", ec.id)
		stdoutMu.Unlock()
		buf, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("[client-%d] spawning a new client and exiting...\n", ec.id)
				break
			} else {
				log.Printf("[client-%d] error reading from stdin: %s\n", ec.id, err)
				continue
			}
		}

		n, err := conn.Write([]byte(strings.TrimSpace(buf)))
		if err != nil {
			log.Printf("[client-%d] error writing to UDP: %s\n", ec.id, err)
		}
		log.Printf("[client-%d] wrote %d byte(s) to UDP\n", ec.id, n)
	}

	go NewEchoClient(ec.id + 1).Run()
}
