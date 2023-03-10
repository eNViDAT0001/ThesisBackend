package wrap_gorm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type SoftDeleteModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
type HardDeleteModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//type DateOnly time.Time
//
//const ctLayout = "2006-01-02 15:04:05 Z07:00"
//
//// UnmarshalJSON Parses the json string in the custom format
//func (ct *DateOnly) UnmarshalJSON(b []byte) (err error) {
//	s := strings.Trim(string(b), `"`)
//	nt, err := time.Parse(ctLayout, s)
//	*ct = DateOnly(nt)
//	return
//}
//
//// MarshalJSON writes a quoted string in the custom format
//func (ct DateOnly) MarshalJSON() ([]byte, error) {
//	fullDate := strings.Trim(ct.String(), `"`)
//	dateOnly := strings.Split(fullDate, ` `)
//	return json.Marshal(fmt.Sprintf("%s %s", dateOnly[0], dateOnly[1]))
//}
//
//// String returns the time in the custom format
//func (ct *DateOnly) String() string {
//	t := time.Time(*ct)
//	return fmt.Sprintf("%q", t.Format(ctLayout))
//}
//
//func StringToDate(s string) (*time.Time, error) {
//	date := strings.Split(s, "-")
//	year, err := strconv.Atoi(date[0])
//	if err != nil {
//		return nil, err
//	}
//	month, err := strconv.Atoi(date[1])
//	if err != nil {
//		return nil, err
//	}
//	day, err := strconv.Atoi(date[2])
//	if err != nil {
//		return nil, err
//	}
//	dateTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
//	return &dateTime, nil
//}
//

const DateOnlyFormat = "2006-01-02 15:04:05 Z07:00"

type DateOnly time.Time

func NewDateOnly(year int, month time.Month, day int) DateOnly {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return DateOnly(t)
}

func (t *DateOnly) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = DateOnly(v)
	case nil:
		*t = DateOnly{}
	default:
		return fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	}
	return nil
}

// MarshalJSON writes a quoted string in the custom format
func (ct DateOnly) MarshalJSON() ([]byte, error) {
	fullDate := strings.Trim(ct.String(), `"`)
	dateOnly := strings.Split(fullDate, ` `)
	return json.Marshal(fmt.Sprintf("%s %s", dateOnly[0], dateOnly[1]))
}

// String returns the time in the custom format
func (ct *DateOnly) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(DateOnlyFormat))
}

func StringToDate(s string) (*time.Time, error) {
	date := strings.Split(s, "-")
	year, err := strconv.Atoi(date[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(date[1])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(date[2])
	if err != nil {
		return nil, err
	}
	dateTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &dateTime, nil
}

func (t DateOnly) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(DateOnlyFormat)), nil
}

func (t *DateOnly) UnmarshalText(value string) error {
	dd, err := time.Parse(DateOnlyFormat, value)
	if err != nil {
		return err
	}
	*t = DateOnly(dd)
	return nil
}

func (DateOnly) GormDataType() string {
	return "TIME"
}
