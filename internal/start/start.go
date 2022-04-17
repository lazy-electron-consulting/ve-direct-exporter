package start

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/goburrow/serial"
	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/config"
	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/metrics"
	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/scanner"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(ctx context.Context, cfg *config.Config) error {
	port, err := serial.Open(&serial.Config{
		Address:  cfg.Serial.Path,
		BaudRate: cfg.Serial.BaudRate,
		DataBits: cfg.Serial.DataBits,
		Parity:   cfg.Serial.Parity,
		StopBits: cfg.Serial.StopBits,
	})
	if err != nil {
		return fmt.Errorf("could not open serial port: %w", err)
	}
	defer port.Close()

	s := scanner.New(port)
	r, err := metrics.New(cfg.Subsystem, cfg.Gauges)
	if err != nil {
		return fmt.Errorf("cannot create metrics registry: %w", err)
	}
	rc := make(chan scanner.Reading)
	go r.Run(ctx, rc)
	go s.Run(ctx, rc)

	return runHttp(ctx, cfg.Address)
}

func runHttp(ctx context.Context, addr string) error {
	http.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{Addr: addr}
	go func() {
		<-ctx.Done()
		log.Println("stopping")
		srv.Close()
	}()

	defer log.Println("http server stopped")
	log.Printf("http server started on %v\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("error serving %v\n", err)
		return err
	}
	return nil
}
