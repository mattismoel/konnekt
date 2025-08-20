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

func (ut1 UnixTime) Equals(ut2 UnixTime) bool {
	return ut1.Time().Equal(ut2.Time())
}

func (uts1 UnixTimestamps) Equals(uts2 UnixTimestamps) bool {
	return uts1.CreatedAt.Equals(uts2.CreatedAt) &&
		uts1.UpdatedAt.Equals(uts2.UpdatedAt)
}
