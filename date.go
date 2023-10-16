package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

var (
	dateFormat = "2006-01-02"
)

type Date time.Time

// NewDate - create Date by time.Time
func NewDate(tm time.Time) Date {
	return Date(tm)
}

// DateToday - returns Today Date
func DateToday() Date {
	return NewDate(time.Now())
}

func (date *Date) Scan(value interface{}) (err error) {
	tm := &sql.NullTime{}
	err = tm.Scan(value)
	*date = Date(tm.Time)
	return
}

func (date Date) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(date).Location()).Format(dateFormat), nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	ts := time.Time(date).Format(dateFormat)

	return []byte(fmt.Sprintf("\"%s\"", ts)), nil
}

func (date *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	s, _ = strings.CutPrefix(s, "\"")
	s, _ = strings.CutSuffix(s, "\"")

	d, err := time.ParseInLocation(dateFormat, s, time.Local)
	if err != nil {
		return fmt.Errorf("1: cannot parse time %s", b)
	}

	*date = Date(d)

	return nil
}
