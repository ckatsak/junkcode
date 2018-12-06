package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

const (
	serverPort = 54242
	bufSize    = 1500
)

var (
	logFile *os.File
)

func init() {
	var err error
	logFile, err = os.OpenFile("b.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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
	out := bytes.NewBufferString(fmt.Sprintf("[server] received %d byte(s) from '%v':\t", n, addr))
	if m, e := fmt.Fprintf(out, "%02x", b[:n]); e != nil {
		log.Panicf("[server] error writing to out buffer: (written %d out of %d byte(s)) %s\n", m, n, e)
	}
	//if m, e := out.Write(b[:n]); m != n || e != nil {
	//	log.Panicf("[server] error writing to out buffer: (written %d out of %d byte(s)) %s\n", m, n, e)
	//}
	if e := out.WriteByte('\n'); e != nil {
		log.Panicf("[server] error writing '\\n\\n' to out buffer:", e)
	}
	if err != nil {
		if _, e := out.WriteString(fmt.Sprintf("\n\tand error: %v\n", err)); e != nil {
			log.Panicf("[server] error writing received error to out buffer: %s\n", e)
		}
	}
	log.Printf(out.String())
}

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

	for i := uint16(0); i < uint16(0xffff); i++ {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, i)
		n, err := conn.Write(buf)
		if err != nil {
			log.Printf("[client-%d] error writing to UDP: %s\n", ec.id, err)
		}
		log.Printf("[client-%d] wrote %d byte(s) to UDP\n", ec.id, n)
	}
}
