package aoj

import (
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

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
