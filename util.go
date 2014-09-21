package aoj

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func UnixToTime(unix uint64) time.Time {
	return time.Unix(int64(unix/1000), int64(unix%1000)*1000000)
}

func APIRequest(api string, values url.Values) ([]byte, error) {
	query := values.Encode()
	if query != "" {
		query = "?" + query
	}
	url := api + query

	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	res := bytes.Map(func(r rune) rune {
		switch r {
		case '\r', '\n':
			return -1
		default:
			return r
		}
	}, body)

	return res, nil
}
