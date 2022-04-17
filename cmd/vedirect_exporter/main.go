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

	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/config"
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

	cfg, err := config.ReadYaml(flag.Arg(0))
	if err != nil {
		log.Fatalf("could not read config %s: %v", flag.Arg(0), err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGABRT)
	defer stop()
	err = start.Run(ctx, cfg)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("exiting with errors %v\n", err)
	}

}
