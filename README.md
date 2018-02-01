# scraperexporter

Simple webserver that scrapes sites on configuration, filter the retrieved data and export them on Prometheus format

[![Build Status](https://travis-ci.org/marcelosousaalmeida/scraperexporter.svg?branch=master)](https://travis-ci.org/marcelosousaalmeida/scraperexporter) [![Go Report Card](https://goreportcard.com/badge/github.com/marcelosousaalmeida/scraperexporter)](https://goreportcard.com/report/github.com/marcelosousaalmeida/scraperexporter)

## Installation

If you are using Go 1.6+ (or 1.5 with the `GO15VENDOREXPERIMENT=1` environment variable), you can install `scraperexporter` with the following command:

```bash
$ go get -u github.com/marcelosousaalmeida/scraperexporter
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
