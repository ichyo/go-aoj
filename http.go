package aoj

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

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
