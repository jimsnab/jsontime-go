package jsontime

import (
	"errors"
	"fmt"
	"time"
)

type (
	JsonTimeSec struct {
		time.Time
	}

	JsonTimeMs struct {
		time.Time
	}

	JsonTimeUs struct {
		time.Time
	}
)

func parseJsonTime(str string) (t time.Time, err error) {
	switch len(str) {
	case 22:
		return time.Parse("2006-01-02T15:04:05Z", str[1:21])
	case 26:
		return time.Parse("2006-01-02T15:04:05.000Z", str[1:25])
	case 29:
		return time.Parse("2006-01-02T15:04:05.000000Z", str[1:28])
	case 32:
		return time.Parse("2006-01-02T15:04:05.000000000Z", str[1:31])
	}

	err = errors.New("unrecognized timestamp format in " + str)
	return
}

func (sec JsonTimeSec) MarshalJSON() ([]byte, error) {
	str := `""`
	if !sec.IsZero() {
		str = fmt.Sprintf(`"%s"`, sec.Format("2006-01-02T15:04:05Z"))
	}
	return []byte(str), nil
}

func (sec *JsonTimeSec) UnmarshalJSON(text []byte) error {
	if len(text) == 2 {
		return nil
	}
	if len(text) < 22 {
		return errors.New("malformed second timestamp")
	}

	t, err := parseJsonTime(string(text))
	if err != nil {
		return err
	}
	sec.Time = time.Unix(t.Unix(), 0).UTC()
	return nil
}

func (ms JsonTimeMs) MarshalJSON() ([]byte, error) {
	str := `""`
	if !ms.IsZero() {
		str = fmt.Sprintf(`"%s"`, ms.Format("2006-01-02T15:04:05.000Z"))
	}
	return []byte(str), nil
}

func (ms *JsonTimeMs) UnmarshalJSON(text []byte) error {
	if len(text) == 2 {
		return nil
	}
	t, err := parseJsonTime(string(text))
	if err != nil {
		return err
	}
	n := t.UnixMilli()
	ms.Time = time.Unix(n/1000, (n%1000)*1000000).UTC()
	return nil
}

func (us JsonTimeUs) MarshalJSON() ([]byte, error) {
	str := `""`
	if !us.IsZero() {
		str = fmt.Sprintf(`"%s"`, us.Format("2006-01-02T15:04:05.000000Z"))
	}
	return []byte(str), nil
}

func (us *JsonTimeUs) UnmarshalJSON(text []byte) error {
	if len(text) == 2 {
		return nil
	}
	t, err := parseJsonTime(string(text))
	if err != nil {
		return err
	}
	n := t.UnixMicro()
	us.Time = time.Unix(n/1000000, (n%1000000)*1000).UTC()
	return nil
}