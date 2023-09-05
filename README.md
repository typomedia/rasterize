# `rasterize` - Convert svg graphics to png images

`rasterize` is a command line tool to convert svg graphics to png images.

[![Go Report Card](https://goreportcard.com/badge/github.com/typomedia/rasterize)](https://goreportcard.com/report/github.com/typomedia/rasterize)
[![Go Reference](https://pkg.go.dev/badge/github.com/typomedia/rasterize.svg)](https://pkg.go.dev/github.com/typomedia/rasterize)

## Install

    go install github.com/typomedia/rasterize@latest

## Usage

    rasterize [options] [file]

## Options

    -i, --input  string   input svg file
    -o, --output string   output png file
    -s, --size   int      output size (default 1000 pixels)
    -d, --debug  bool     debug mode

## Example

    rasterize -i example.svg
    rasterize --input example.svg --size 512
    
## Debug mode

    rasterize --input example.svg --output image.png --debug
    width:  100
    height:  80
    width:  1000
    height:  800
    Output: image.png
    Elapsed: 206.2827ms

---
Copyright © 2023 Typomedia Foundation. All rights reserved.