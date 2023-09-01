package util

import "time"

func NowUnixTimeMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
