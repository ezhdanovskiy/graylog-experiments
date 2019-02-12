package main

import (
	"flag"
	"gopkg.in/Graylog2/go-gelf.v1/gelf"
	"io"
	"log"
	"os"
)

func main() {
	var graylogAddr string

	flag.StringVar(&graylogAddr, "graylog", "", "graylog server addr")
	flag.Parse()

	graylogAddr = "localhost:12201"
	if graylogAddr != "" {
		gelfWriter, err := gelf.NewWriter(graylogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// log to both stderr and graylog2
		log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
		log.Printf("logging to stderr & graylog2@'%s'", graylogAddr)
	}

	// From here on out, any calls to log.Print* functions
	// will appear on stdout, and be sent over UDP to the
	// specified Graylog2 server.

	log.Printf("Hello gray World")
	log.Printf(`{"id":123, "version":1}`)

	// ...
}
