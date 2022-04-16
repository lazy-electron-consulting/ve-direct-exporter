package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/start"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] config-file\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGABRT)
	defer stop()
	err := start.Run(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("exiting with errors %v\n", err)
	}

}
