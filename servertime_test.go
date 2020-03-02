package servertime

import (
	"testing"
	"time"
)

func TestGetServerTime(t *testing.T) {
	url := "https://github.com"
	date, err := GetServerTime(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(date.Format(time.RFC3339))

	url = "abcde"
	date, err = GetServerTime(url)
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}