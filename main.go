package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJob struct {
	id       string
	title    string
	location string
	// salary   string
	// summary  string
}

func main() {
	var jobs []extractedJob

	totalPages := getPages()
	fmt.Println("get", totalPages, "pages")

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	// fmt.Println("Total jobs:", len(jobs))
	writeJobs(jobs)
	fmt.Println("Finish!")
}

func getPage(page int) []extractedJob {
	pageUrl := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobCard_mainContent ")

	var jobs []extractedJob

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs

}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	defer w.Flush() // 아래 작업이 끝나고 파일에 입력

	headers := []string{"ID", "Title", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location}
		jErr := w.Write(jobSlice)
		checkErr(jErr)
	}

}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Find("a").Attr("data-jk")
	title := cleanString(card.Find("a>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	return extractedJob{id: id, title: title, location: location}
}

func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
