package main

import (
	"fmt"

	"github.com/ihbang/go-tutorial/job_scrapper/scrapper"
)

func main() {
	baseURL := "https://kr.indeed.com/jobs?q=python"
	suffix := ""
	jobs := []scrapper.JobCard{}

	for {
		url := baseURL + "&" + suffix
		fmt.Println("Requesting", url)

		slice := scrapper.GetCards(url)
		fmt.Println(slice)
		jobs = append(jobs, slice...)

		suffix = scrapper.GetNextPage(url)
		if suffix == "" {
			break
		}
	}
	fmt.Println(jobs)
}
