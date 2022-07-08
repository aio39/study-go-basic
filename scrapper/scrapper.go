package scrapper

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

type extractedJob struct {
	id       string
	title    string
	location string
	// salary   string
	// summary  string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var url string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"

	var jobs []extractedJob

	c := make(chan []extractedJob)

	totalPages := getPages(url)
	fmt.Println("get", totalPages, "pages")

	for i := 0; i < totalPages; i++ {
		go getPage(i, url, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	// fmt.Println("Total jobs:", len(jobs))
	writeJobs(jobs)
	fmt.Println("Finish!")
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	pageUrl := url + "&start=" + strconv.Itoa(page*50)
	c := make(chan extractedJob)
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
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs

}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	defer w.Flush() // 아래 작업이 끝나고 파일에 입력

	headers := []string{"ID", "Title", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	// csv. Write 는 goroutine-safe 하지 않음. 코드를 잘 못 짜면 runtime panic이 일어남.
	// concurrency를 지원하는 라이브러리 사용  or Write All로 Batch 처리
	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location}
		jErr := w.Write(jobSlice)
		checkErr(jErr)
	}

}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Find("a").Attr("data-jk")
	title := CleanString(card.Find("a>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	c <- extractedJob{id: id, title: title, location: location}
}

func getPages(url string) int {
	pages := 0

	res, err := http.Get(url)
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

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
