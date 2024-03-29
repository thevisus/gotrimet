package trimet

import (
	"errors"
	"time"
)

// Time is a wrapper for time.Time to workaround format issues.
//
// Times returned from TriMet are not in the proper RFC3339 format, as they
// lack the "Z" specifier separating seconds from timezone.
type Time struct {
	*time.Time
}

const trimetTime = `2006-01-02T15:04:05.999-0700`

// UnmarshalJSON parses a TriMet time into a time.Time.
func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"`+trimetTime+`"`, string(data))
	if nil == err {
		t.Time = new(time.Time)
		*t.Time = parsed
	}
	return err
}

// MarshalJSON formats a Time as a JSON string in RFC3339 format.
func (t *Time) MarshalJSON() ([]byte, error) {
	if nil != t && nil != t.Time {
		if y := t.Year(); y < 0 || y >= 10000 {
			return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
		}
		return []byte(t.Time.Format(`"` + time.RFC3339 + `"`)), nil
	}
	return nil, nil
}

// NewTime returns a new Time wrapping the given time.
func NewTime(t time.Time) *Time {
	return &Time{
		Time: &t,
	}
}

// ParseTime attempts to parse the timestamp using the TriMet format.
func ParseTime(timestamp string) (*Time, error) {
	parsed, err := time.Parse(trimetTime, string(timestamp))
	if nil != err {
		return nil, err
	}

	return NewTime(parsed), nil
}
