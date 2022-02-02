[![Go Report Card](https://goreportcard.com/badge/github.com/cnvergence/invoice-generator)](https://goreportcard.com/report/github.com/cnvergence/invoice-generator)
[![GoDoc](https://godoc.org/github.com/cnvergence/invoice-generator?status.svg)](https://godoc.org/github.com/cnvergence/invoice-generator)

# Invoice generator

Simple invoice generator app, written in Go. Generating PDF invoice from YAML, using https://github.com/johnfercher/maroto library.

# Generate

Install the package:

`go get github.com/cnvergence/invoice-generator`

Generate the invoice with this command:

`invoice-generator generate --yaml=<path to yaml values> --out=<path to save a PDF file>

For the example, please check the example folder.