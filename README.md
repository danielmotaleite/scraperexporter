# scraperexporter

Simple webserver that scrapes sites on configuration, filter the retrieved data and export them on Prometheus format

[![Build Status](https://travis-ci.org/marceloalmeida/scraperexporter.svg?branch=master)](https://travis-ci.org/marceloalmeida/scraperexporter) [![Go Report Card](https://goreportcard.com/badge/github.com/marceloalmeida/scraperexporter)](https://goreportcard.com/report/github.com/marceloalmeida/scraperexporter) [![Maintainability](https://api.codeclimate.com/v1/badges/2621f3d48115aa78b57c/maintainability)](https://codeclimate.com/github/marceloalmeida/scraperexporter/maintainability)

## Installation

If you are using Go 1.6+ (or 1.5 with the `GO15VENDOREXPERIMENT=1` environment variable), you can install `scraperexporter` with the following command:

```bash
$ go get -u github.com/marceloalmeida/scraperexporter
```

## Usage

```bash
$ ./scraperexporter --help
Usage of ./scraperexporter:
  -configuration-file string
    	Configuration file (default "conf.json")
  -generate-config
    	Generate configuration example file

$ ./scraperexporter
```
## Contributing

All contributions are welcome, but if you are considering significant changes, please open an issue beforehand and discuss it with us.

## License

MIT. See the `LICENSE` file for more information.
