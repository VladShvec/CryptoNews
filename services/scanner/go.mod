module cryptonews/scanner

go 1.24.0

replace cryptonews/shared => ../../shared

require (
	cryptonews/shared v0.0.0
	github.com/k0kubun/pp/v3 v3.5.1
	github.com/segmentio/kafka-go v0.4.50
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.32.0 // indirect
)
