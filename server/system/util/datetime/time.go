package datetime

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

// timestamp
func String(f ...string) string {
	format := "2006-01-02 15:04:05"
	if len(f) > 0 {
		format = f[0]
	}
	return time.Now().Format(format)
}

func Hour(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Hour()
}

func Minute(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Minute()
}

func Second(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Second()
}

// String to timestamp
func Timestamp(args ...string) int64 {
	var timestamp int64 = 0
	l := len(args)
	if l == 0 {
		timestamp = time.Now().Unix()
	}
	if l > 0 {
		if reflect.TypeOf(args[0]).String() == "string" {
			t, err := Strtotime(string(args[0]), "2006-01-02 15:04:05")
			if err != nil {
				log.Println("datetime.Timestamp error:", err)
				return -1
			}
			timestamp = t.Unix()
		}
	}
	return timestamp
}

func Millisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func Microsecond() int64 {
	return time.Now().UnixNano() / 1e3
}

func Nanosecond() int64 {
	return time.Now().UnixNano()
}

// GMT TIME
func Gmtime() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
}

func Localtime() string {
	now := time.Now().Local()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	zone, _ := now.Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
}

// String to time.Time
func Strtotime(s string, args ...string) (time.Time, error) {
	format := "2006-01-02 15:04:05"
	if len(args) > 0 {
		format = strings.Trim(args[0], " ")
	}
	if len(s) != len(format) {
		return time.Now(), errors.New("String to time: parameter format error")
	}
	return time.ParseInLocation(format, s, time.Local)
}

func charToCode(layout string) string {
	characters := []string{
		"y", "06",
		"m", "1",
		"d", "2",
		"Y", "2006",
		"M", "01",
		"D", "02",

		"h", "03",
		"H", "15",
		"i", "4",
		"s", "5",
		"I", "04",
		"S", "05",

		"t", "pm",
		"T", "PM",
	}
	replacer := strings.NewReplacer(characters...)
	return replacer.Replace(layout)
}

//返回时间范围
func DateRange(t string, n int64) (time.Time, time.Time) {
	formatStr := "2006-01-02 15:04:05"
	timeBeginStr := time.Now().Format("2006-01-02") + " 00:00:00"
	timeEndStr := time.Now().Format("2006-01-02") + " 23:59:59"
	var First time.Time
	var Last time.Time

	switch t {
	case "d":
		if r, err := time.ParseInLocation(formatStr, timeBeginStr, time.Local); err == nil {
			First = r.AddDate(0, 0, int(n))
		}
		if r, err := time.ParseInLocation(formatStr, timeEndStr, time.Local); err == nil {
			Last = r
		}
	case "w":
		if r, err := time.ParseInLocation(formatStr, timeBeginStr, time.Local); err == nil {
			First = r.AddDate(0, 0, int(-n*7))
		}
		if r, err := time.ParseInLocation(formatStr, timeEndStr, time.Local); err == nil {
			Last = r
		}
	case "m":
		timeBeginStr = time.Now().Format("2006-01") + "-01 00:00:00"
		timeEndStr = time.Now().Format("2006-01-02") + " 23:59:59"
		if r, err := time.ParseInLocation(formatStr, timeBeginStr, time.Local); err == nil {
			First = r.AddDate(0, int(-n), 0)
		}
		if r, err := time.ParseInLocation(formatStr, timeEndStr, time.Local); err == nil {
			Last = r
		}
	case "y":
		timeBeginStr = time.Now().Format("2006") + "-01-01 00:00:00"
		timeEndStr = time.Now().Format("2006-01-02") + " 23:59:59"
		if r, err := time.ParseInLocation(formatStr, timeBeginStr, time.Local); err == nil {
			First = r.AddDate(int(-n), 0, 0)
		}
		if r, err := time.ParseInLocation(formatStr, timeEndStr, time.Local); err == nil {
			Last = r
		}

	}

	return First, Last

}
