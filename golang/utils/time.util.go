package utils

import "time"

func TimeDurationToInt(duration time.Duration) int {
	return int(duration.Seconds())
}
