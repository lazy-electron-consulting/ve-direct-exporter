package config_test

import (
	"embed"
	"testing"
	"time"

	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/config"
	"github.com/stretchr/testify/require"
)

//go:embed testdata
var testdata embed.FS

func TestParseYaml(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path string
		want *config.Config
	}{
		"full": {
			path: "testdata/full.yml",
			want: &config.Config{
				Address: ":9090",
				Serial: config.Serial{
					Path:     "/dev/ttyUSB0",
					BaudRate: 9600,
					DataBits: 7,
					StopBits: 2,
					Parity:   "Y",
					Timeout:  time.Hour,
				},
				Gauges: []config.Gauge{
					{
						Name:       "battery_volts",
						Help:       "Main battery voltage",
						Label:      "V",
						Multiplier: 0.01,
					},
				},
			},
		},
		"empty": {
			path: "testdata/empty.yml",
			want: &config.Config{
				Address: config.DefaultAddress,
				Serial: config.Serial{
					Path:     "/dev/ttyUSB1",
					BaudRate: config.DefaultBaudRate,
					DataBits: config.DefaultDataBits,
					StopBits: config.DefaultStopBits,
					Parity:   config.DefaultParity,
					Timeout:  config.DefaultTimeout,
				},
			},
		},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			file, err := testdata.Open(tc.path)
			require.NoError(t, err)

			got, err := config.ParseYaml(file)
			require.NoError(t, err)
			require.EqualValues(t, tc.want, got)
		})
	}
}
