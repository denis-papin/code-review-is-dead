package marketdata

import (
	"regexp"
	"time"
)

var symbolPattern = regexp.MustCompile("^[A-Za-z0-9^][A-Za-z0-9.^=-]{0,19}$")

func utc(seconds int64) time.Time { return time.Unix(seconds, 0).UTC() }
