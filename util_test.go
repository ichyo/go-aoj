package aoj

import (
	"testing"
	"time"
)

func TestUnixToTime(t *testing.T) {
	d := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	if !d.Equal(UnixToTime(uint64(d.Unix()) * 1000)) {
		t.Error("fail: UnixToTime()")
		t.Error("d =", d)
		t.Error("d.Unix() =", d.Unix())
		t.Error("UnixToTime(d.Unix() * 1000) =", UnixToTime(uint64(d.Unix())*1000))
	}
}

func TestAPIRequest(t *testing.T) {
	test_url := "http://judge.u-aizu.ac.jp/onlinejudge/webservice/solved_record?user_id=solver"
	_, err := APIRequest(test_url, nil)
	if err != nil {
		t.Error("fail: APIRequest()")
		t.Error(err)
	}
}
