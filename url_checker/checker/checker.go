package checker

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed error = errors.New("request failed")

func HitUrl(url string) error {
	fmt.Println("Checking", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
