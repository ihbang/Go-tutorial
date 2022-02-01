package scrapper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("status code error: %d %s", res.StatusCode, res.Status)
	}
}

func GetNextPage(res *http.Response) string {
	result := ""
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("div.pagination > ul > li > a").Each(func(i int, s *goquery.Selection) {
		label, ok := s.Attr("aria-label")
		if ok && label == "다음" {
			href, _ := s.Attr("href")
			slice := strings.Split(href, "&")
			result = slice[1]
		}
	})
	return result
}

func GetPage(url string) *http.Response {
	res, err := http.Get(url)

	checkErr(err)
	checkCode(res)
	return res
}
