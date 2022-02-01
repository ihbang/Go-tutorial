package main

import (
	"fmt"

	"github.com/ihbang/go-tutorial/job_scrapper/scrapper"
)

func main() {
	baseURL := "https://kr.indeed.com/jobs?q=golang"
	suffix := ""

	for {
		url := baseURL + "&" + suffix
		fmt.Println("Requesting", url)

		res := scrapper.GetPage(url)
		defer res.Body.Close()

		suffix = scrapper.GetNextPage(res)
		if suffix == "" {
			break
		}
	}
}
