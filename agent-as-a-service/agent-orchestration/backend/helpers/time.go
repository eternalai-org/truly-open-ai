package helpers

import (
	"time"
)

const (
	DEFAULT_DATE_FORMAT      = "2006-01-02"
	DEFAULT_TIME_FORMAT      = "2006-01-02 15:04:05"
	TWITTER_DATE_TIME_FORMAT = "2006-01-02T15:04:05.000Z"

	DEFAULT_TIME_ONLY_FORMAT   = "15:04:05"
	DEFAULT_FILTER_TIME_FORMAT = "20060102150405"
)

func TimeFromUnix(n int64) *time.Time {
	t := time.Unix(n, 0)
	if t.Year() > 9999 {
		t = time.Unix(n/1000, 0)
	}
	return &t
}

func ParseStringToDateDefault(value string) *time.Time {
	return ParseStringToTime(DEFAULT_DATE_FORMAT, value)
}

func ParseStringToDateTimeDefault(value string) *time.Time {
	return ParseStringToTime(DEFAULT_TIME_FORMAT, value)
}

func ParseStringToDateTimeTwitter(value string) *time.Time {
	return ParseStringToTime(TWITTER_DATE_TIME_FORMAT, value)
}

func ParseStringToTime(layout string, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil
	}
	return &t
}

func ParseTimeToStringDateDefault(value *time.Time) string {
	return ParseTimeToString(DEFAULT_DATE_FORMAT, value)
}

func ParseTimeToStringDateTimeDefault(value *time.Time) string {
	return ParseTimeToString(DEFAULT_TIME_FORMAT, value)
}

func ParseTimeToString(layout string, value *time.Time) string {
	if value == nil {
		return ""
	}
	return value.Format(layout)
}

func NewDate(day int, month time.Month, year int) *time.Time {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	return &t
}

func NewDateForDayOfWeek(dayOffWeek time.Weekday, month time.Month, year int, number int) *time.Time {
	if number == 0 {
		return nil
	}
	var result *time.Time
	var num int
	for i := 1; i <= 31; i++ {
		t := time.Date(year, month, i, 0, 0, 0, 0, time.Now().Location())
		if t.Month() != month {
			break
		}
		if t.Weekday() == dayOffWeek {
			num++
		}
		if num == number {
			result = &t
			break
		}
	}
	return result
}

func NewLastDateForDayOfWeek(dayOffWeek time.Weekday, month time.Month, year int) *time.Time {
	var result *time.Time
	var num int
	for i := 1; i <= 31; i++ {
		t := time.Date(year, month, i, 0, 0, 0, 0, time.Now().Location())
		if t.Month() != month {
			break
		}
		if t.Weekday() == dayOffWeek {
			num++
		}
		if num == num {
			result = &t
		}
	}
	return result
}

func TimeNow() *time.Time {
	t := time.Now()
	return &t
}

func TimeNowAdd(d time.Duration) *time.Time {
	t := time.Now().Add(d)
	return &t
}

func TimeAdd(t time.Time, d time.Duration) *time.Time {
	ts := t.Add(d)
	return &ts
}

func TruncateDate(toRound time.Time) time.Time {
	rounded := time.Date(toRound.Year(), toRound.Month(), toRound.Day(), 0, 0, 0, 0, toRound.Location())
	return rounded
}

func NewNearbyDayOfWeekAt(t *time.Time, dayOffWeek time.Weekday) *time.Time {
	for i := 1; i <= 7; i++ {
		t1 := t.AddDate(0, 0, -1)
		t = &t1
		if t.Weekday() == dayOffWeek {
			break
		}
	}
	return t
}

func NewNearbyDayOfMonthAt(t *time.Time, dayOffMonth int) *time.Time {
	ts := *t
	for i := 1; i <= 31; i++ {
		ts = ts.Add(-24 * time.Hour)
		if ts.Month() == t.Month() {
			if ts.Day() == dayOffMonth {
				break
			}
		} else {
			if ts.Day() == dayOffMonth {
				break
			}
			if ts.Day() < dayOffMonth {
				break
			}
		}
	}
	return &ts
}

func ToWorkHours(t1 *time.Time, t2 *time.Time) float64 {
	var workHours float64
	if t1.Unix() >= t2.Unix() {
		return 0
	}
	for t1.Unix() <= t2.Unix() {
		if t1.Weekday() != time.Sunday &&
			t1.Weekday() != time.Saturday {
			workHours += 8
		}
		t := t1.AddDate(0, 0, 1)
		t1 = &t
	}
	return workHours
}

func GetStartDayOfWeek(t time.Time) time.Time { //get monday 00:00:00
	t = t.UTC()
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := t.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour)
}

func GetStartDayOfMonth(t time.Time) time.Time { //get monday 00:00:00
	t = t.UTC()
	year, month, _ := t.Date()
	r := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return r
}

func ParseStringToFilterDateDefault(value string) *time.Time {
	return ParseStringToTime(DEFAULT_FILTER_TIME_FORMAT, value)
}

func GetNextScheduleTime(interval time.Duration, delay time.Duration) *time.Time {
	t := time.Now().Truncate(interval).Add(delay)
	if t.Before(time.Now()) {
		t = t.Add(interval)
	}
	return &t
}
