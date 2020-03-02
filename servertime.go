package main

import (
	"errors"
	"net/http"
	"time"
)

func Get(url string) (time.Time, error) {
	t := time.Time{}

	res, err := http.Get(url)
	if err != nil {
		return t, err
	}

	date := res.Header.Get("Date")

	if date == "" {
		return t, errors.New("did not provide time from server")
	}

	t, err = time.Parse(time.RFC1123, date)
	if err != nil {
		return t, err
	}
	return t, nil
}