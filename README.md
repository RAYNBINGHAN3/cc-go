# CC-GO
English | [简体中文](./README.zh.md)

Script to perform a CC DDOS by Go (for test only)

## Usage

From terminal:
`go run main.go http://127.0.0.1/ccgo  -w 500 -t 20`

## Parameters

<pre>
Flags:
  -h, --help            help for cc-go
  -t, --time string     How long(s) (default "10")
  -w, --worker string   The number of worker threads executing concurrently (default "100")
</pre>

## Requirements
- Go 1.12 version or greater

## Disclaimer

This tool is written for educational purpose only, **please** use it on your own good faith.
