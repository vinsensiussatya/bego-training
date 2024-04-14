package http

import (
	"time"

	"github.com/vinsensiussatya/bego-training/internal/pkg/util"
)

type Date time.Time

func (d *Date) Time() time.Time {
	return time.Time(*d)
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (d *Date) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*d = Date(time.Time{})
		return nil
	}
	t, err := time.Parse(DateLayout, string(data))
	*d = Date(t)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (d *Date) MarshalText() ([]byte, error) {
	b := make([]byte, 0, len(DateLayout))
	return d.Time().AppendFormat(b, DateLayout), nil
}

func (d *Date) String() string {
	return d.Time().Format(DateLayout)
}

func (d *Date) StartOfDate() time.Time {
	if d.Time().IsZero() {
		return time.Time{}
	}
	return util.GetStartOfTheDay(d.Time())
}
func (d *Date) EndOfDate() time.Time {
	if d.Time().IsZero() {
		return time.Time{}
	}
	return util.GetEndOfTheDay(d.Time())
}
