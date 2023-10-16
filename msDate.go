package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type MsDate time.Time

func (date *MsDate) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = MsDate(nullTime.Time)
	return
}

func (date MsDate) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(date).Location()), nil
}

func (date MsDate) MarshalJSON() ([]byte, error) {
	ts := time.Time(date).Format("2006-01-02 15:04")

	return []byte(fmt.Sprintf("\"%s\"", ts)), nil
}

func (date *MsDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s, _ = strings.CutPrefix(s, "\"")
	s, _ = strings.CutSuffix(s, "\"")

	d, err := time.ParseInLocation("2006-01-02T15:04:05.999Z", s, time.Local)
	if err != nil {
		return fmt.Errorf("1: cannot parse time %s", b)
	}

	*date = MsDate(d)

	return nil
}
