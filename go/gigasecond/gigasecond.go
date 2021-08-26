package gigasecond

// import path for the time package from the standard library
import "time"

// Return the time after a gigasecond has passed
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
