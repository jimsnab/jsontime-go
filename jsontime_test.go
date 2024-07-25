package jsontime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type (
	timeStruct struct {
		When time.Time `json:"when"`
	}

	secStruct struct {
		When SecRes `json:"when"`
	}

	msStruct struct {
		When MsRes `json:"when"`
	}

	usStruct struct {
		When UsRes `json:"when"`
	}

	nsStruct struct {
		When NsRes `json:"when"`
	}

	formType int
)

const (
	formRegular formType = iota
	formSec
	formMs
	formUs
	formNs
)

func testVerifyTime(t *testing.T, parseTime, renderTime string, ft formType) {
	tt, err := time.Parse(time.RFC3339Nano, renderTime)
	if err != nil {
		panic("bad test input")
	}

	raw := fmt.Sprintf(`{"when":"%s"}`, parseTime)

	var parsed time.Time
	var rendered []byte
	var parseErr, renderErr error
	switch ft {
	case formRegular:
		var ts timeStruct
		parseErr = json.Unmarshal([]byte(raw), &ts)
		parsed = ts.When
	case formSec:
		var ts secStruct
		parseErr = json.Unmarshal([]byte(raw), &ts)
		parsed = ts.When.Time
		rendered, renderErr = json.Marshal(ts)
	case formMs:
		var ts msStruct
		parseErr = json.Unmarshal([]byte(raw), &ts)
		parsed = ts.When.Time
		rendered, renderErr = json.Marshal(ts)
	case formUs:
		var ts usStruct
		parseErr = json.Unmarshal([]byte(raw), &ts)
		parsed = ts.When.Time
		rendered, renderErr = json.Marshal(ts)
	case formNs:
		var ts nsStruct
		parseErr = json.Unmarshal([]byte(raw), &ts)
		parsed = ts.When.Time
		rendered, renderErr = json.Marshal(ts)
	}

	if parseErr != nil {
		t.Fatal(parseErr)
	}
	if renderErr != nil {
		t.Fatal(renderErr)
	}

	if !parsed.Equal(tt) {
		t.Errorf("timestamp %s is not the expected time %s", parseTime, tt.Format(time.RFC3339))
	}

	if rendered != nil {
		expected := fmt.Sprintf(`{"when":"%s"}`, renderTime)
		if string(rendered) != expected {
			t.Errorf("timestamp %s was rendered as %s, not %s", parseTime, string(rendered), expected)
		}
	}
}

func testVerifyTimeParseErr(t *testing.T, testData string) {
	raw := fmt.Sprintf(`{"when":"%s"}`, testData)

	var ts timeStruct
	if err := json.Unmarshal([]byte(raw), &ts); err == nil {
		t.Fatal("did not get expected error")
	}
}

func TestJsonTime(t *testing.T) {
	testVerifyTime(t, "2024-07-22T15:05:52.338001008Z", "2024-07-22T15:05:52.338001008Z", formRegular)
	testVerifyTime(t, "2024-07-22T15:05:52.338001Z", "2024-07-22T15:05:52.338001Z", formRegular)
	testVerifyTime(t, "2024-07-22T15:05:52.338Z", "2024-07-22T15:05:52.338Z", formRegular)
	testVerifyTime(t, "2024-07-22T15:05:52Z", "2024-07-22T15:05:52Z", formRegular)
}

func TestSecRes(t *testing.T) {
	testVerifyTime(t, "2024-07-22T15:05:52.338001008Z", "2024-07-22T15:05:52Z", formSec)
	testVerifyTime(t, "2024-07-22T15:05:52.338001Z", "2024-07-22T15:05:52Z", formSec)
	testVerifyTime(t, "2024-07-22T15:05:52.338Z", "2024-07-22T15:05:52Z", formSec)
	testVerifyTime(t, "2024-07-22T15:05:52Z", "2024-07-22T15:05:52Z", formSec)

	testVerifyTime(t, "2024-07-22T15:05:52.838001008Z", "2024-07-22T15:05:53Z", formSec)
	testVerifyTime(t, "2024-07-22T15:05:52.838001Z", "2024-07-22T15:05:53Z", formSec)
	testVerifyTime(t, "2024-07-22T15:05:52.838Z", "2024-07-22T15:05:53Z", formSec)
}

func TestMsRes(t *testing.T) {
	testVerifyTime(t, "2024-07-22T15:05:52.338001008Z", "2024-07-22T15:05:52.338Z", formMs)
	testVerifyTime(t, "2024-07-22T15:05:52.338001Z", "2024-07-22T15:05:52.338Z", formMs)
	testVerifyTime(t, "2024-07-22T15:05:52.338Z", "2024-07-22T15:05:52.338Z", formMs)
	testVerifyTime(t, "2024-07-22T15:05:52Z", "2024-07-22T15:05:52.000Z", formMs)

	testVerifyTime(t, "2024-07-22T15:05:52.338500008Z", "2024-07-22T15:05:52.339Z", formMs)
	testVerifyTime(t, "2024-07-22T15:05:52.338500Z", "2024-07-22T15:05:52.339Z", formMs)
}

func TestUsRes(t *testing.T) {
	testVerifyTime(t, "2024-07-22T15:05:52.338001008Z", "2024-07-22T15:05:52.338001Z", formUs)
	testVerifyTime(t, "2024-07-22T15:05:52.338001Z", "2024-07-22T15:05:52.338001Z", formUs)
	testVerifyTime(t, "2024-07-22T15:05:52.338Z", "2024-07-22T15:05:52.338000Z", formUs)
	testVerifyTime(t, "2024-07-22T15:05:52Z", "2024-07-22T15:05:52.000000Z", formUs)

	testVerifyTime(t, "2024-07-22T15:05:52.338001508Z", "2024-07-22T15:05:52.338002Z", formUs)
}

func TestNsRes(t *testing.T) {
	testVerifyTime(t, "2024-07-22T15:05:52.338001008Z", "2024-07-22T15:05:52.338001008Z", formNs)
	testVerifyTime(t, "2024-07-22T15:05:52.338001Z", "2024-07-22T15:05:52.338001000Z", formNs)
	testVerifyTime(t, "2024-07-22T15:05:52.338Z", "2024-07-22T15:05:52.338000000Z", formNs)
	testVerifyTime(t, "2024-07-22T15:05:52Z", "2024-07-22T15:05:52.000000000Z", formNs)
}

func TestJsonTimeMalformed(t *testing.T) {
	testVerifyTimeParseErr(t, "2024-07-22T15:05:52.33800!Z")
}

func TestNow(t *testing.T) {
	sr := SecResNow()
	msr := MsResNow()
	usr := UsResNow()
	nsr := NsResNow()

	d := nsr.Sub(sr.Time)
	if d.Abs().Seconds() > 1 {
		t.Error("sr fail")
	}

	d = nsr.Sub(msr.Time)
	if d.Abs().Seconds() > 1 {
		t.Error("msr fail")
	}

	d = nsr.Sub(usr.Time)
	if d.Abs().Seconds() > 1 {
		t.Error("usr fail")
	}
}
