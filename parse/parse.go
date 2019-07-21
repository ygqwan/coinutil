package parse

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func ParseMsToUtime(ms interface{}) (time.Time, error) {
	if ms == nil {
		return time.Now(), errors.New("ms is empty")
	}
	var m int64 = 0
	var err error
	switch ms.(type) {
	case string:
		m, err = strconv.ParseInt(ms.(string), 10, 64)
		if err != nil {
			return time.Now(), err
		}
	case float64:
		m = int64(ms.(float64))
	case int64:
		m = ms.(int64)
	case int:
		m = int64(ms.(int))
	case int32:
		m = int64(ms.(int32))
	default:
		return time.Now(), fmt.Errorf("not support ms, ms:%+v", ms)
	}
	     //1563689757037992000
	if m > 999999999999999999 {
		return time.Unix(m / 1000000000, m % 1000000000), nil
	} else if m > 999999999999999 {
		return time.Unix(m / 1000000, m % 1000000 * 1000), nil
	} else if m > 999999999999 { //毫秒
		return time.Unix(m / 1000, m % 1000 * 1000000), nil
	} else {
		return time.Unix(m, 0), nil
	}
}