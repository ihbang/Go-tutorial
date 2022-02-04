package scrapper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type JobCard struct {
	Id       string
	Title    string
	Company  string
	Location string
	Salary   string
	Summary  string
}

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

func GetCards(url string) []JobCard {
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	result := []JobCard{}
	doc.Find("div.mosaic-provider-jobcards > a").Each(func(i int, s *goquery.Selection) {
		var job JobCard
		job.Id, _ = s.Attr("data-jk")
		job.Title = s.Find(".jobTitle").Text()
		job.Company = s.Find(".companyName").Text()
		job.Location = s.Find(".companyLocation").Text()
		job.Salary = s.Find(".salary-snippet").Text()
		job.Summary = s.Find(".job-snippet").Text()
		result = append(result, job)
	})
	return result
}

func GetNextPage(url string) string {
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	result := ""
	doc.Find(".pagination-list > li > a").Each(func(i int, s *goquery.Selection) {
		label, _ := s.Attr("aria-label")
		if label == "다음" {
			href, _ := s.Attr("href")
			slice := strings.Split(href, "&")
			result = slice[1]
		}
	})
	return result
}
