package sqlite

import "time"

type UnixTime int64

type UnixTimestamps struct {
	CreatedAt UnixTime
	UpdatedAt UnixTime
}

func (uts UnixTimestamps) Times() (time.Time, time.Time) {
	return uts.CreatedAt.Time(), uts.UpdatedAt.Time()
}

func (ut UnixTime) Time() time.Time {
	return time.Unix(int64(ut), 0)
}
