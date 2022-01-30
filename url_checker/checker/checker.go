package checker

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed error = errors.New("request failed [%d]")

type Website struct {
	url    string
	status error
}

func HitUrl(url string, c chan<- Website) {
	resp, err := http.Get(url)
	if resp.StatusCode >= 400 {
		err = fmt.Errorf(errRequestFailed.Error(), resp.StatusCode)
	}
	c <- Website{url: url, status: err}
}

func (w Website) Url() string {
	return w.url
}

func (w Website) Status() error {
	return w.status
}
