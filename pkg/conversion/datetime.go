package conversion

import (
	"strings"
	"time"
)

func ConvertMysqlTimeUnixTime(mysqlTime string) int64 {
	res1 := strings.Replace(mysqlTime, "T", " ", 1)
	res2 := res1[:19]

	// YYYY-MM-DD
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, res2)
	if err != nil {
		panic(err)
	}

	return t.Unix()
}

func ConvertUnixTimeMySqlTime(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func CompareTime(start int64, end int64, check int64) bool {
	return false
}
