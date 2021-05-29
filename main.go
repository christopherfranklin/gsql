package main

import (
	"flag"
	"os"
	"time"

	"github.com/christopherfranklin/gsql/cli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Debug bool

func main() {
	debug := flag.Bool("debug", false, "Turns on debug mode")
	flag.Parse()
	Debug = *debug

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	cli.StartPrompt("some file")
}
