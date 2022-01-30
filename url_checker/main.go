package main

import (
	"fmt"

	"github.com/ihbang/go-tutorial/url_checker/checker"
)

func main() {
	c := make(chan checker.Website)
	websites := []checker.Website{}
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
		go checker.HitUrl(url, c)
	}
	for i := 0; i < len(urls); i++ {
		website := <-c
		websites = append(websites, website)
	}
	for _, website := range websites {
		status := website.Status()
		if status != nil {
			fmt.Println(website.Url(), status)
		} else {
			fmt.Println(website.Url(), "success")
		}
	}
}
