package constant

import "time"

const (
	ValueEmpty         = 0
	StringEmpty string = ""
)

const (
	PageSizeDefault  = 20
	PageIndexDefault = 1
)

const (
	TimeoutRequestDefault = 5 * time.Second
	TimeoutServerDefault  = 5 * time.Second
)

const (
	TimeFormatDefault string = "02-01-2006 15:04:05"
)

const (
	NanoIdAlphabet string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NanoIdSize     int    = 7
)
