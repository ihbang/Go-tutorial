package main

import (
	"fmt"

	"github.com/ihbang/go-tutorial/url_checker/checker"
)

func main() {
	results := map[string]string{}
	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.youtube.com",
		"https://www.madeedam.com",
		"https://www.instagram.com",
	}
	for _, url := range urls {
		result := "SUCCESS"
		err := checker.HitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
		fmt.Println(url, result)
	}
}
