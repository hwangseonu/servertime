package servertime

import (
	"errors"
	"net/http"
	urlUtil "net/url"
	"time"
)

func parseUrl(url string) (string, error) {
	u, err := urlUtil.Parse(url)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" || u.Host == "" {
		return "", errors.New("url parsing error: " + url)
	}
	return u.Scheme + "://" + u.Host, err
}

func getDateHeader(url string) (string, error) {
	res, err := http.Head(url)
	if err != nil {
		return "", err
	}
	date := res.Header.Get("Date")
	if date == "" {
		return "", errors.New("did not provide time from server")
	}
	return date, nil
}

func getDateHeaderWithLatency(url string) (string, time.Duration, error) {
	s := time.Now()
	date, err := getDateHeader(url)
	l := time.Now().Sub(s)
	return date, l, err
}

func GetServerTime(url string) (time.Time, error) {
	t := time.Time{}
	u, err := parseUrl(url)
	if err != nil {
		return t, err
	}
	date, l, err := getDateHeaderWithLatency(u)
	if err != nil {
		return t, err
	}
	t, err = time.Parse(time.RFC1123, date)
	if err != nil {
		return t, err
	}
	return t.Add(l / 2), nil
}