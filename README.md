# ve-direct-exporter

[Prometheus exporter](https://prometheus.io/docs/instrumenting/exporters/)  for [Victron](https://www.victronenergy.com) devices using the [VE.Direct protocol](https://www.victronenergy.com/live/vedirect_protocol:faq).

## Usage

- download a binary from the github releases or build from source using `make
  build`
- connect your computer to a victron device, I'm using a victron brand [VE.Direct to USB interface](https://www.victronenergy.com/accessories/ve-direct-to-usb-interface)
- run `./vedirect-exporter config.yaml`
- see your metrics via http

See `configs/` for an example of the config syntax.

## Releasing

- merging to `main` will make new "latest" binaries
- push a tag w/ `v$SEMVER` to make a versioned binaries
