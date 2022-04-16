package config

import (
	"fmt"
	"io"

	"github.com/lazy-electron-consulting/ve-direct-exporter/internal/util"
	"gopkg.in/yaml.v2"
)

const (
	DefaultAddress  = ":8000"
	DefaultBaudRate = 19200
	DefaultDataBits = 8
	DefaultStopBits = 1
	DefaultParity   = "N"
)

type Serial struct {
	Path     string `json:"path,omitempty" yaml:"path,omitempty"`
	BaudRate int    `json:"baudRate,omitempty" yaml:"baudRate,omitempty"`
	DataBits int    `json:"dataBits,omitempty" yaml:"dataBits,omitempty"`
	StopBits int    `json:"stopBits,omitempty" yaml:"stopBits,omitempty"`
	Parity   string `json:"parity,omitempty" yaml:"parity,omitempty"`
}

func (m *Serial) defaults() {
	m.BaudRate = util.Default(m.BaudRate, DefaultBaudRate)
	m.DataBits = util.Default(m.DataBits, DefaultDataBits)
	m.StopBits = util.Default(m.StopBits, DefaultStopBits)
	m.Parity = util.Default(m.Parity, DefaultParity)
}

type Gauge struct {
	Name       string  `json:"name,omitempty" yaml:"name,omitempty"`
	Label      string  `json:"label,omitempty" yaml:"label,omitempty"`
	Help       string  `json:"help,omitempty" yaml:"help,omitempty"`
	Multiplier float32 `json:"multiplier,omitempty" yaml:"multiplier,omitempty"`
}

type Config struct {
	Address string  `json:"address,omitempty" yaml:"address,omitempty"`
	Serial  Serial  `json:"serial,omitempty" yaml:"serial,omitempty"`
	Gauges  []Gauge `json:"gauges,omitempty" yaml:"gauges,omitempty"`
}

func (c *Config) defaults() {
	c.Address = util.Default(c.Address, DefaultAddress)
	c.Serial.defaults()

}

func ParseYaml(r io.Reader) (*Config, error) {
	decoder := yaml.NewDecoder(r)
	decoder.SetStrict(true)

	var config Config
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}
	config.defaults()
	return &config, nil
}
