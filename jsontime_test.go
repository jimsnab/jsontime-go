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
)

func testVerifyTime(t *testing.T, testData string, expected time.Time) {
	raw := fmt.Sprintf(`{"when":"%s"}`, testData)

	var ts timeStruct
	if err := json.Unmarshal([]byte(raw), &ts); err != nil {
		t.Fatal(err)
	}

	if !ts.When.Equal(expected) {
		t.Errorf("timestamp %s is not the expected time %s", testData, expected.Format(time.RFC3339))
	}
}

func testVerifyNotTime(t *testing.T, testData string, notExpected time.Time) {
	raw := fmt.Sprintf(`{"when":"%s"}`, testData)

	var ts timeStruct
	if err := json.Unmarshal([]byte(raw), &ts); err != nil {
		t.Fatal(err)
	}

	if ts.When.Equal(notExpected) {
		t.Errorf("timestamp %s matchines wrong time %s", testData, notExpected.Format(time.RFC3339))
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
	refTime, err := time.Parse(time.RFC3339Nano, "2024-07-22T15:05:52.338001000Z")
	if err != nil {
		t.Fatal(err)
	}
	testVerifyTime(t, "2024-07-22T15:05:52.338001Z", refTime)
	testVerifyNotTime(t, "2024-07-22T15:05:52.338Z", refTime)
	testVerifyNotTime(t, "2024-07-22T15:05:52Z", refTime)

	refTime, err = time.Parse(time.RFC3339Nano, "2024-07-22T15:05:52.338000000Z")
	if err != nil {
		t.Fatal(err)
	}

	testVerifyTime(t, "2024-07-22T15:05:52.338000Z", refTime)
	testVerifyTime(t, "2024-07-22T15:05:52.338Z", refTime)
	testVerifyNotTime(t, "2024-07-22T15:05:52Z", refTime)

	refTime, err = time.Parse(time.RFC3339Nano, "2024-07-22T15:05:52.000000000Z")
	if err != nil {
		t.Fatal(err)
	}
	testVerifyTime(t, "2024-07-22T15:05:52Z", refTime)
	testVerifyTime(t, "2024-07-22T15:05:52.000Z", refTime)
	testVerifyTime(t, "2024-07-22T15:05:52.000000Z", refTime)

}

func TestJsonTimeMalformed(t *testing.T) {
	testVerifyTimeParseErr(t, "2024-07-22T15:05:52.33800!Z")
}
