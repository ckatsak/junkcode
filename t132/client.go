// A simple HTTP (forever-)streaming client of JSON data for testing with
// server.py.
//
// Expects chunked responses with JSON lists of strings.
//
// Usage example:
//	$ go run client.go stream
//	$ go run client.go range 0 103 7
//	$ go run client.go rand
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	// Subcommands; no flags at all for now
	streamCmd = flag.NewFlagSet("stream", flag.ExitOnError)
	rangeCmd  = flag.NewFlagSet("range", flag.ExitOnError)
	randCmd   = flag.NewFlagSet("rand", flag.ExitOnError)
)

func getStream() {
	results := []string{}

	resp, err := http.Get("http://localhost:58080/stream")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		dec := json.NewDecoder(resp.Body)
		for {
			err := dec.Decode(&results)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			fmt.Println("Received chunk:", results)
		}
	} else {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Printf("Received: %q\n", resp.Status)
			if err != io.EOF {
				panic(err)
			}
		} else {
			fmt.Printf("Received: %s\n\t%q\n", resp.Status, string(body))
		}
	}
}

func getRange(start, end, step int) {
	results := []string{}

	resp, err := http.Get(fmt.Sprintf("http://localhost:58080/range/%d/%d/%d", start, end, step))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		dec := json.NewDecoder(resp.Body)
		for {
			err := dec.Decode(&results)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			fmt.Println("Received chunk:", results)
		}
	} else {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Printf("Received: %q\n", resp.Status)
			if err != io.EOF {
				panic(err)
			}
		} else {
			fmt.Printf("Received: %s\n\t%q\n", resp.Status, string(body))
		}
	}
}

func getRand() {
	results := []string{}

	resp, err := http.Get("http://localhost:58080/rand")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		dec := json.NewDecoder(resp.Body)
		for {
			err := dec.Decode(&results)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			fmt.Println("Received chunk:", results)
		}
	} else {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Printf("Received: %q\n", resp.Status)
			if err != io.EOF {
				panic(err)
			}
		} else {
			fmt.Printf("Received: %s\n\t%q\n", resp.Status, string(body))
		}
	}
}

func main() {
	defer fmt.Println("Exiting...")
	flag.Parse()
	switch flag.Arg(0) {
	case "stream":
		streamCmd.Parse(flag.Args())
		getStream()
	case "range":
		rangeCmd.Parse(flag.Args())
		if rangeCmd.NArg() != 4 {
			fmt.Fprintf(os.Stderr, "Usage:\n\t$ %s range <start> <end> <step>\n", os.Args[0])
			os.Exit(1)
		}
		start, err := strconv.Atoi(rangeCmd.Arg(1))
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(rangeCmd.Arg(2))
		if err != nil {
			panic(err)
		}
		step, err := strconv.Atoi(rangeCmd.Arg(3))
		if err != nil {
			panic(err)
		}
		getRange(start, end, step)
	case "rand":
		randCmd.Parse(flag.Args())
		getRand()
	default:
		fmt.Fprintln(os.Stderr, "Unkown command.")
	}
}
