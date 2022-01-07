package tests

import (
	"testing"
	"time"
)

func GetTimezone(timezoneName string) *time.Location {
	loc, errLoadLoc := time.LoadLocation(timezoneName)
	if errLoadLoc != nil {
		panic(errLoadLoc)
	}
	return loc
}

func TestLocation(t *testing.T) {
	curDateTimeStr := time.Now().Format("2006-01-02 15:04:05")
	t.Log("Local DateTime String: ", curDateTimeStr)

	parsedTime, errParse := time.Parse(
		"2006-01-02 15:04:05",
		curDateTimeStr)
	if errParse == nil {
		t.Log("Parsed Local DateTime String without Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.ParseInLocation(
		"2006-01-02 15:04:05",
		curDateTimeStr,
		GetTimezone("Asia/Shanghai"))
	if errParse == nil {
		t.Log("Parsed Local DateTime String with Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.Parse(
		time.RFC3339,
		time.Now().Format(time.RFC3339))
	if errParse == nil {
		t.Log("Parsed Local RFC3339 String without Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.ParseInLocation(
		time.RFC3339,
		time.Now().Format(time.RFC3339),
		GetTimezone("Asia/Shanghai"))
	if errParse == nil {
		t.Log("Parsed Local RFC3339 String with Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.ParseInLocation(
		time.RFC3339,
		time.Now().Format(time.RFC3339),
		GetTimezone("US/Hawaii"))
	if errParse == nil {
		t.Log("Parsed Local RFC3339 String with another Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.Parse(
		time.RFC3339,
		time.Now().In(GetTimezone("US/Hawaii")).Format(time.RFC3339))
	if errParse == nil {
		t.Log("Parsed US/Hawaii RFC3339 String without Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.ParseInLocation(
		time.RFC3339,
		time.Now().In(GetTimezone("US/Hawaii")).Format(time.RFC3339),
		GetTimezone("US/Hawaii"))
	if errParse == nil {
		t.Log("Parsed US/Hawaii RFC3339 String with Timezone: ", parsedTime.String())
	}

	parsedTime, errParse = time.ParseInLocation(
		time.RFC3339,
		time.Now().In(GetTimezone("US/Hawaii")).Format(time.RFC3339),
		GetTimezone("Asia/Shanghai"))
	if errParse == nil {
		t.Log("Parsed US/Hawaii RFC3339 String with another Timezone: ", parsedTime.String())
	}

	curUnixTs := time.Now().Unix()
	t.Log("Local UNIX Timestamp: ", curUnixTs)

	t.Log("Parsed Local UNIX Timestamp without Timezone: ", time.Unix(curUnixTs, 0).String())
	t.Log("Parsed Local UNIX Timestamp with UTC Timezone: ", time.Unix(curUnixTs, 0).
		In(time.UTC).String())

	anotherUnixTs := time.Now().In(GetTimezone("US/Hawaii")).Unix()
	t.Log("UNIX Timestamp with another Timezone: ", anotherUnixTs)
}
