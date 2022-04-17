package metrics

import (
	"context"
	"log"
	"strconv"

	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/config"
	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/scanner"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type gauge struct {
	prometheus.Gauge
	cfg config.Gauge
}

func newGauge(subsystem string, cfg config.Gauge) gauge {
	return gauge{
		Gauge: promauto.NewGauge(prometheus.GaugeOpts{
			Subsystem: subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
		}),
		cfg: cfg,
	}
}

func (g gauge) update(value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	v := float64(i) * g.cfg.Multiplier
	g.Set(v)
	return nil
}

type Registry struct {
	gauges map[string]gauge
}

func New(subsystem string, gs []config.Gauge) (*Registry, error) {
	r := Registry{
		gauges: make(map[string]gauge, len(gs)),
	}

	for _, g := range gs {
		r.gauges[g.Label] = newGauge(subsystem, g)
	}

	return &r, nil
}

func (r *Registry) Run(ctx context.Context, in <-chan scanner.Reading) {
	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-in:
			if !ok {
				return
			}
			err := r.update(d.Label, d.Value)
			if err != nil {
				log.Printf("failed to update %v, %v\n", d, err)
			}
		}
	}
}

func (r *Registry) update(label, value string) error {
	g, ok := r.gauges[label]
	if ok {
		return g.update(value)
	}
	return nil
}
